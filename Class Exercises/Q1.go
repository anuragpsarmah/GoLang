/*
You're developing an online store application in GoLang.
As part of the application, you need to keep track of various product details such as name, price, and quantity in stock.
Design a set of variables and assign values to represent a specific product in the inventory.
Ensure you use appropriate data types for each variable to accurately capture the information.
*/

package main

import (
	"fmt"
)

func main() {
	// I have defined a struct for product.
	type Product struct {
		Name     string
		Price    float64
		Quantity int
	}

	// I have create an empty slice to store objects of Product Structute. I used slice because it is of variable size.
	var products []Product

	// Here I am getting the number of products.
	var numProducts int
	fmt.Print("Enter the number of products: ")
	fmt.Scanln(&numProducts)

	// Here I am running a loop to get the product details.
	for i := 0; i < numProducts; i++ {
		var productName string
		var productPrice float64
		var productQuantity int

		fmt.Printf("\nEnter details for Product %d:\n", i+1)

		fmt.Print("Name: ")
		fmt.Scanln(&productName)

		fmt.Print("Price: ")
		fmt.Scanln(&productPrice)

		fmt.Print("Quantity: ")
		fmt.Scanln(&productQuantity)

		// I creating a new Product object and appending it to the slice.
		product := Product{Name: productName, Price: productPrice, Quantity: productQuantity}
		products = append(products, product)
	}

	// I am finally printing all the products in the inventory.
	fmt.Println("\nProducts in Inventory:")
	for i, product := range products {
		fmt.Printf("Product %d:\n", i+1)
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Quantity in Stock:", product.Quantity)
		fmt.Println()
	}
}
