package types

// SystemInfo holds information about the system.
type SystemInfo struct {
	OS           string
	Architecture string
	NumCPU       int
	GoVersion    string
}
