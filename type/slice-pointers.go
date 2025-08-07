package main

import "fmt"

func SlicePointers() {
	s := [3]string{"I", "love", "you"}
	fmt.Println("Original slice: ", s)

	p := s[0:2]
	fmt.Println("Slice pointer: ", p)

	p[0] = "hate"
	fmt.Println("Modified slice: ", p)
	fmt.Println("Original after modification: ", s)
}