package main

import (
	// Only for the purpose of syntax clarity, we are using dot import notation below.
	// It should be noted, however, that the Go team does not recommend using the dot
	// import since it can cause some odd behaviour in certain cases.
	. "github.com/gopherlabs/gopher"
)

func main() {

	Log.Debug("Useful debugging information.")

	Log.Info("Something noteworthy happened!")

	Log.Warn("You should probably take a look at this.")

	Log.Error("Something failed but I'm not quitting.")

	// Calls os.Exit(1) after logging
	Log.Fatal("Bye.")

	// Calls panic() after logging
	Log.Panic("I'm bailing.")
}
