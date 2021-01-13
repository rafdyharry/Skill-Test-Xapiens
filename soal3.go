package main

import (
	"fmt"
	"strings"
)

func main() {

	// bikin kalimat yang ingin di oprasikan
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio."

	fmt.Println("String 1: ", str)

	// bikin karakter yang ingin di cari
	res1 := strings.Count(str, "a")
	res2 := strings.Count(str, "b")
	res3 := strings.Count(str, "c")
	res4 := strings.Count(str, "z")

	// tampilkan hasil pencarian
	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)

}
