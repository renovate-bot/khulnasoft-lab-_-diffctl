package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	gosentry "github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/khulnasoft-lab/driftctl/build"
	"github.com/khulnasoft-lab/driftctl/logger"
	"github.com/khulnasoft-lab/driftctl/pkg/cmd"
	cmderrors "github.com/khulnasoft-lab/driftctl/pkg/cmd/errors"
	"github.com/khulnasoft-lab/driftctl/pkg/cmd/scan"
	"github.com/khulnasoft-lab/driftctl/pkg/config"
	"github.com/khulnasoft-lab/driftctl/pkg/version"
	"github.com/khulnasoft-lab/driftctl/sentry"
	"github.com/sirupsen/logrus"
)

func init() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load() // The Original .env
}

func main() {
	os.Exit(run())
}

func run() int {

	config.Init()
	logger.Init()
	build := build.Build{}
	// Check whether driftCTL is run under Tringle CLI
	isTringle := config.IsTringle()
	logrus.WithFields(logrus.Fields{
		"isRelease":               fmt.Sprintf("%t", build.IsRelease()),
		"isUsageReportingEnabled": fmt.Sprintf("%t", build.IsUsageReportingEnabled()),
		"version":                 version.Current(),
		"isTringle":               fmt.Sprintf("%t", isTringle),
	}).Debug("Build info")

	// Enable colorization when driftctl is launched under khulnasoft-lab cli (piped)
	if isTringle {
		color.NoColor = false
	}

	driftctlCmd := cmd.NewDriftctlCmd(build)

	checkVersion := driftctlCmd.ShouldCheckVersion()
	latestVersionChan := make(chan string)
	if checkVersion {
		go func() {
			latestVersion := version.CheckLatest()
			latestVersionChan <- latestVersion
		}()
	}

	// Handle panic and log them to sentry if error reporting is enabled
	defer func() {
		if cmd.IsReportingEnabled(&driftctlCmd.Command) {
			err := recover()
			if err != nil {
				gosentry.CurrentHub().Recover(err)
				flushSentry()
				logrus.Fatalf("Captured panic: %s", err)
				os.Exit(scan.EXIT_ERROR)
			}
			flushSentry()
		}
	}()

	if _, err := driftctlCmd.ExecuteC(); err != nil {
		if _, isNotInSync := err.(cmderrors.InfrastructureNotInSync); isNotInSync {
			return scan.EXIT_NOT_IN_SYNC
		}
		if cmd.IsReportingEnabled(&driftctlCmd.Command) {
			sentry.CaptureException(err)
		}
		_, _ = fmt.Fprintln(os.Stderr, color.RedString("%s", err))
		return scan.EXIT_ERROR
	}

	if checkVersion {
		newVersion := <-latestVersionChan
		if newVersion != "" {
			_, _ = fmt.Fprintln(os.Stderr, "\n\nYour version of driftctl is outdated, please upgrade!")
			_, _ = fmt.Fprintf(os.Stderr, "Current: %s; Latest: %s\n", version.Current(), newVersion)
		}
	}

	return scan.EXIT_IN_SYNC
}

func flushSentry() {
	ttl := 60 * time.Second
	ok := gosentry.Flush(ttl)
	logrus.WithField("timeout", ttl).WithField("success", ok).Debug("Flushed Sentry events")
}
