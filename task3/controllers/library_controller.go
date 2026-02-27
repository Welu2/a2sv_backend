package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

type LibraryController struct {
	Service services.LibraryManager
}

func (c *LibraryController) Start() {

	for {
		fmt.Println("\n===== Library Menu =====")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Member Borrowed Books")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var id int
			var title, author string

			fmt.Print("Book ID: ")
			fmt.Scanln(&id)

			fmt.Print("Title: ")
			fmt.Scanln(&title)

			fmt.Print("Author: ")
			fmt.Scanln(&author)

			c.Service.AddBook(models.Book{
				ID:     id,
				Title:  title,
				Author: author,
			})

		case 2:
			var id int
			fmt.Print("Book ID: ")
			fmt.Scanln(&id)
			c.Service.RemoveBook(id)

		case 3:
			var bookID, memberID int
			fmt.Print("Book ID: ")
			fmt.Scanln(&bookID)

			fmt.Print("Member ID: ")
			fmt.Scanln(&memberID)

			err := c.Service.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println(err)
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Book ID: ")
			fmt.Scanln(&bookID)

			fmt.Print("Member ID: ")
			fmt.Scanln(&memberID)

			err := c.Service.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println(err)
			}

		case 5:
			books := c.Service.ListAvailableBooks()
			for _, b := range books {
				fmt.Println(b.ID, "-", b.Title, "-", b.Author)
			}

		case 6:
			var memberID int
			fmt.Print("Member ID: ")
			fmt.Scanln(&memberID)

			books := c.Service.ListBorrowedBooks(memberID)
			for _, b := range books {
				fmt.Println(b.ID, "-", b.Title)
			}

		case 0:
			fmt.Println("Goodbye!")
			return
		}
	}
}