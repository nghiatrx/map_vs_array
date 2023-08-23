package main

import (
	"testing"
)

func benchMap(myMap *map[int]int, index int) int {
	return (*myMap)[index]
}

func benchArray(myArray *[]int, index int) int {
	return (*myArray)[index]
}

func Benchmark1M(b *testing.B) {
	num := 1000000
	myMap := make(map[int]int, num)
	myArray := make([]int, num)

	for i := 0; i < num; i++ {
		myMap[i] = 1
		myArray[i] = i
	}

	b.Run("map", func(b *testing.B) {
		benchMap(&myMap, num-1)
	})
	b.Run("array", func(b *testing.B) {
		benchArray(&myArray, num-1)
	})
}

func Benchmark50M(b *testing.B) {
	num := 50000000
	myMap := make(map[int]int, num)
	myArray := make([]int, num)

	for i := 0; i < num; i++ {
		myMap[i] = 1
		myArray[i] = i
	}

	b.Run("map", func(b *testing.B) {
		benchMap(&myMap, num-1)
	})
	b.Run("array", func(b *testing.B) {
		benchArray(&myArray, num-1)
	})
}

func Benchmark100M(b *testing.B) {
	num := 100000000
	myMap := make(map[int]int, num)
	myArray := make([]int, num)

	for i := 0; i < num; i++ {
		myMap[i] = 1
		myArray[i] = i
	}

	b.Run("map", func(b *testing.B) {
		benchMap(&myMap, num-1)
	})
	b.Run("array", func(b *testing.B) {
		benchArray(&myArray, num-1)
	})
}
