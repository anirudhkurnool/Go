package main

import (
	"errors"
	"fmt"
	"iter"
	"math"
	"unicode/utf8"
)

func add(a int, b int) int {
	return a + b
}

func mul(a int, b int) int {
	return a * b
}

func sub(a int, b int) int {
	return a - b
}

func div(a int, b int) int {
	if b == 0 {
		panic("divide by zero error.")
	}
	return a / b
}

func is_even(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func printShape(s Shape) {
	fmt.Println("******* SHAPE *******")
	fmt.Println(s)
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) CreateRectangle(length, breadth float64) *Rectangle {
	return &Rectangle{length, breadth}
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	//2 automatically gets casted to float64???
	return 2 * (r.length + r.breadth)
}

func (r Rectangle) Info() (float64, float64) {
	return r.length, r.breadth
}

func (r Rectangle) String() string {
	return fmt.Sprintf("===== RECTANGLE =====\nlength     : %f\nbreadth    : %f\narea       : %f\nperimeter  : %f", r.length, r.breadth, r.Area(), r.Perimeter())
}

func printShapeRectagle(s Shape) {
	fmt.Println("RECTANGLE SHAPE ONLY")
	//type assertion
	if rect, ok := s.(Rectangle); ok {
		fmt.Println(rect)
	}
}

// struct embedding
// as rectangle in embedded in this new struct this one also implements Shape interface
type SquareWithRectEmbedded struct {
	sideLength float64
	Rectangle
}

type Square struct {
	sideLength float64
	rect       Rectangle
}

func (s *Square) CreateSquare(sideLength float64) *Square {
	s1 := Square{
		sideLength: sideLength,
		rect:       Rectangle{sideLength, sideLength},
	}

	return &s1
}

func (s *Square) Area() float64 {
	return s.rect.Area()
}

func (s *Square) Perimeter() float64 {
	return s.rect.Perimeter()
}

func (s *Square) String() string {
	return fmt.Sprintf("===== SQUARE =====\nside-length: %f\narea       : %f\nperimeter  : %f", s.sideLength, s.Area(), s.Perimeter())
}

func is_prime(num int) bool {
	end := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 2; i <= end; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func sieveOfEratosthenes(num int) []int {
	//var bool_arr bool[num] //can't use arrays as array lengths have to be compile time constants
	bool_arr := make([]bool, num)
	for index, _ := range bool_arr {
		bool_arr[index] = true
	}

	bool_arr[0] = false
	var res []int
	for i := range num {
		if bool_arr[i] == true {
			for j := (2*i + 1); j < num; j += (i + 1) {
				bool_arr[j] = false
			}

			res = append(res, (i + 1))
		}
	}

	return res
}

// take any number of integer arguments and also a integer slice
func sum(nums ...int) int {
	res := 0
	for _, val := range nums {
		res += val
	}

	return res
}

// "Operationally, a closure is a record storing a function[a] together with an environment." - https://en.wikipedia.org/wiki/Closure_(computer_programming)
// Closures ???
func nextInteger() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func factorial(num int) int {
	if num < 0 {
		panic("factorial of negitive number is undefined")
	}
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

// func fibonacci {

// }
//

func addNextInt(arr *[]int) {
	lastNum := (*arr)[len(*arr)-1]
	*arr = append(*arr, lastNum+1)
}

func passRuneToFunc(r rune) {
	if r == 'a' {
		fmt.Println("a")
	} else {
		fmt.Println("something else")
	}
}

func SearchInGenericList[genLst ~[]genericElement, genericElement comparable](lst genLst, element genericElement) bool {
	//return lst.Contains(element)

	for _, value := range lst {
		if value == element {
			return true
		}
	}

	return false
}

type GenericLinkedListNode[T any] struct {
	data T
	next *GenericLinkedListNode[T]
	prev *GenericLinkedListNode[T]
}

func genFib() iter.Seq[int] {
	//from  - https://gobyexample.com/range-over-iterators
	//
	//example of a infinite iterator
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !(yield(a)) {
				return
			}

			a, b = b, a+b
		}
	}
}

func div_err(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero error")
	}

	return (a / b), nil
}

func simpleError() error {
	err2 := fmt.Errorf("simple error")
	if 1 == 1 {
		return err2
	}

	return nil
}

