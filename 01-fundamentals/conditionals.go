package fundamentals

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func processAccess(role, day int) string {

	granted := "Access Granted"
	denied := "Access Denied"

	if role <= Manager {
		return granted
	}
	if role == Contractor {
		if day > Friday {
			return granted
		}
	}
	if role == Member {
		if day <= Friday {
			return granted
		}
	}
	if role == Guest {
		if day == Monday || day == Wednesday || day == Friday {
			return granted
		}
	}

	return denied
}

func Conditionals() {
	fmt.Println(processAccess(Manager, Wednesday))
	fmt.Println(processAccess(Guest, Wednesday))
	fmt.Println(processAccess(Contractor, Saturday))
	fmt.Println(processAccess(Contractor, Friday))
	fmt.Println(processAccess(Admin, Monday))
	fmt.Println(processAccess(Guest, Thursday))
}
