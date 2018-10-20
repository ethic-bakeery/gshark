package util

import (
	"gshark/util/githubsearch"

	"github.com/urfave/cli"

	"time"
	"gshark/logger"
)

func Scan(ctx *cli.Context) {
	// seconds
	var Interval time.Duration = 900

	if ctx.IsSet("time") {
		Interval = time.Duration(ctx.Int("time"))
	}

	logger.Log.Println("scan github code ")
	// use go keyword or not
	go githubsearch.ScheduleTasks(Interval)
}