func errorInsideError() (int, error) {
	res, err := div_err(3, 0)
	if err != nil {
		return -1, fmt.Errorf("error while division : %s", err)
	} else {
		return res, nil
	}
}

func checkErrors(err error, errors1 []error) {
	for index := range errors1 {
		if errors.Is(err, errors1[index]) {
			fmt.Printf("err - 2 => %s\n", errors1[index])
		}
	}
}

func main() {
	fmt.Println("Hello, World!!!")
	fmt.Println("1 + 2 = ", add(1, 2))
	fmt.Println("3 * 4 = ", mul(3, 4))
	fmt.Println("5 / 2 = ", div(5, 2))
	fmt.Println("3 - 4 = ", sub(3, 4))
	fmt.Printf("3 is : %s\n", is_even(3))
	fmt.Printf("4 is : %s\n", is_even(4))
	fmt.Printf("is 3 prime : %t\n", is_prime(3))
	fmt.Printf("is 4 prime : %t\n", is_prime(4))
	fmt.Println(sieveOfEratosthenes(100))

	var r Rectangle
	r = *(r.CreateRectangle(3.1234, 4.5466))
	fmt.Println(r)
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	rLength, rBreadth := r.Info()
	fmt.Println(rLength, rBreadth)

	var s *Square
	s = s.CreateSquare(5.14234)
	fmt.Println(s)

	fmt.Println(sum(1, 2, 3, 4, 5))
	nums := []int{6, 7, 8, 9, 10}
	fmt.Println(sum(nums...))

	nextInt := nextInteger()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Printf("Factorial of 5  : %d\n", factorial(5))
	fmt.Printf("Factorial of 10 : %d\n", factorial(10))

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before : %d\n", arr2)
	addNextInt(&arr2)
	fmt.Printf("After : %d\n", arr2)

	//strings are read only slices of bytes
	//strinfs are UTF-8 encoded
	//a char in go is called a rune which is a character's UTF-8 int representation(Unicode code point)

	str := "Hello"
	for _, val := range str {
		fmt.Printf("%x\n", val)
	}

	//rune  ????
	//In UTF-8 some characters can span more than one byte
	//rune literals are kept in '' in golang
	//Iteration happens over raw bytes
	const str2 = "สวัสดี" //string from https://gobyexample.com/strings-and-runes
	fmt.Printf("length of string :  %d", len(str2))

	//"A range loop handles strings specially and decodes each rune along with its offset in the string." - https://gobyexample.com/strings-and-runes
	for i, val := range str2 {
		fmt.Printf("%d : %x\n", i, val)
	}

	fmt.Printf("rune count in str2 : %d\n", utf8.RuneCountInString(str2))

	passRuneToFunc('a')

	runeValue, width := utf8.DecodeRuneInString(str2)
	fmt.Println(runeValue, width)

	printShape(r)
	printShape(s)
	printShapeRectagle(r)

	s1 := SquareWithRectEmbedded{
		sideLength: 4.56,
		Rectangle: Rectangle{
			length:  4.56,
			breadth: 4.56,
		},
	}

	printShape(s1)
	//fmt.Println(SearchInGenericList[[]int, int]([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(SearchInGenericList([]int{1, 2, 3, 4, 5}, 4)) //type infered

	lstNode := GenericLinkedListNode[float64]{3.14, nil, nil}
	fmt.Println(lstNode)

	//Iterator in go is a function with special signature - https://pkg.go.dev/iter#Seq
	//anything which returns a iterator can be used in for range loop
	fmt.Println("--------- Fibonacci Sequence ---------")
	for n := range genFib() {
		if n > 10 {
			break
		}

		fmt.Println(n)
	}

	//"In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C."
	// - https://gobyexample.com/errors
	res1, err1 := div_err(3, 0)
	fmt.Println(res1, err1)

	// official go errors - https://go.dev/blog/go1.13-errors

	//simple error
	fmt.Println(simpleError())

	//error inside a error
	fmt.Println(errorInsideError())

	//check which error
	err2 := fmt.Errorf("error type - 1")
	err3 := fmt.Errorf("error type - 2")
	errors := []error{err2, err3}
	checkErrors(err2, errors)
	checkErrors(err3, errors)
}
