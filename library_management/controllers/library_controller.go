package controllers

import (
	"fmt"
	"main/models"
	"main/services"
	"strings"

)

type LibraryController struct {
	service services.LibraryService
}


func NewLibraryController() *LibraryController {
	return &LibraryController{
		service: *services.NewLibraryService(),
	}
}

func (c *LibraryController) AddBook(reader models.Book) {
	
	c.service.AddBook(reader)
	fmt.Println("Book added successfully.")
}

func (c *LibraryController) RemoveBook(rm_id int) {
	
	c.service.RemoveBook(rm_id)
	fmt.Println("Book removed successfully.")
}

func (c *LibraryController) BorrowBook(bookID int, memberID int) {	
	

	err := c.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func (c *LibraryController) ReturnBook(bookID int, memberID int) {	
	err := c.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func (c *LibraryController) ListAvailableBooks() {
	books := c.service.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) ListBorrowedBooks(memberID int) {	
	
	books := c.service.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (c *LibraryController) Conversaion() {

	var choice int
	var rm_id int
	
	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

	

		switch choice {
		case 1:
			var title string
			var id int
			var author string

			fmt.Print("Enter Book ID: ")
			fmt.Scanln(&id)

			fmt.Print("Enter Book Title: ")
			fmt.Scanln(&title)

			fmt.Print("Enter Book Author: ")
			fmt.Scanln(&author)

			book := models.Book{
				ID:     id,
				Title:  strings.TrimSpace(title),
				Author: strings.TrimSpace(author),
				Status: "Available",
			}
			c.AddBook(book)
		case 2:
			fmt.Print("Enter Book ID to remove: ")
			fmt.Scanln(&rm_id)
			c.RemoveBook(rm_id)
			
		case 3:
			var bookID int
			var memberID int
			fmt.Print("Enter Book ID to borrow: ")
			fmt.Scanln(&bookID)
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)
			c.BorrowBook(bookID, memberID)
		case 4:
			var bookID int
			var memberID int
			fmt.Print("Enter Book ID to return: ")
	
			fmt.Scanln(&bookID)

			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)	
		
			c.ReturnBook(bookID, memberID)
		case 5:
			c.ListAvailableBooks()
		case 6:
			var memberID int
			fmt.Print("Enter Member ID: ")
			fmt.Scanln(&memberID)

			c.ListBorrowedBooks(memberID)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	} 
}
