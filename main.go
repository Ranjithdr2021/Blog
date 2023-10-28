package main

import "blog/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}
