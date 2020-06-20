package main

import "locationapi/http"

func main() {
	httpServer := http.NewServer()
	httpServer.Open()
}
