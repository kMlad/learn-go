package types

import (
	"fmt"
	. "time"
)

type Book struct {
	title       string
	isAvailable bool
}

type LibraryEvent struct {
	event string
	time  Time
}

type Library struct {
	name           string
	books          []Book
	members        map[string][]Book
	lendingHistory []LibraryEvent
}

func checkOut(bookTitle string, memberName string, library *Library) {
	var book *Book
	bookExists := false
	memberExists := false

	// Change isAvailable to false
	for i := range library.books {
		book = &library.books[i]
		if book.title == bookTitle && book.isAvailable {
			bookExists = true
			book.isAvailable = false
			break
		}
	}
	if !bookExists {
		panic("Error checking out: No such book exists!")
	}

	// Add book to member's  list
	if library.members[memberName] != nil {
		memberExists = true
		library.members[memberName] = append(library.members[memberName], *book)
	}
	if !memberExists {
		panic("Error checking out: No such member exists!")
	}

	// Log event in lending history
	event := fmt.Sprintf("%s checked out %s", memberName, bookTitle)
	library.lendingHistory = append(library.lendingHistory, LibraryEvent{event, Now()})
}

func checkIn(bookTitle string, memberName string, library *Library) {
	var book *Book
	bookIndex := -1
	bookInSystem := false
	memberHasBook := false

	// Check if book is in system
	for i := range library.books {
		book = &library.books[i]
		if book.title == bookTitle {
			bookInSystem = true
			break
		}
	}
	if !bookInSystem {
		panic("Error checking in: Book is not registered in system!")
	}

	// Check if user has book
	if library.members[memberName] != nil {
		for i, book := range library.members[memberName] {
			if book.title == bookTitle {
				memberHasBook = true
				bookIndex = i
				break
			}
		}
	}
	if !memberHasBook {
		panic("Error checking in: This member is not in possesion of the book!")
	}

	// Remove book from member and Change isAvailable to True
	book.isAvailable = true
	library.members[memberName] = append(library.members[memberName][:bookIndex], library.members[memberName][bookIndex+1:]...)

	// Log change
	event := fmt.Sprintf("%s checked out %s", memberName, bookTitle)
	library.lendingHistory = append(library.lendingHistory, LibraryEvent{event, Now()})
}

func printWorkingDay(library *Library) {
	for _, event := range library.lendingHistory {
		fmt.Printf("%s at %s\n", event.event, event.time.Local().Format("2006-01-02 15:04:05"))
	}
}

func SectionReviewLibrary() {
	skopjeLibrary := Library{
		name: "Skopje Library",
		books: []Book{
			{"Atomic Habits", true},
			{"Meditations", true},
			{"The Book of Five Rings", true},
			{"Endure", true},
		},
		members: map[string][]Book{
			"Peter": make([]Book, 0, 5),
			"Paul":  make([]Book, 0, 5),
			"John":  make([]Book, 0, 5),
			"Mary":  make([]Book, 0, 5),
		},
		lendingHistory: []LibraryEvent{{"Opening Hours", Now()}},
	}

	checkOut("Atomic Habits", "Peter", &skopjeLibrary)
	checkOut("Endure", "Paul", &skopjeLibrary)
	checkIn("Atomic Habits", "Peter", &skopjeLibrary)
	checkOut("Meditations", "Peter", &skopjeLibrary)
	printWorkingDay(&skopjeLibrary)
	// fmt.Println(skopjeLibrary)
}
