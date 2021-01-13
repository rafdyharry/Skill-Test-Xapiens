package main

import (
	"fmt"
	"sort"
)

// bikin fungsi untuk penjumlahan
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	// input array
	array := []int{1, 2, 33, 44, 55, 33, 44, 22, 44, 33, 2, 77, 66, 1, 2, 3, 4, 5, 6, 7, 89, 3, 3, 8, 9, 75, 4, 3, 2, 2, 1, 3}

	// bagi array sesuai pembagiannya

	var array_1 = array[0:10]

	array_2 := array[10:20]

	array_3 := array[20:32]

	// menampilkan hasil pembagian aray
	fmt.Println("Pembagian Array Menjadi 3: ")
	println("Hasil pembagian anggota array_1: 0-", len(array)/3)
	fmt.Println("New Array 1: ", array_1)
	println("Hasil pembagian anggota array_2: 10-", len(array)/3+10)
	fmt.Println("New Array 2: ", array_2)
	println("Hasil pembagian anggota array_3 di tambah mod hasil bagi array: 20-", len(array)/3+10+2)
	fmt.Println("New Array 3: ", array_3)
	// Mengurutkan angka terbesar ke kecil
	sort.Sort(sort.Reverse(sort.IntSlice(array_1)))
	sort.Sort(sort.Reverse(sort.IntSlice(array_2)))
	sort.Sort(sort.Reverse(sort.IntSlice(array_3)))

	// pengurutan angka
	fmt.Println(" ")
	fmt.Println("Pengurutan Array Dari Besar ke Kecil: ")
	fmt.Println("Array 1: ", array_1)
	fmt.Println("Array 2: ", array_2)
	fmt.Println("Array 3: ", array_3)

	//SUM
	fmt.Println(" ")
	fmt.Println("Jumlah Array :")
	fmt.Println("Jumalah Array 1 :", sum(array_1))
	fmt.Println("Jumlah Array 2 :", sum(array_2))
	fmt.Println("Jumlah Array 3 :", sum(array_3))
	//mean
	fmt.Println(" ")
	fmt.Println("Rata-rata Array :")
	fmt.Println("Rata-rata Array 1 :", sum(array_1)/len(array_1))
	fmt.Println("Rata-rata Array 2 :", sum(array_2)/len(array_2))
	fmt.Println("Rata-rata Array 3 :", sum(array_3)/len(array_3))

	// Min Max

	min := array_1[0]
	for _, v := range array_1 {
		if v < min {
			min = v
		}
	}
	fmt.Println("Nilai Manimal Pada array_1: ", min)
	max := array_1[0]
	for _, v := range array_1 {
		if v > max {
			max = v
		}
	}
	fmt.Println("Nilai Maximal pada array_1: ", max)
	fmt.Println("")

	min1 := array_2[0]
	for _, v := range array_2 {
		if v < min1 {
			min1 = v
		}
	}
	fmt.Println("Nilai Manimal Pada array_2: ", min1)
	max1 := array_2[0]
	for _, v := range array_2 {
		if v > max1 {
			max1 = v
		}
	}
	fmt.Println("Nilai Maximal pada array_2: ", max1)
	fmt.Println(" ")

	min2 := array_3[0]
	for _, v := range array_3 {
		if v < min2 {
			min2 = v
		}
	}
	fmt.Println("Nilai Manimal Pada array_3: ", min2)
	max2 := array_3[0]
	for _, v := range array_3 {
		if v > max2 {
			max2 = v
		}
	}
	fmt.Println("Nilai Maximal pada array_1: ", max2)
}
