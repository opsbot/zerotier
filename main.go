package main

import (
	"os"

	"github.com/opsbot/zerotier/cli"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	log.SetFormatter(&logrus.TextFormatter{})

	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat:  "",
	// 	DisableTimestamp: false,
	// 	DataKey:          "",
	// 	PrettyPrint:      true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(logrus.FatalLevel)

	// log.SetReportCaller(true)
}

func main() {
	cli.Execute(version, commit, date)
}
