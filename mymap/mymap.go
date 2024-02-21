package mymap

import (
	"errors"
	"fmt"
)

func CreateMap() map[string]int {
	//two ways to create a map make and litrals
	// newMap := map[string]int{
	// 	"employee_age":    30,
	// 	"employee_salary": 20000,
	// }
	newMap := make(map[string]int)
	newMap["employee_age"] = 30
	newMap["employee_salary"] = 20000
	return newMap
}

func SliceToMap(mystring string) (map[string]uint, error) {
	myMap := make(map[string]uint)
	if len(mystring) > 0 {
		for _, j := range mystring {
			stringJ := string(j)
			if myMap != nil && myMap[stringJ] > 0 {
				fmt.Println("myMap test ", myMap[stringJ])
				myMap[stringJ] = myMap[stringJ] + 1 //b <--hupesh
			} else {
				myMap[stringJ] = 1
			}
		}
		return myMap, nil
	} else {
		err := errors.New("slice is empty")
		return nil, err
	}
}
