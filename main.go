package main

import (
	"rURLr/model"
	"rURLr/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
