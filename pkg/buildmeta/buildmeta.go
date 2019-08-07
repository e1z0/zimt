package buildmeta

import "runtime"

// Build meta information, populated at build-time
var (
	GitTag    string
	GitCommit string
	GitBranch string
	BuildDate string
	BuildOS   string
	Platform  string
	Compiler  = runtime.Version()
)
