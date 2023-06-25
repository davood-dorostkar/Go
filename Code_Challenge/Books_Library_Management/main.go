package library

import (
	"fmt"
)
// import (
// 	"fmt"
// 	"strings"
// )

type Book struct {
	name string;
	borrowed bool;
	person string;
}

type Library struct {
	books []*Book
}

func toLower(s string) string {
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}

func exists(lib *Library, book string)(status bool,index int){
	book = toLower(book)
	for index,b := range(lib.books){
		if book == b.name{
			return true,index;
		}
	}
	return false,-1;
}

func NewLibrary(capacity int) *Library {
	books := make([]*Book, capacity)
	for i := 0; i < capacity; i++ {
		books[i] = &Book{borrowed: false}
	}
	return &Library{books: books}
}

func (library *Library) AddBook(name string) string {
	name = toLower(name);
	status, _ := exists(library,name);
	if status{
		return "The book is already in the library"
	} else {
		book := Book{
			name:     name,
			borrowed: false,
			person: "",
		}
		library.books = append(library.books,&book)
		return "OK"
	}
}

func (library *Library) BorrowBook(bookName, personName string) string {
	bookName = toLower(bookName);
	status, index := exists(library,bookName);
	if status{
		borrowed := library.books[index].borrowed
		if borrowed{
			person := library.books[index].person
			return fmt.Sprintf("The book is already borrowed by, %s!", person)
		}
		library.books[index].borrowed = true;
		library.books[index].person = personName;
		return "OK"
	} else {
		return "The book is not defined in the library"
	}
}

func (library *Library) ReturnBook(bookName string) string {
	bookName = toLower(bookName);
	status, index := exists(library,bookName);
	if status{
		borrowed := library.books[index].borrowed
		if borrowed{
			library.books[index].borrowed = false;
			return "OK";
		}
		return "The book has not been borrowed"
	} else {
		return "The book is not defined in the library"
	}
}