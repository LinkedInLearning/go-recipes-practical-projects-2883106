package main

import "fmt"

// isbn -> book
var booksDB = map[string]Book{
	"0062225677": {
		Title:  "The Colour of Magic",
		Author: "Terry Pratchett",
		ISBN:   "0062225677",
	},
	"0765394855": {
		Title:  "Old Man's War",
		Author: "John Scalzi",
		ISBN:   "0765394855",
	},
}

func getBook(isbn string) (Book, error) {
	book, ok := booksDB[isbn]
	if !ok {
		return Book{}, fmt.Errorf("unknown ISBN: %q", isbn)
	}

	return book, nil
}
