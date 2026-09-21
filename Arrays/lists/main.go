package lists

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 11}
	fmt.Println(prices[0:1])
	prices[1] = 16.99

	prices = append(prices, 5.99, 6.99, 20.0)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{100.99, 200.99, 300.99}

	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

// func main() {
// 	var productNames [4]string = [4]string{"A Book", "A Chair"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"
// 	productNames[3] = "A Game"

// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 100.99

// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
