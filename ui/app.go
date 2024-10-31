package main

import (
	"context"
	"fmt"
	"gmountie/pkg/client"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/common"
	"gmountie/pkg/utils/log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

// App struct
type App struct {
	ctx    context.Context
	appCtx *client.AppContext
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// shutdown is called when the app is shutting down
func (a *App) shutdown(ctx context.Context) {
	if a.appCtx != nil {
		err := a.appCtx.Close()
		if err != nil {
			log.Log.Error("error closing app context", zap.Error(err))
		}
	}
}

// Login is called when the user logs in
func (a *App) Login(endpoint, username, password string) (bool, error) {
	c, err := grpc.NewClient(endpoint, grpc.WithBasicAuth(username, password))
	if err != nil {
		return false, err
	}
	c.Connect()
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Log.Sugar().Fatalf("failed to get home directory: %v", err)
	}
	path := filepath.Join(homePath, "mnt", "gmountie")
	a.appCtx = client.NewAppContext(c, path)
	err = a.appCtx.MultiVolumeMounter.Start()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Logout is called when the user logs out
func (a *App) Logout() error {
	if a.appCtx == nil {
		return nil
	}
	err := a.appCtx.Close()
	if err != nil {
		return err
	}
	a.appCtx = nil
	return nil
}

// GetVolumes returns a list of volumes
func (a *App) GetVolumes() ([]common.Volume, error) {
	if a.appCtx == nil {
		return nil, fmt.Errorf("not logged in")
	}
	return a.appCtx.VolumeService.GetVolumes(a.ctx)
}

// IsMounted checks if a volume is mounted
func (a *App) IsMounted(volume common.Volume) (bool, error) {
	if a.appCtx == nil {
		return false, fmt.Errorf("not logged in")
	}
	return a.appCtx.MultiVolumeMounter.IsVolumeMounted(volume.Name), nil
}

// Mount mounts a volume
func (a *App) Mount(volume common.Volume) error {
	return a.appCtx.MultiVolumeMounter.Mount(volume.Name)
}

// Unmount unmounts a volume
func (a *App) Unmount(volume common.Volume) error {
	if a.appCtx == nil {
		return fmt.Errorf("not logged in")
	}
	return a.appCtx.MultiVolumeMounter.Unmount(volume.Name)
}
