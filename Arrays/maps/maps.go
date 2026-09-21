package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon"])
	websites["Facebook"] = "https://facebook.com"

	fmt.Println(websites)

	delete(websites, "Facebook")
	fmt.Println(websites)
}
