package main

import "categoryht"

func main() {
	router := categoryht.SetupRouter()
	router.Run(":8080")
}
