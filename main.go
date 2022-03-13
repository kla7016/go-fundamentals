package main

import (
	"fmt"
	"fundamentals/book"
)

func main() {
	a := book.Book{
		Name: "Harry Potter",
		Author: "J. K. Rowling",
	}

	fmt.Println(a.String())
	fmt.Println(a.Set("Korkla eiei"))
}

