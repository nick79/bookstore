package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	var result []Book
	for _, b := range catalog {
		result = append(result, b)
	}

	return result
}

func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]

	if !ok {
		return Book{}, fmt.Errorf("ID %d doens't exist", ID)
	}

	return b, nil
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}
