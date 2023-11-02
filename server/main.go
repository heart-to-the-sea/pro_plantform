package main

import (
	"server/router"
)

func main() {
	server := router.GetRouter()
	server.Run(":9999")
}
