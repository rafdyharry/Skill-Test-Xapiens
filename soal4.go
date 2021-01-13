// input package main
package main

// input import yang di butuhkan
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// bikin random array 1-100
	rand.Seed(time.Now().UnixNano())

	v := rand.Perm(100)
	// bikin inputan oleh user
	fmt.Println(v)
	fmt.Println("angka yang ingin di cari: ")
	var a int
	fmt.Scanln(&a)
	res := sort.SearchInts(v, a)

	// Displaying the results
	fmt.Println("Angka berada pada urutan ke : ", res)

}
