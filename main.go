package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Books struct {
	Items []BookItem
}

type BookItem struct {
	Name, Author, Date string
	ID                 int
	is_deleted         bool
}

// AddItem appends an instance of bookitem to books struct
func (b *Books) AddItem(item BookItem) []BookItem {
	b.Items = append(b.Items, item)
	return b.Items
}

// FindItem searches a bookitem with a specific name inside Books struct
func (b *Books) FindItem_Name(s string) []BookItem {
	item := []BookItem{}
	for i := range b.Items {
		if strings.Contains(strings.ToLower(b.Items[i].Name), s) {
			item = append(item, b.Items[i])
		}
	}

	return item
}

// FindItem searches a bookitem with an ID
func (b *Books) FindItem_ID(id int) BookItem {
	item := BookItem{}
	for i := range b.Items {
		if b.Items[i].ID == id {
			return b.Items[i]
		}
	}

	return item
}

// FindItem searches a bookitem with an ID
func (b *Books) Deleteitem_ID(id int) {
	for i := range b.Items {
		if b.Items[i].ID == id {
			b.Items[i].is_deleted = true
		}
	}

}

// ListItems lists all book_items with a pretty print format(visualized with empty spaces)
func (b *Books) ListItems() {
	column_length := 50
	empty_string := strings.Repeat("\t", 6)
	fmt.Printf("\tNAME%sAUTHOR%sDATE\n", empty_string, empty_string)
	for index, book := range b.Items {
		if book.is_deleted == false {
			fmt.Printf("%d %s", index+1, book.Name)
			fmt.Printf("%s", strings.Repeat(" ", column_length-len(book.Name)))
			fmt.Printf("%s", book.Author)
			fmt.Printf("%s", strings.Repeat(" ", column_length-len(book.Author)))
			fmt.Printf("%s\n", book.Date)
		}

	}
}
func main() {
	arg := os.Args[1:]

	to_string := strings.Join(arg, " ")
	to_string = strings.ToLower(to_string)

	book_list := map[string][]string{
		"The History of Tom Jones": {"Henry Fielding", "28 February 1749"},
		"Pride and Prejudice":      {"Jane Austen", "28 January 1813"},
		"The Red and the Black":    {"Stendhal", "22 November 1830"},
		"Le Pere Goriot":           {"Honore de Balzac", "8 March 1835 "},
		"David Copperfield":        {"Charles Dickens", "6 November 1850"},
		"Madame Bovary":            {"Gustave Flaubert", "5 April 1857"},
		"Moby-Dick":                {"Herman Melville", "18 October 1851"},
		"Wuthering Heights":        {"Emily Bronte", "3 December 1847"},
		"The Brothers Karamazov":   {"Dostoevsky", "10 January 1879"},
		"War and Peace":            {"Tolstoy", "19 June 1869"},
	}
	book_struct := new(Books)
	count_id := 0
	for k, v := range book_list {
		book := new(BookItem)
		book.Name = k
		book.Author = v[0]
		book.Date = v[1]
		book.ID = count_id
		book.is_deleted = false
		count_id += 1
		book_struct.AddItem(*book)

	}
	if to_string == "list" {
		book_struct.ListItems()
	} else if strings.HasPrefix(to_string, "search ") {
		exp_book := to_string[len("search "):]
		result := book_struct.FindItem_Name(exp_book)
		if len(result) != 0 {
			fmt.Printf("\nThe book has found in the library !\n\n")
			for _, book := range result {
				fmt.Printf("Book Name:%s\nAuthor:%s\nDate:%s\nID:%d\n", book.Name, book.Author, book.Date, book.ID)
			}
		} else {
			fmt.Println("No books found!")
		}
	} else if strings.HasPrefix(to_string, "get ") {
		exp_id, _ := strconv.Atoi(to_string[len("get "):])
		fmt.Println(exp_id)
		result := book_struct.FindItem_ID(exp_id)
		fmt.Println(result)
		if !(result == BookItem{}) {
			fmt.Printf("\nThe book has found in the library !\n\n")

			fmt.Printf("Book Name:%s\nAuthor:%s\nDate:%s\nID:%d\n", result.Name, result.Author, result.Date, result.ID)
		}
	} else if strings.HasPrefix(to_string, "delete ") {
		exp_id, _ := strconv.Atoi(to_string[len("delete "):])
		book_deleted := book_struct.FindItem_ID(exp_id)
		book_struct.Deleteitem_ID(exp_id)
		fmt.Printf("\nThe book has been deleted !\n\n")

		fmt.Printf("Book Name:%s\nAuthor:%s\nDate:%s\nID:%d\n", book_deleted.Name, book_deleted.Author, book_deleted.Date, book_deleted.ID)
		book_struct.ListItems()
	} else {
		fmt.Printf("Invalid input! You must enter 'list' or 'search <bookName>'!")
	}
}
