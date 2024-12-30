package main

import (
	"github.com/aleksnew2/deltarune_engine/io/console"
)

func main() {
	var console console.Console
	console.Settings.IsWithTimestamp = false
	console.Settings.IsWithExceptionCode = false
	
	console.Error(-1, "HI!")
}
