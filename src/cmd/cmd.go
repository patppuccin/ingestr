package cmd

const (
	helpRootCmd    = "File ingestion tool for hoarders and archivers"
	helpCleanupCmd = "Remove remnant files from managed temporary locations"
	helpInitCmd    = "Initialize example config file"
	helpRunCmd     = "Runs the application"
)

var (
	flagGlobalLogLevel string
	flagGlobalQuiet    bool
)
