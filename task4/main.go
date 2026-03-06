package main

import (
	"library_management/controllers"
	"library_management/concurrency"
	"library_management/services"
)

func main() {

	library := services.NewLibrary()

	reservationChannel := make(chan concurrency.ReservationRequest)

	go concurrency.StartReservationWorker(
		library,
		reservationChannel,
	)

	controller := controllers.LibraryController{
		Service:         library,
		ReservationChan: reservationChannel,
	}

	controller.Start()
}