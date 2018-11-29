package logger

import (
	"github.com/azer/logger"
	"os"
	"strings"
)

// Setup sets up global logger configuration.
func Setup(output *os.File, verbosityLevel, filterSettings string) {
	ow := logger.NewStandardOutput(output)
	logWriter, _ := ow.(logger.StandardWriter)

	// JSON will be printed if it is `false`.
	logWriter.ColorsEnabled = true

	defaultOutputSettings := parseVerbosityLevel(verbosityLevel)
	logWriter.Settings = parsePackageSettings(filterSettings, defaultOutputSettings)

	// TODO: it is more safe to replace runtime.writers[0]
	logger.Hook(logWriter)
}

// Accepts: foo,bar,qux@timer
//          *
//          *@error
//          *@error,database@timer
func parsePackageSettings(input string, defaultOutputSettings *logger.OutputSettings) map[string]*logger.OutputSettings {
	all := map[string]*logger.OutputSettings{}
	items := strings.Split(input, ",")

	for _, item := range items {
		name, verbosity := parsePackageName(item)
		if verbosity == nil {
			verbosity = defaultOutputSettings
		}

		all[name] = verbosity
	}

	return all
}

// Accepts: users
//          database@timer
//          server@error
func parsePackageName(input string) (string, *logger.OutputSettings) {
	parsed := strings.Split(input, "@")
	name := strings.TrimSpace(parsed[0])

	if len(parsed) > 1 {
		return name, parseVerbosityLevel(parsed[1])
	}

	return name, nil
}

func parseVerbosityLevel(val string) *logger.OutputSettings {
	val = strings.ToUpper(strings.TrimSpace(val))

	if val == "MUTE" {
		return &logger.OutputSettings{}
	}

	s := &logger.OutputSettings{
		Info:  true,
		Timer: true,
		Error: true,
	}

	if val == "TIMER" {
		s.Info = false
	}

	if val == "ERROR" {
		s.Info = false
		s.Timer = false
	}

	return s
}
