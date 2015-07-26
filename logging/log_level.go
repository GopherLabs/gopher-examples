package main

import (
	// Only for the purpose of syntax clarity, we are using dot import notation below.
	// It should be noted, however, that the Go team does not recommend using the dot
	// import since it can cause some odd behaviour in certain cases.
	"time"

	. "github.com/gopherlabs/gopher"
)

func main() {

	App.Config(Config{
		KEY_LOGGER: ConfigLogger{
			TimestampFormat: time.RFC822,
			LogLevel:        LEVEL_INFO,
		},
	})

	// Debug logs will not be logged since we set the Log Level to LEVEL_INFO
	Log.Debug("Useful debugging information.")

	// Anything with severity Info or above it will be logged
	Log.Info("Something noteworthy happened!")

	Log.Warn("You should probably take a look at this.")

	Log.Error("Something failed but I'm not quitting.")

	ListenAndServe()
}
