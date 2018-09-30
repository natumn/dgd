package main

import (
	"os"
)

func main() {
	cli := New()

	os.Exit(
		cli.Run(),
	)
}
