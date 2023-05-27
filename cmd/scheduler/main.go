package main

import "scheduler/cmd/scheduler/server"

func main() {
	sv := server.New()
	sv.Start()
}
