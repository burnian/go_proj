package main

import (
	"fmt"
	"reflect"
)

// iota can be used only in const() block
const (
	BEIJING = 3 * iota
	SHANGHAI
	GUANGZHOU
	SHENZHEN
)

type Book struct {
	Title  string `json:"jtitle"`  //info:"抬头"
	Author string `json:"jauthor"` //info:"作者"
}

func init() {
	fmt.Println("main.go init()")
}

// func f1() {
// 	fmt.Println("f1")
// }

// func f2() {
// 	fmt.Println("f2")
// }

// func f3() {
// 	fmt.Println("f3")
// }

// func f4(arr [5]int) {
// 	arr[0] = 99
// }

// type assertion test
func generalTypeTest(arg interface{}) {
	v1, flag1 := arg.(int)
	fmt.Println("v1", v1)
	if flag1 {
		fmt.Println("all_type_test is int:", v1)
		return
	}

	v2, flag2 := arg.(string)
	fmt.Println("v2", v2)
	if flag2 {
		fmt.Println("all_type_test is string:", v2)
		return
	}

	v3, flag := arg.(Book)
	fmt.Println("v3", v3)
	if flag {
		fmt.Println("all_type_test is Book:", v3.Title, "from", v3.Author)
		return
	}
}

func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)

	fmt.Println("inputType:", inputType)
	fmt.Println("inputValue:", inputValue)

	// get field
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Println("field:", field.Name, "type:", field.Type, "value:", value)
	}

	// get method
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Println("method:", method.Name, "type:", method.Type)
	}
}

func findTag(input interface{}) {
	elems := reflect.TypeOf(input).Elem()
	for i := 0; i < elems.NumField(); i++ {
		dic_tag := elems.Field(i).Tag
		fmt.Println("tag info:", dic_tag.Get("info"), "json:", dic_tag.Get("json"))
	}
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x += y
			x, y = y, x
		case <-quit:
			return
		}
	}
}

func main() {
	//////////////////////////////////////////////////
	// defer will be executed after the function returns as LIFO
	// defer f1()
	// defer f2()
	// defer f3()

	//////////////////////////////////////////////////
	// channel test
	// c := make(chan int)
	// quit := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()

	// fibonacci(c, quit)

	// c := make(chan int, 3)

	// go func() {
	// 	defer fmt.Println("sub goroutine end...")

	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 		fmt.Println("write:", i, ", len(c):", len(c), ", cap(c):", cap(c))
	// 	}

	// 	close(c)
	// }()

	// // for {
	// // 	if num, ok := <-c; ok {
	// // 		fmt.Println("read:", num, ", len(c):", len(c), ", cap(c):", cap(c))
	// // 	} else {
	// // 		break
	// // 	}
	// // }
	// // can also use range to read from channel
	// for num := range c {
	// 	fmt.Println("read:", num, ", len(c):", len(c), ", cap(c):", cap(c))
	// }

	// fmt.Println(len(c))
	// fmt.Println("main goroutine end...")

	//////////////////////////////////////////////////
	// type test
	// generalTypeTest(12)
	// generalTypeTest("hello")
	// generalTypeTest(Book{"little prince", "John"})

	//////////////////////////////////////////////////
	// reflect test
	// book := Book{"little_prince", "John"}
	// DoFieldAndMethod(book)
	// findTag(&book)

	//////////////////////////////////////////////////
	// json test
	// book := Book{"little_prince", "John"}
	// jsonStr, err := json.Marshal(book)
	// if err != nil {
	// 	fmt.Println("json.Marshal failed:", err)
	// 	return
	// }
	// fmt.Printf("jsonStr: %s\n", jsonStr)

	// newBook := Book{}
	// err = json.Unmarshal(jsonStr, &newBook)
	// if err != nil {
	// 	fmt.Println("json.Unmarshal failed:", err)
	// 	return
	// }
	// fmt.Println("newBook:", newBook)

	//////////////////////////////////////////////////
	// fmt.Println(common.str_split("a,b,c", ","))

	// fmt.Println(BEIJING, SHANGHAI, GUANGZHOU, SHENZHEN)

	// arr1 := [5]int{1, 2, 3, 4, 5}
	// fmt.Printf("arr1 type: %T\n", arr1)
	// for i, v := range arr1 {
	// 	fmt.Println(i, v)
	// }

	// f4(arr1)
	// fmt.Println("arr1:", arr1)

	// arr2 := arr1[:]
	// fmt.Printf("arr2 type: %T, arr2:%v\n", arr2, arr2)

	// arr3 := make([]int, 5, 10)
	// fmt.Printf("arr3 len:%d, cap:%d, %v\n", len(arr3), cap(arr3), arr3)
	// arr3 = append(arr3, 1, 2, 3)
	// fmt.Println("arr3:", arr3)

	// copy(arr3, arr2)
	// fmt.Println("arr3:", arr3)

	//////////////////////////////////////////////////
	// map
	// m1 := make(map[string]int)
	// m1["aaa"] = 111
	// m1["bbb"] = 222
	// m1["ccc"] = 333
	// fmt.Println("m1:", m1)

	// delete(m1, "bbb")

	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }

	// m2 := make(map[string]int)
	// fmt.Println("m2:", m2)
}
