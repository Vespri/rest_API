package main

import "api_go/routers"

func main() {
	var PORT = ":2121"

	routers.StartServer().Run(PORT)
}
