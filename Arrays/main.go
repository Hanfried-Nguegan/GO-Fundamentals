package main

import "fmt"

type FloatMap map[string]float64

func (m FloatMap) output() {
	fmt.Println(m)
}

func main() {
	// userNames := []string{}
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Hanfried")

	fmt.Println(userNames)

	courseRatings := make(FloatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.5

	courseRatings.output()

	//fmt.Println(courseRatings)

	for index, value := range userNames {
		// ...
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}

}
