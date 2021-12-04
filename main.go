package main

import "fmt"

func main() {
	// fmt.Print("aku sayang golang")

	//ARRAY
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "javascript"
	languages[2] = "ruby"
	languages[3] = "php"
	languages[4] = "java"

	languages2 := [5]string{"Ruby", "GO", "javascript", "java", "C"}
	fmt.Println(languages2)

	for index, lang := range languages {
		fmt.Println("index:", index, "languange:", lang)
	}
}
