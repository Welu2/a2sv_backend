package controllers

import (
	"fmt"
	"library_management/concurrency"
	"library_management/services"
)

type LibraryController struct {
	Service         services.LibraryManager
	ReservationChan chan concurrency.ReservationRequest
}

func (c *LibraryController) Start() {

	for {

		fmt.Println("\nLibrary Menu")
		fmt.Println("1 Add Book")
		fmt.Println("2 Reserve Book")
		fmt.Println("3 List Available Books")
		fmt.Println("0 Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var id int
			var title string
			var author string

			fmt.Println("ID:")
			fmt.Scanln(&id)

			fmt.Println("Title:")
			fmt.Scanln(&title)

			fmt.Println("Author:")
			fmt.Scanln(&author)

			c.Service.AddBook(
				struct {
					ID     int
					Title  string
					Author string
					Status string
					Reserved bool
				}{
					ID: id,
					Title: title,
					Author: author,
				},
			)

		case 2:

			var bookID int
			var memberID int

			fmt.Println("Book ID:")
			fmt.Scanln(&bookID)

			fmt.Println("Member ID:")
			fmt.Scanln(&memberID)

			c.ReservationChan <- concurrency.ReservationRequest{
				BookID:   bookID,
				MemberID: memberID,
			}

		case 3:

			books := c.Service.ListAvailableBooks()

			for _, b := range books {
				fmt.Println(b.ID, b.Title)
			}

		case 0:
			return
		}
	}
}