package main

import (
	"go101/conf"
	"go101/server"
)

func main() {
	conf.Init()

	r := server.NewRouter()
	r.Run(":3000")
}
