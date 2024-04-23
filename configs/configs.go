package configs

import (
	"flag"
	"os"

	"github.com/mist-gopher/feature-tag/pkg/logger"
)

type Configuration struct {
	LogLevel logger.Level
}

type configflags struct {
	debug bool
}

type configenv struct {
	loglevel string
}

func loadflags() configflags {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set app to debug level")
	flag.Parse()

	return configflags{
		debug: debug,
	}
}

func loadenv() configenv {
	return configenv{
		loglevel: os.Getenv("LOG_LEVEL"),
	}
}

func Load() Configuration {
	f := loadflags()
	e := loadenv()

	level := logger.INFO

	if e.loglevel != "" {
		level = logger.Level(e.loglevel)
	}

	if f.debug {
		level = logger.DEBUG
	}

	return Configuration{
		LogLevel: level,
	}
}
