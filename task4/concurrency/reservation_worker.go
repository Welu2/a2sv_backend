package concurrency

import (
	"fmt"
	"library_management/services"
	"time"
)

type ReservationRequest struct {
	BookID   int
	MemberID int
}

func StartReservationWorker(
	service services.LibraryManager,
	requestChannel chan ReservationRequest,
) {

	for request := range requestChannel {

		go handleReservation(service, request)
	}
}

func handleReservation(service services.LibraryManager, req ReservationRequest) {

	err := service.ReserveBook(req.BookID, req.MemberID)

	if err != nil {
		fmt.Println("Reservation failed:", err)
		return
	}

	fmt.Println("Book reserved:", req.BookID, "by member:", req.MemberID)

	go func() {

		time.Sleep(5 * time.Second)

		fmt.Println("Reservation expired for book:", req.BookID)

	}()
}