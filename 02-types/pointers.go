package types

import "fmt"

type StoreItem struct {
	name                string
	isSecurityTagActive bool
}

func activateTag(item *StoreItem) {
	item.isSecurityTagActive = true
}

func deactivateTag(item *StoreItem) {
	item.isSecurityTagActive = false
}

func checkout(items []*StoreItem) {
	for _, item := range items {
		deactivateTag(item)
	}
}

func printItems(items []*StoreItem) {
	fmt.Println("Items:")
	for _, item := range items {
		fmt.Println(*item)
	}
	fmt.Println("\n")
}

func Pointers() {
	tshirt := StoreItem{
		name:                "T-Shirt",
		isSecurityTagActive: true,
	}
	coat := StoreItem{
		name:                "Coat",
		isSecurityTagActive: true,
	}
	onesie := StoreItem{
		name:                "Onesie",
		isSecurityTagActive: true,
	}
	cap := StoreItem{
		name:                "Cap",
		isSecurityTagActive: true,
	}

	items := []*StoreItem{&tshirt, &coat, &onesie, &cap}
	printItems(items)

	deactivateTag(&cap)
	printItems(items)

	checkout(items)
	printItems(items)

}
