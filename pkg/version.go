package pkg

var (
	version = "unknown"
	date    = "unknown"
	commit  = "unknown"
)

// BuildInfo contains the build information
type BuildInfo struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Commit  string `json:"commit"`
}

// GetBuildInfo returns the build information
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version: version,
		Date:    date,
		Commit:  commit,
	}
}
