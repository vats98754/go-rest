package main

import (
	"github.com/vats98754/restapi/controller"
	"github.com/vats98754/restapi/model"
)

func main() {
	model.Init()
	controller.Start()
}