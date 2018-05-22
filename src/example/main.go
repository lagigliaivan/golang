package main

import "example/categoryht"

func main() {
	router := categoryht.SetupRouter()
	router.Run(":8080")
}
