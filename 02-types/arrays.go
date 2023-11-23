package types

import "fmt"

type Product struct {
	price int
	name  string
}

func totalCost(shoppingList [4]Product) int {
	total := 0
	for i := 0; i < len(shoppingList); i++ {
		shoppingItem := shoppingList[i]
		total += shoppingItem.price
	}

	return total
}

func Arrays() {
	products := [4]Product{
		{
			price: 60,
			name:  "Bimilk 2.8%",
		},
		{
			price: 79,
			name:  "Bimilk 0.9%",
		},
		{
			price: 11,
			name:  "Choko biskvit",
		},
		{
			price: 59,
			name:  "Zito Luks miks od zrna",
		},
	}

	fmt.Println("The total price is ", totalCost(products), "MKD")

}
