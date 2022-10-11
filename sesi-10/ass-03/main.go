package main

import (
	"ass-03/routers"
)

func main() {
	r := routers.SetupRoutes()
	r.Run()
}
