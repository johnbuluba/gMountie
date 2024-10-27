package utils

import (
	"gmountie/pkg/utils/log"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

const (
	VolumeSrc   = "src"
	VolumeMount = "mount"
)

// TestVolume is a struct that holds the test volume.
type TestVolume struct {
	// Name is the name of the volume.
	Name string
	// GeneratedFiles
	GeneratedFiles []string
	// path is the path of the volume.
	path string
}

// NewTestVolume creates a new TestVolume.
func NewTestVolume(name string, createRandomFiles bool) (*TestVolume, error) {
	path, err := os.MkdirTemp("", "gmountie-test-*")
	if err != nil {
		return nil, err
	}
	v := &TestVolume{
		Name: name,
		path: path,
	}

	// Create subdirectories
	err = os.Mkdir(v.GetSrcPath(), 0o700)
	if err != nil {
		return nil, err
	}
	err = os.Mkdir(v.GetMountPath(), 0o700)
	if err != nil {
		return nil, err
	}
	log.Log.Info("created test volume", zap.String("name", name), zap.String("path", path))
	if createRandomFiles {
		err = v.createRandomFiles()
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}

// Close removes the volume.
func (v *TestVolume) Close() error {
	log.Log.Info("removing test volume", zap.String("name", v.Name), zap.String("path", v.path))
	return os.RemoveAll(v.path)
}

// GetSrcPath returns the source path of the volume.
func (v *TestVolume) GetSrcPath() string {
	return filepath.Join(v.path, VolumeSrc)
}

// GetMountPath returns the mount path of the volume.
func (v *TestVolume) GetMountPath() string {
	return filepath.Join(v.path, VolumeMount)
}

// createRandomFiles creates random files in the source path.
func (v *TestVolume) createRandomFiles() error {
	g := &RandomFileGenerator{
		fileSize:     1024 * 1024,
		fanoutDepth:  2,
		fanoutFiles:  2,
		fanoutDirs:   2,
		randomSize:   true,
		randomFanout: true,
	}
	log.Log.Info("creating random files", zap.String("path", v.GetSrcPath()))
	err := g.WriteRandomFiles(v.GetSrcPath())
	if err != nil {
		return err
	}
	// Remove the source path from the generated files.
	v.GeneratedFiles = make([]string, 0, len(g.files))
	for _, f := range g.files {
		trimmed := strings.TrimPrefix(f, v.GetSrcPath())
		trimmed = strings.TrimPrefix(trimmed, "/")
		v.GeneratedFiles = append(v.GeneratedFiles, trimmed)
	}
	return nil
}
