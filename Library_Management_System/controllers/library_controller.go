package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
)

func StartLibrarySystem(library *services.Library) {

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("8. Reserve Book")
		fmt.Println("9. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter Book ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Book Title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter Book Author: ")
			fmt.Scanln(&author)
			library.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})

		case 2:
			var bookID int
			fmt.Print("Enter Book ID to remove: ")
			fmt.Scanln(&bookID)
			library.RemoveBook(bookID)

		case 3:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println(err)
			}

		case 4:
			var bookID, memberID int
			fmt.Print("Enter Book ID: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println(err)
			}

		case 5:
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}

		case 6:
			var memberID int
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			books, err := library.ListBorrowedBooks(memberID)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s\n", book.ID, book.Title)
			}
		case 7:
			var id int
			var name string
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Member Name: ")
			fmt.Scanln(&name)
			library.AddMember(models.Member{ID: id, Name: name})
			fmt.Println("Member added successfully!")
		case 8:
			var bookID, memberID int
			fmt.Print("Enter Book ID to reserve: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)

			if err := library.ReserveBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book reserved successfully. You have 5 seconds to borrow it.")
			}

		case 9:
			fmt.Println("Exiting Library System.")
			return

		default:
			fmt.Println("Invalid option! Try again.")
		}
	}
}
