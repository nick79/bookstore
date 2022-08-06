package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1

	result, err := bookstore.Buy(b)

	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies

	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of 2, got %d", want, got)
	}

}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("want error buying book when zero copies left, but got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{
			ID:    1,
			Title: "For the Love of Go",
		},
		{
			ID:    2,
			Title: "The Power of Go: Tools",
		},
	}
	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go: Tools",
	}

	got := bookstore.GetBook(catalog, 2)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
