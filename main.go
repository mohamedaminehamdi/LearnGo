package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//first code
/* multi
line
comment */
const PI float64 = 3.1415926

func main() {
	title := "Hello World!"
	var s string = "good bye"
	var b bool = true
	var f float64 = 77.89
	var i int = 5
	fmt.Printf("%d", i)
	fmt.Print(s, i, b, f)
	fmt.Println(title)

	/*count, err := fmt.Scanf("%s %t", &s, &b)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(count)
	fmt.Println(err)*/

	var m float64 = float64(i)
	var mm int = int(m)
	fmt.Println(mm)

	fmt.Printf("%v %T\n", s, s)
	fmt.Println(reflect.TypeOf(title))

	s = strconv.Itoa(mm)
	fmt.Println(s)

	i, err := strconv.Atoi(s)
	fmt.Println(i, err)

	const age = 12
	fmt.Println(age)

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	//arrays
	var grades [5]int = [5]int{1, 2, 3, 4, 5}
	grades[0] = 1
	notes := [...]int{3, 2}
	fmt.Println(len(notes))
	for index, note := range notes {
		fmt.Println(index, note)
	}
	//slices
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	slice_1 := grades[1:4]
	slice_2 := make([]int, 5, 10)
	fmt.Println(slice_1, slice_2)
	fmt.Println(cap(slice_2))
	slice_1 = append(slice_1, 0)
	slice_1 = append(slice_1, slice_2...)
	num := copy(slice_1, slice_2)
	fmt.Println(num)
	for _, value := range slice_1 {
		fmt.Println(value)
	}

}
