package common

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

const (
	VolumeHeader = "Volume"
)

// AddVolumeHeader adds the volume header to the context.
func AddVolumeHeader(ctx context.Context, volume string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, VolumeHeader, volume)
	return ctx
}

// GetVolumeHeader gets the volume header from the context.
func GetVolumeHeader(ctx context.Context) (string, error) {
	md := metadata.ValueFromIncomingContext(ctx, VolumeHeader)
	if len(md) != 1 {
		return "", fmt.Errorf("expected exactly one volume in metadata, got %d", len(md))
	}
	return md[0], nil
}
