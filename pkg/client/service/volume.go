package service

import (
	"context"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/common"
	"gmountie/pkg/proto"
)

// VolumeService is a service that provides information about the volumes
type VolumeService struct {
	client *grpc.Client
}

// NewVolumeService creates a new VolumeService
func NewVolumeService(client *grpc.Client) *VolumeService {
	return &VolumeService{
		client: client,
	}
}

// GetVolumes returns a list of volumes
func (v *VolumeService) GetVolumes(ctx context.Context) ([]common.Volume, error) {
	reply, err := v.client.Volume.List(ctx, &proto.VolumeListRequest{})
	if err != nil {
		return nil, err
	}
	volumes := make([]common.Volume, 0)
	for _, v := range reply.Volumes {
		volumes = append(volumes, common.Volume{Name: v.Name})
	}
	return volumes, nil
}
