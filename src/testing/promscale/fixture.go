package promscale

import (
	"fmt"
	"log"
	"strings"
	"time"

	testutils "github.com/tricorder/src/testing/bazel"
	docker "github.com/tricorder/src/testing/docker"
	pgutils "github.com/tricorder/src/utils/pg"
)

const (
	promscalePort = 9201
	promscalePath = "src/testing/promscale/promscale-image.tar"

	timescaledbPort = 5432
	timescaledbPath = "src/testing/timescaledb/image.tar"
)

var (
	promscaleImageName   string
	timescaledbImageName string
)

// init loads the pulled tsdb and promscale container image when importing this package.
// So the postgres image only needs to be pulled once during one bazel command run.
func init() {
	var cli docker.CLI
	promscaleImageName = cli.Load(testutils.TestFilePath(promscalePath))
	timescaledbImageName = cli.Load(testutils.TestFilePath(timescaledbPath))
}

// LaunchContainer returns a cleaner function, the client connected to the started container, and error if failed.
// You can then destroy the fixtures by deferring statement
func LaunchContainer() (func() error, *pgutils.Client, error) {
	log.Println("Not launching promscale container, will do that for testing prom/otel ingestion ", promscaleImageName)

	tsdbRunner := &docker.Runner{
		ImageName: timescaledbImageName,
		RdyMsg:    "database system is ready to accept connections",
		Options:   []string{"--env=POSTGRES_PASSWORD=passwd"},
	}
	err := tsdbRunner.Launch(10 * time.Second)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to start tsdb server, error: %v", err)
	}

	tsdbGatewayIP, err := tsdbRunner.GetGatewayIP()
	if err != nil {
		return nil, nil, err
	}

	tsdbPort, err := tsdbRunner.GetExposedPort(timescaledbPort)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"failed to get postgres container's exposed port for %d, error: %v",
			timescaledbPort,
			err,
		)
	}

	tsdbURL := fmt.Sprintf("postgresql://postgres:passwd@%s:%d", strings.TrimSpace(tsdbGatewayIP), tsdbPort)
	pgClient := pgutils.NewClient(tsdbURL)
	if err := pgClient.Connect(); err != nil {
		return nil, nil, fmt.Errorf("Unable to create client to tsdb at %s, error: %v", tsdbURL, err)
	}

	// todo: add prometheus and jaeger client connector

	cleanerFn := func() error {
		pgClient.Close()
		if err := tsdbRunner.Stop(); err != nil {
			return err
		}
		return nil
	}
	return cleanerFn, pgClient, nil
}
