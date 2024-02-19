package myslices

import "fmt"

// 1 copy one slice to another
func CreatingSlice() {
	numSlice := []uint{1, 2, 3}
	fmt.Println(numSlice)
	numSlice2 := make([]uint, 3, 10)
	fmt.Println(numSlice2)
	numSlice = append(numSlice, 4, 5, 6)
	fmt.Println(numSlice)
	numSlice2 = numSlice[:]
	fmt.Println(numSlice2)
	//dummy1 := []uint{}
	dummy1 := make([]uint, len(numSlice2))
	copy(dummy1, numSlice2[:])
	fmt.Println(dummy1)
	fmt.Println(myappend(dummy1, 5))
	fmt.Println("########")
	testCopy()
}

func myappend(mylist []uint, count uint) []uint {
	myvar1 := make([]uint, 0)
	var i uint
	for i = 0; i < count; i++ {
		myvar1 = append(myvar1, mylist[i])
	}
	return myvar1
}

func testCopy() {
	s1 := make([]int, 3, 6)
	s2 := s1[1:3] //only 1st and 2nd (0,0), as a part of list.
	s1 = append(s1, 40)
	//fmt.Printf("%p %p\n", &s1, &s2)
	s2 = append(s2, 400)
	fmt.Println(s1, s2)
	s1 = append(s1, 30, 50)
	fmt.Println(s1, s2)
	s2 = append(s2, 500, 600)
	fmt.Println(s1, s2, len(s1), cap(s1), len(s2), cap(s2))
	s2 = append(s2, 100)
	fmt.Println(s1, s2, len(s1), cap(s1), len(s2), cap(s2))
}

func 