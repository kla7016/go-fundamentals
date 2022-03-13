package book

import "fmt"

type Book struct {
	Name   string
	Author  string
}
//** วิธีทดสอบ
//a := book.Book{
//Name: "Harry Potter",
//Author: "J. K. Rowling",
//}
//
//fmt.Println(a.String())
//fmt.Println(a.Set("Korkla eiei"))

func (bookData Book) String() string {
	return fmt.Sprintf("%s by %s", bookData.Name, bookData.Author)
}

func (bookData Book) Set(name string) Book {
	 bookData.Name = name
	 return bookData
}