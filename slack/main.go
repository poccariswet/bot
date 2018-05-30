package main

import "os"

const (
	port = os.Getenv("PORT")
)

func main() {
	os.Exit(run())
}

func run() int {
	return 0
}
