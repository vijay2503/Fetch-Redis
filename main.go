package main

import (
	"qwik/router"
	s "qwik/server"
)

func main() {
	s.Redis()
	s.ConnectDrive()
	e := router.Routing()
	e.Logger.Fatal(e.Start(":8000"))
}
