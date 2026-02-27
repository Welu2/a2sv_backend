package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {

	library := services.NewLibrary()

	controller := controllers.LibraryController{
		Service: library,
	}

	controller.Start()
}