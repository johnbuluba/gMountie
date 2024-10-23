package service

import (
	"context"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/client/io"
)

// MounterService is a service that mounts volumes
type MounterService struct {
	client        *grpc.Client
	volumeService *VolumeService
}

// NewMounterService creates a new MounterService
func NewMounterService(client *grpc.Client, volumeService *VolumeService) *MounterService {
	return &MounterService{
		client:        client,
		volumeService: volumeService,
	}
}

// Mount mounts a volume
func (m *MounterService) Mount(ctx context.Context, volume, path string) error {
	fs := io.NewLocalFileSystem(m.client)
}
