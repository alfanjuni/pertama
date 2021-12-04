package main

import "fmt"

func main() {
	// fmt.Print("aku sayang golang")

	//ARRAY
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "javascript"
	// languages[2] = "ruby"
	// languages[3] = "php"
	// languages[4] = "java"

	// languages2 := [5]string{"Ruby", "GO", "javascript", "java", "C"}
	// fmt.Println(languages2)

	// for index, lang := range languages {
	// 	fmt.Println("index:", index, "languange:", lang)
	// }

	//SLICE
	// var gamingConsoles []string

	// gamingConsoles = append(gamingConsoles, "PlayStation4")
	// gamingConsoles = append(gamingConsoles, "xbox1")

	// fmt.Println(gamingConsoles)

	//MAP

	// var myMap map[string]int
	// myMap = map[string]int{}
	// myMap["ruby"] = 9
	// myMap["java"] = 5
	// myMap["c"] = 10

	// myMap2 := map[string]string{"ruby": "is beautiful", "java": "awesome"}
	// fmt.Println(myMap["c"])
	// fmt.Println(myMap2)

	// for key, value := range myMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	// //delete map

	// fmt.Println("===============================")

	// delete(myMap, "ruby")
	// for key, value := range myMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	// value, isAvailable := myMap["REACT"]
	// fmt.Println(value, isAvailable)

	students := []map[string]string{
		{"name": "agung", "score": "A"},
		{"name": "hasan", "score": "A"},
		{"name": "cucuk", "score": "C"},
		{"name": "item", "score": "B"},
	}

	fmt.Println(students)

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}

}
