# Library Management System Documentation

## Overview

This is a simple console-based library management system implemented in Go. It allows you to manage books and members, including adding and removing books, borrowing and returning books, and listing available and borrowed books.

## Structs

### Book

- ID (int)
- Title (string)
- Author (string)
- Status (string)

### Member

- ID (int)
- Name (string)
- BorrowedBooks ([]Book)

## Interfaces

### LibraryManager

- AddBook(book Book)
- RemoveBook(bookID int)
- BorrowBook(bookID int, memberID int) error
- ReturnBook(bookID int, memberID int) error
- ListAvailableBooks() []Book
- ListBorrowedBooks(memberID int) []Book

## Implementation

### LibraryService

Implements the `LibraryManager` interface.

- `AddBook(book models.Book)`
- `RemoveBook(bookID int)`
- `BorrowBook(bookID int, memberID int) error`
- `ReturnBook(bookID int, memberID int) error`
- `ListAvailableBooks() []models.Book`
- `ListBorrowedBooks(memberID int) []models.Book`

## Console Interaction

### main.go

The entry point of the application. It provides a console-based menu for interacting with the library management system.

## Folder Structure
library_management/
├── main.go
├── controllers/
│ └── library_controller.go
├── models/
│ └── book.go
│ └── member.go
├── services/
│ └── library_service.go
├── docs/
│ └── documentation.md
└── go.mod


## Usage

1. Run the application: `go run main.go`
2. Follow the on-screen instructions to manage books and members.

## Error Handling

- Ensures appropriate error handling for scenarios where books or members are not found, or books are already borrowed.
- Prints error messages to the console for invalid operations.

## Dependencies

- No external dependencies.



