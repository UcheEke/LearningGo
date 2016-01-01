package main

// Solution to the sort package challenge in lecture 119

// Regarding Reverse(data Interface) Interface:
// Reverse() DOES NOT return a sorted slice. It returns an Interface
// which allows sort.Sort to sort the data in reverse order(!!)
// so you end up with the following construct
//
// sort.Sorted(      sort.Reverse(          sort.Interface(datatype))) in order to reverse datatype
//        ^               ^                          ^
// i.e. Sort func -takes- Interface - acting upon- Interface (via type conversion!!)
// Confused.com? I was for a while!! :)

import (
	"fmt"
	"sort"
)

// User defined type 'people' with underlying []string datatype
type people []string

// Implementing my own methods to satisfy the sort interface 'Interface'
// Len(), Swap(i,j int) and Less(i,j int) bool

// Note p is already a reference type (via the underlying []string) so func receiver is of "value" type.
func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	// Check i, j are within bounds
	if (i >= 0 && i < len(p)) && (j >= 0 && j < len(p)) {
		return p[i] < p[j]
	} else {
		return false
	}
}

func (p people) Swap(i, j int) {
	// Check i,j are within bounds
	if (i >= 0 && i < len(p)) && (j >= 0 && j < len(p)) {
		p[i], p[j] = p[j], p[i]
	}
}

// Wrapper functions for basic illustrations
func SortIntSlice(i []int, reverse bool) {
	fmt.Println("Before sorting:", i)
	if reverse {
		fmt.Println("Reverse sort...")
		sort.Sort(sort.Reverse(sort.IntSlice(i)))
	} else {
		sort.Ints(i)
	}
	fmt.Println("After sorting:", i)
}

func SortStrSlice(s []string, reverse bool) {
	fmt.Println("Before sorting:", s)
	if reverse {
		fmt.Println("Reverse sort...")
		sort.Sort(sort.Reverse(sort.StringSlice(s)))
	} else {
		sort.Strings(s)
	}
	fmt.Println("After sorting:", s)
}

func SortPeopleSlice(p people, reverse bool) {
	fmt.Println("Before sorting:", p)
	if reverse {
		fmt.Println("Reverse sort...")
		sort.Reverse(p)
	} else {
		sort.Sort(p)
	}
	fmt.Println("After sorting:", p)
}

func main() {
	// 1. Sorting a user defined type
	fmt.Println("1. User Defined type")
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	SortPeopleSlice(studyGroup, false)

	// 2. Sorting a string slice
	fmt.Println("2. String slice")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	SortStrSlice(s, false)

	// 3. Sorting an int slice
	fmt.Println("3. Int Slice")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	SortIntSlice(n, false)
}
