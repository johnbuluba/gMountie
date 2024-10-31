package fs

import (
	"embed"
	"gmountie/test/e2e/utils"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Embed fio config files
//
//go:embed fio/*.fio
var configs embed.FS

type FioTestSuite struct {
	suite.Suite
	testAppCtx *utils.AppTestingContext
	volume     *utils.TestVolume
}

func (s *FioTestSuite) SetupSuite() {
	// Start the app.
	testAppCtx, err := utils.NewAppTestingContext(
		utils.WithBasicAuth("test", "test"),
		utils.WithRandomTestVolume(false),
	)
	utils.Must0(s, err)
	utils.Must0(s, testAppCtx.Start())

	s.testAppCtx = testAppCtx
	// Copy the fio config files to the volume.
	s.volume = s.testAppCtx.GetVolumes()[0]
	scriptsPath := filepath.Join(s.volume.GetRootPath(), "scripts")
	utils.Must0(s, os.Mkdir(scriptsPath, 0o700))
	utils.Must0(s, CopyEmbedFiles(configs, scriptsPath))

	// Mount the volume.
	s.testAppCtx.MountVolume(s.volume)
}

func (s *FioTestSuite) TestFS() {
	entries, err := configs.ReadDir("fio")
	if err != nil {
		s.T().Fatal(err)
	}

	for _, entry := range entries {
		s.Run(entry.Name(), func() {
			path := filepath.Join(s.volume.GetRootPath(), "scripts", entry.Name())
			cmd := exec.Command("fio", "--output-format=json+", path)
			cmd.Dir = s.volume.GetMountPath()
			out, err := cmd.Output()

			if err != nil {
				s.T().Fatal(string(out), err)
			}
			s.T().Log(string(out))
		})
	}
}

func (s *FioTestSuite) TearDownSuite() {
	err := s.testAppCtx.Close()
	if err != nil {
		s.T().Fatal(err)
	}
}

// ----------- Helpers

// CopyEmbedFiles copies the embed files to the destination.
func CopyEmbedFiles(src embed.FS, dest string) error {
	entries, err := src.ReadDir("fio")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		data, err := src.ReadFile(filepath.Join("fio", entry.Name()))
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(dest, entry.Name()), data, fs.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestFioTestSuite(t *testing.T) {
	_, err := exec.LookPath("fio")
	if err != nil {
		t.Skip("fio not found, skipping test")
	}
	suite.Run(t, new(FioTestSuite))
}
