package main

import (
	"fmt"
	"go-mistakes/mymap"
	"go-mistakes/myslices"
)

func main() {
	newlist := myslices.CreatingSlice()
	myslices.ReadSlicesIndex(newlist, 3)
	myslices.IterateSlices(newlist)

	var updatedSlice []uint
	fmt.Printf("before updateSlice  %v\n", newlist)
	updatedSlice, err := myslices.UpdateSlices(newlist, 2, 55)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("after updateSlice %v\n", updatedSlice)
	}
	//fmt.Println(newlist) //actual changing in underlying array.
	fmt.Printf("after delete index %v\n", newlist)
	updatedSlice, err = myslices.DeleteSlicesElementUsingReslicing(newlist, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("after delete index %v\n", updatedSlice)
	}
	fmt.Printf("after delete index using swap and shift %v\n", newlist)
	updatedSlice, err = myslices.DeleteSlicesElementUsingSwapShift(newlist, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("after delete index using swap and shift %v\n", updatedSlice)
	}

	fmt.Printf("after delete index before copy %v\n", updatedSlice)
	updatedSlice, err = myslices.DeleteSlicesElementUsingCopy(updatedSlice, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("after delete index using copy %v\n", updatedSlice)
	}
	fmt.Printf("before reverse: %d\n", updatedSlice)
	updatedSlice, _ = myslices.ReverseSlice(updatedSlice)
	fmt.Printf("afer reverse: %d\n", updatedSlice)

	elementFirstAppearance, err := myslices.SearchElementsFirstAppearance(updatedSlice, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("first appearance of element at position: %v\n", elementFirstAppearance)
	}
	// map
	newMap := mymap.CreateMap()
	fmt.Printf("%v\n", newMap)

	str := "bhupesh thakur"
	myMap, err := mymap.SliceToMap(str)
	fmt.Printf("%v\n", myMap)
}
