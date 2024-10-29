package controller

import (
	"context"
	"gmountie/pkg/proto"
	"gmountie/pkg/server/service"
	"sync/atomic"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/puzpuzpuz/xsync/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fileEntry struct {
	file nodefs.File
	path string
	fd   uint64
}

type RpcFileServerImpl struct {
	fsService service.VolumeService
	files     *xsync.MapOf[uint64, *fileEntry]
	proto.UnimplementedRpcFileServer
	fdNum atomic.Uint64
}

// Verify that RpcFileServerImpl implements proto.RpcFileServer
var _ proto.RpcFileServer = (*RpcFileServerImpl)(nil)

// NewRpcFileServer creates a new RpcFileServerImpl
func NewRpcFileServer(fsService service.VolumeService) *RpcFileServerImpl {
	return &RpcFileServerImpl{
		fsService: fsService,
		files:     xsync.NewMapOf[uint64, *fileEntry](),
	}
}

// Register registers the gRPC server
func (r *RpcFileServerImpl) Register(server *grpc.Server) {
	proto.RegisterRpcFileServer(server, r)
}

func (r *RpcFileServerImpl) Open(ctx context.Context, request *proto.OpenRequest) (*proto.OpenReply, error) {
	fs, err := r.fsService.GetVolumeFileSystem(request.Volume)
	if err != nil {
		return nil, err
	}
	fd, s := fs.Open(request.Path, request.Flags, createContext(ctx, request.Caller))
	privateFd := r.fdNum.Add(1)
	reply := &proto.OpenReply{
		Fd:     privateFd,
		Status: int32(s),
	}
	r.addFile(privateFd, request.Path, fd)
	return reply, nil
}

func (r *RpcFileServerImpl) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateReply, error) {
	fs, err := r.fsService.GetVolumeFileSystem(request.Volume)
	if err != nil {
		return nil, err
	}
	fd, s := fs.Create(request.Path, request.Flags, request.Mode, createContext(ctx, request.Caller))
	privateFd := r.fdNum.Add(1)
	reply := &proto.CreateReply{
		Fd:     privateFd,
		Status: int32(s),
	}
	r.addFile(privateFd, request.Path, fd)
	return reply, nil
}

func (r *RpcFileServerImpl) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, request.Size)
	n, s := file.file.Read(buf, request.Offset)
	if s != fuse.OK {
		return &proto.ReadReply{
			Status: int32(s),
		}, nil
	}
	buf, s = n.Bytes(buf)
	reply := &proto.ReadReply{
		Size:   int64(n.Size()),
		Bytes:  buf,
		Status: int32(s),
	}
	return reply, nil
}

func (r *RpcFileServerImpl) Write(ctx context.Context, request *proto.WriteRequest) (*proto.WriteReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	written, s := file.file.Write(request.Bytes, request.Offset)

	reply := &proto.WriteReply{
		Written: written,
		Status:  int32(s),
	}
	return reply, nil
}

// Fsync is called to flush the file data
func (r *RpcFileServerImpl) Fsync(ctx context.Context, request *proto.FsyncRequest) (*proto.FsyncReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	s := file.file.Fsync(int(request.Flags))
	return &proto.FsyncReply{
		Status: int32(s),
	}, nil
}

// Release is called when the file is closed
func (r *RpcFileServerImpl) Release(ctx context.Context, request *proto.ReleaseRequest) (*proto.ReleaseReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	file.file.Release()
	r.files.Delete(request.Fd)
	return &proto.ReleaseReply{}, nil
}

// Flush is called when the file is closed
func (r *RpcFileServerImpl) Flush(ctx context.Context, request *proto.FlushRequest) (*proto.FlushReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	s := file.file.Flush()
	return &proto.FlushReply{
		Status: int32(s),
	}, nil
}

// ----- File locking

// GetLk returns existing lock information for file.
func (r *RpcFileServerImpl) GetLk(ctx context.Context, request *proto.GetLkRequest) (*proto.GetLkReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	lock := &fuse.FileLock{
		Start: request.Lk.Start,
		End:   request.Lk.End,
		Typ:   request.Lk.Typ,
		Pid:   request.Lk.Pid,
	}
	out := &fuse.FileLock{}
	s := file.file.GetLk(request.Owner, lock, request.Flags, out)
	reply := &proto.GetLkReply{
		Lk: &proto.FileLock{
			Start: out.Start,
			End:   out.End,
			Typ:   out.Typ,
			Pid:   out.Pid,
		},
		Status: int32(s),
	}
	return reply, nil
}

// SetLk sets or clears the lock described by lk on file.
func (r *RpcFileServerImpl) SetLk(ctx context.Context, request *proto.SetLkRequest) (*proto.SetLkReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	lock := &fuse.FileLock{
		Start: request.Lk.Start,
		End:   request.Lk.End,
		Typ:   request.Lk.Typ,
		Pid:   request.Lk.Pid,
	}
	s := file.file.SetLk(request.Owner, lock, request.Flags)
	return &proto.SetLkReply{
		Status: int32(s),
	}, nil
}

// SetLkw sets or clears the lock described by lk. This call blocks until the operation can be completed.
func (r *RpcFileServerImpl) SetLkw(ctx context.Context, request *proto.SetLkwRequest) (*proto.SetLkwReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	lock := &fuse.FileLock{
		Start: request.Lk.Start,
		End:   request.Lk.End,
		Typ:   request.Lk.Typ,
		Pid:   request.Lk.Pid,
	}
	s := file.file.SetLkw(request.Owner, lock, request.Flags)
	return &proto.SetLkwReply{
		Status: int32(s),
	}, nil
}

// Allocate allocates space for a file
func (r *RpcFileServerImpl) Allocate(ctx context.Context, request *proto.AllocateRequest) (*proto.AllocateReply, error) {
	file, err := r.getFile(request.Fd)
	if err != nil {
		return nil, err
	}
	s := file.file.Allocate(request.Off, request.Size, request.Mode)
	return &proto.AllocateReply{
		Status: int32(s),
	}, nil
}

func (r *RpcFileServerImpl) addFile(fd uint64, path string, file nodefs.File) {
	r.files.Store(fd, &fileEntry{
		file: file,
		path: path,
		fd:   fd,
	})
}

func (r *RpcFileServerImpl) getFile(fd uint64) (*fileEntry, error) {
	file, ok := r.files.Load(fd)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "file with fd: %d not found", fd)
	}
	return file, nil
}
