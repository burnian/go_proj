package main

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/google/uuid"
)

// iota can be used only in const() block
const (
	BEIJING = 3 * iota
	SHANGHAI
	GUANGZHOU
	SHENZHEN
)

const (
	Summer string = "summer"
	Autumn        = "autumn"
	Winter        = "winter"
	Spring        = "spring"
)

type Book struct {
	Title  string `json:"jtitle"`  //info:"抬头"
	Author string `json:"jauthor"` //info:"作者"
}

var nSum int = 0
var mu sync.Mutex

func Add(i int) {
	// mu.Lock()
	nSum += 10
	// mu.Unlock()
}

func Coordinates(value ...int) {
	a := value
	fmt.Println(a)
}

func main() {
	task_id := uuid.New().String()
	fmt.Println(task_id)

	// whole_time := time.Now().Format("2006.01.02 15:04:05")
	// fmt.Println(whole_time)
	// date := time.Now().Format("0102")
	// fmt.Println(date)

	// a := 123
	// b := strconv.Itoa(a)
	// b += "hello"
	// fmt.Println(b)

	// a := "1234567891234567891"
	// u64, err := strconv.ParseUint(a, 10, 32)
	// fmt.Println(u64, err)

	// b := "1234567891234567891"
	// str, err := strconv.Atoi(b)
	// fmt.Println(str, err)

	// books := []Book{}
	// book := Book{"little_prince", "John"}
	// p := book.Clone()
	// fmt.Println(*p)
	// p.Title = "new_title"
	// p.Author = "new_author"
	// fmt.Println(*p)
	// fmt.Println(book)

	// // append one element to an array from a function
	// books = append(books, book)
	// books = append(books, book)
	// books = append(books, *p)
	// fmt.Println(len(books))
	// f6(&books)
	// fmt.Println(len(books))
	// fmt.Println(books)

	// a := 2
	// switch a {
	// case 1:
	// 	fmt.Println("a is 1")
	// case 2:
	// 	fmt.Println("a is 2")
	// case 3:
	// 	fmt.Println("a is 3")
	// default:
	// 	fmt.Println("a is not 1, 2, 3")
	// }

	//////////////////////////////////////////////////
	// goroutine test
	// var wg sync.WaitGroup
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		Add(0)
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>nSum:", nSum)

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
	// map test
	// m := make(map[string]int, 0)
	// m["a"] = 1
	// m["b"] = 2
	// v, ok := m["c"]
	// fmt.Println(v, ok)

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

	//////////////////////////////////////////////////
	// file test
	// fp, err := os.Create("/home/zhoubl1/dev/test.txt")
	// if err != nil {
	// fmt.Println("create file failed:", err)
	// }
	// defer fp.Close()

	// file_path := "/nas/cache/dev/test.txt"
	// fmt.Println(path.Base(file_path))
	// fmt.Println(path.Dir(file_path))
	// fmt.Println(path.Ext(file_path))
	// fmt.Println(path.Split(file_path))
	// str := strings.TrimSuffix(path.Base(file_path), ".txt")
	// fmt.Println(str)

	// m := map[string]bool{"a": true, "b": true}
	// fmt.Println(m["a"])
	// fmt.Println(m["c"])

	// tile_id_dir := "/data/tomtom/country_tile_ids"
	// file_paths, err := filepath.Glob(path.Join(tile_id_dir, "*.txt"))
	// for _, file_path := range file_paths {
	// 	fmt.Println(file_path)
	// }
	// fmt.Println(err)
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

// func f5() Book {
// 	return Book{"little_prince", "John"}
// }

func f6(a *[]Book) {
	fmt.Println("f6", len(*a), cap(*a))
	book := Book{"aaa", "bbb"}
	*a = append(*a, book)
	fmt.Println("f6", len(*a), cap(*a))
}

func (x Book) Clone() *Book {
	p := new(Book)
	*p = x
	return p
}

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
