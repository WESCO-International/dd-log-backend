package myslices

import (
	"errors"
	"fmt"
)

// 1 copy one slice to another
func CreatingSlice() []uint {
	numSlice := []uint{1, 2, 3}
	numSlice2 := make([]uint, 3, 10)
	numSlice = append(numSlice, 4, 5, 6)
	numSlice2 = numSlice[:]
	//dummy1 := []uint{}
	dummy1 := make([]uint, len(numSlice2))
	copy(dummy1, numSlice2[:])
	//fmt.Println(myappend(dummy1, 5))
	//fmt.Println("########")
	shalowCopy()
	return numSlice2
}

func myappend(mylist []uint, count uint) []uint {
	myvar1 := make([]uint, 0)
	var i uint
	for i = 0; i < count; i++ {
		myvar1 = append(myvar1, mylist[i])
	}
	return myvar1
}

func shalowCopy() {
	s1 := make([]int, 3, 6)
	s2 := s1[1:3] //only 1st and 2nd (0,0), as a part of list.
	//s1 = append(s1, 40)
	//fmt.Printf("%p %p\n", &s1, &s2)
	s2 = append(s2, 400)
	//fmt.Println(s1, s2)
	s1 = append(s1, 30, 50)
	//fmt.Println(s1, s2)
	s2 = append(s2, 500, 600)
	//fmt.Println(s1, s2, len(s1), cap(s1), len(s2), cap(s2))
	s2 = append(s2, 100)
	//fmt.Println(s1, s2, len(s1), cap(s1), len(s2), cap(s2))
}

func ReadSlicesIndex(i []uint, index int) {
	if index >= len(i) {
		err := errors.New("index out of range while reading the slice")
		fmt.Println(err)
		//	err.error.Error("out of range")
	} else {
		fmt.Printf("value of index %v is %v\n", index, i[index])
	}
}
func IterateSlices(myslice []uint) {
	for _, i := range myslice {
		fmt.Println(i)
	}
}
func UpdateSlices(i []uint, index int, value uint) ([]uint, error) {
	if index >= len(i) {
		err := errors.New("index out of range while updating")
		return nil, err
	} else {
		i[index] = value
		return i, nil
	}
}

func DeleteSlicesElementUsingReslicing(i []uint, index int) ([]uint, error) {
	if index >= len(i) {
		err := errors.New("index out of range")
		return nil, err
	} else {
		newReslice := append(i[:index], i[index+1:]...)
		return newReslice, nil
	}
}

// basic way but it will alter the order of list.
func DeleteSlicesElementUsingSwapShift(i []uint, index int) ([]uint, error) {

	if index >= len(i) {
		err := errors.New("index out of range in delete slices element using swapshift")
		return nil, err
	} else {
		lastElement := i[len(i)-1] //get last element
		i[index] = lastElement
		i = i[:len(i)-1]
		return i, nil
	}
}

func DeleteSlicesElementUsingCopy(i []uint, index int) ([]uint, error) {
	if index >= len(i) {
		err := errors.New("index out of range in Delete Slices Element Using Copy")
		return nil, err
	} else {
		newslice := make([]uint, len(i)-1)  //newslice-[0,0,0,0],i-1,2,3,4,5 //index-2
		copy(newslice[:index], i[:index])   //newslice-[1,2,0,0]
		copy(newslice[index:], i[index+1:]) //newslice-[1,2,4,5]
		return newslice, nil
	}
}

func ReverseSlice(i []uint) ([]uint, error) {
	newslice := make([]uint, len(i))
	reverseIndex := len(i) - 1
	for element := 0; element < len(i); element++ {
		newslice[reverseIndex] = i[element]
		reverseIndex -= 1
	}
	return newslice, nil
}

func SearchElementsFirstAppearance(i []uint, searchElement uint) (uint, error) {
	for x, y := range i {
		if y == searchElement {
			return uint(x), nil
		}
	}
	err := errors.New("this element doesn't exists")
	var nilVar uint
	return nilVar, err
}

//O-logN reverse and sort implement and search element's first appearance, if not present msg:- number is not available.
//remove duplicates from slice
