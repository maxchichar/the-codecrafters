package main

import (
	"fmt"
)
var buffer [256]byte

func AddOneToEachElement(slice []byte)  {
	for i := range slice{
		slice[i]++
	}
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice) - 1]
	return slice
}

func main()  {
	//AddOneToEachElement
	slice := buffer[10:60]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("Add One To Each Element")
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	fmt.Println()

	//SubtractOneFromLength
	fmt.Println("Subtract One From Length")
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After: len(slice) =", len(slice))
	fmt.Println("After: len(newSlice) =", len(newSlice))
}