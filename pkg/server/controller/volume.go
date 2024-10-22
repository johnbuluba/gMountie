package controller

import (
	"context"
	"gmountie/pkg/proto"
	"gmountie/pkg/server/service"

	"google.golang.org/grpc"
)

// VolumeServiceImpl implements the VolumeService gRPC service
type VolumeServiceImpl struct {
	service service.VolumeService
	proto.UnimplementedVolumeServiceServer
}

// NewVolumeService creates a new VolumeServiceImpl
func NewVolumeService(service service.VolumeService) *VolumeServiceImpl {
	return &VolumeServiceImpl{
		service: service,
	}
}

// Register registers the gRPC server
func (v *VolumeServiceImpl) Register(server *grpc.Server) {
	proto.RegisterVolumeServiceServer(server, v)
}

// List lists all volumes
func (v *VolumeServiceImpl) List(context.Context, *proto.VolumeListRequest) (*proto.VolumeListReply, error) {
	// Find volumes
	volumes, err := v.service.List()
	if err != nil {
		return nil, err
	}
	// Convert volumes to gRPC reply
	volumesReply := make([]*proto.Volume, 0, len(volumes))
	for _, volume := range volumes {
		volumesReply = append(volumesReply, &proto.Volume{
			Name: volume.Name,
		})
	}

	reply := &proto.VolumeListReply{
		Volumes: volumesReply,
	}
	return reply, nil
}
