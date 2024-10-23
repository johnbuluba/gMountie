package controller

import (
	"context"
	"errors"
	"gmountie/pkg/proto"
	"gmountie/pkg/server/service"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/puzpuzpuz/xsync/v3"
	"google.golang.org/grpc"
)

type fileEntry struct {
	file nodefs.File
	path string
}

type RpcFileServerImpl struct {
	fsService service.VolumeService
	files     *xsync.MapOf[string, *fileEntry]
	proto.UnimplementedRpcFileServer
}

// Verify that RpcFileServerImpl implements proto.RpcFileServer
var _ proto.RpcFileServer = (*RpcFileServerImpl)(nil)

// NewRpcFileServer creates a new RpcFileServerImpl
func NewRpcFileServer(fsService service.VolumeService) *RpcFileServerImpl {
	return &RpcFileServerImpl{
		fsService: fsService,
		files:     xsync.NewMapOf[string, *fileEntry](),
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
	fd, status := fs.Open(request.Path, request.Flags, createContext(ctx, request.Caller))
	reply := &proto.OpenReply{
		Status: int32(status),
	}
	r.addFile(request.Path, fd)
	return reply, nil
}

func (r *RpcFileServerImpl) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateReply, error) {
	fs, err := r.fsService.GetVolumeFileSystem(request.Volume)
	if err != nil {
		return nil, err
	}
	fd, status := fs.Create(request.Path, request.Flags, request.Mode, createContext(ctx, request.Caller))
	reply := &proto.CreateReply{
		Status: int32(status),
	}
	r.addFile(request.Path, fd)
	return reply, nil
}

func (r *RpcFileServerImpl) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, request.Size)
	n, status := file.file.Read(buf, request.Offset)
	if status != fuse.OK {
		return &proto.ReadReply{
			Status: int32(status),
		}, nil
	}
	buf, status = n.Bytes(buf)
	reply := &proto.ReadReply{
		Size:   int64(n.Size()),
		Bytes:  buf,
		Status: int32(status),
	}
	return reply, nil
}

func (r *RpcFileServerImpl) Write(ctx context.Context, request *proto.WriteRequest) (*proto.WriteReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	written, status := file.file.Write(request.Bytes, request.Offset)

	reply := &proto.WriteReply{
		Written: written,
		Status:  int32(status),
	}
	return reply, nil
}

// Fsync is called to flush the file data
func (r *RpcFileServerImpl) Fsync(ctx context.Context, request *proto.FsyncRequest) (*proto.FsyncReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	status := file.file.Fsync(int(request.Flags))
	return &proto.FsyncReply{
		Status: int32(status),
	}, nil
}

// Release is called when the file is closed
func (r *RpcFileServerImpl) Release(ctx context.Context, request *proto.ReleaseRequest) (*proto.ReleaseReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	file.file.Release()
	r.files.Delete(request.Path)
	return &proto.ReleaseReply{}, nil
}

// Flush is called when the file is closed
func (r *RpcFileServerImpl) Flush(ctx context.Context, request *proto.FlushRequest) (*proto.FlushReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	status := file.file.Flush()
	return &proto.FlushReply{
		Status: int32(status),
	}, nil
}

// ----- File locking

// GetLk returns existing lock information for file.
func (r *RpcFileServerImpl) GetLk(ctx context.Context, request *proto.GetLkRequest) (*proto.GetLkReply, error) {
	file, err := r.getFile(request.Path)
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
	status := file.file.GetLk(request.Owner, lock, request.Flags, out)
	reply := &proto.GetLkReply{
		Lk: &proto.FileLock{
			Start: out.Start,
			End:   out.End,
			Typ:   out.Typ,
			Pid:   out.Pid,
		},
		Status: int32(status),
	}
	return reply, nil
}

// SetLk sets or clears the lock described by lk on file.
func (r *RpcFileServerImpl) SetLk(ctx context.Context, request *proto.SetLkRequest) (*proto.SetLkReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	lock := &fuse.FileLock{
		Start: request.Lk.Start,
		End:   request.Lk.End,
		Typ:   request.Lk.Typ,
		Pid:   request.Lk.Pid,
	}
	status := file.file.SetLk(request.Owner, lock, request.Flags)
	return &proto.SetLkReply{
		Status: int32(status),
	}, nil
}

// SetLkw sets or clears the lock described by lk. This call blocks until the operation can be completed.
func (r *RpcFileServerImpl) SetLkw(ctx context.Context, request *proto.SetLkwRequest) (*proto.SetLkwReply, error) {
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	lock := &fuse.FileLock{
		Start: request.Lk.Start,
		End:   request.Lk.End,
		Typ:   request.Lk.Typ,
		Pid:   request.Lk.Pid,
	}
	status := file.file.SetLkw(request.Owner, lock, request.Flags)
	return &proto.SetLkwReply{
		Status: int32(status),
	}, nil
}

func (r *RpcFileServerImpl) addFile(path string, file nodefs.File) {
	r.files.Store(path, &fileEntry{
		file: file,
		path: path,
	})
}

func (r *RpcFileServerImpl) getFile(path string) (*fileEntry, error) {
	file, ok := r.files.Load(path)
	if !ok {
		return nil, errors.New("file not found")
	}
	return file, nil
}
