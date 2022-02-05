package main

import (
	"fmt"
	"sort"
	"sync"
)


func GetSliceLength() int {
	var length int
	fmt.Println("Enter the number of elements in array (minimum 4 are reqired):")
	for length<4 {
		fmt.Scan(&length)
		if length<4 {
			fmt.Println("The minimum number is 4.")
			fmt.Println("Please, enter the new number:")
		}
	}
	fmt.Println("")
	return length
}


func FillSlice(sli []int) {
	var number int
	for i:=0; i<len(sli); i++ {
		fmt.Printf("Enter element %d: ", i)
		fmt.Scan(&number)
		sli[i] = number
	}
	fmt.Println("\nThe unsorted array is:")
	fmt.Println(sli, "\n")
}


func GetSepIndecies(sliLength int) []int {
	chunkLength := sliLength  / 4.0
	splits := make([]int, 5)
	splits[0] = 0
	for i:=1; i<4; i++ {
		splits[i] = splits[i - 1] + chunkLength
	}
	splits[4] = sliLength
	return splits
}


func SortInts(sli []int, wg *sync.WaitGroup) {
	sort.Ints(sli)
	wg.Done()
}


func SortChunks(sli, sepIndicies []int) {
	var wg sync.WaitGroup
	wg.Add(4)
	for i:=0; i<4; i++ {
		s := sepIndicies[i]
		e := sepIndicies[i+1]
		go SortInts(sli[s:e], &wg)
	}
	wg.Wait()
}


func Merge(sli1, sli2 []int) []int {
	l1 := len(sli1)
	l2 := len(sli2)
	sliOut := make([]int, l1+l2)
	
	i, j := 0, 0
	for index := 0; index < l1+l2; index++ {

		if i == l1 {
			sliOut[index] = sli2[j]
			j++
		} else if j == l2 {
			sliOut[index] = sli1[i]
			i++
		} else if sli1[i] < sli2[j] && i < l1 {
			sliOut[index] = sli1[i]
			i++
		} else if sli1[i] >= sli2[j] && j < l2 {
			sliOut[index] = sli2[j]
			j++
		}
	}
	return sliOut
}


func MergeChunks(sli, sepIndicies []int) []int {
	outSli := sli[0:sepIndicies[1]]
	for i:=1; i<4; i++ {
		m := sepIndicies[i]
		e := sepIndicies[i+1]
		outSli = Merge(outSli, sli[m:e])
	}
	return outSli
}


func main() {
	sliLength := GetSliceLength()
	sli := make([]int, sliLength)
	sepIndicies := GetSepIndecies(sliLength)
	FillSlice(sli)
	SortChunks(sli, sepIndicies)
	sli = MergeChunks(sli, sepIndicies)
	fmt.Println("Sorted array:")
	fmt.Println(sli)
}
