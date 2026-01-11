package main

import (
	"errors"
	"fmt"
	"iter"
	"maps"
	"math"
	"slices"
	"time"
	"unicode/utf8"
)

// go has automatic type inference
// default integer type is int
// default floating point type is float64
// default complex type is complex128
// := can be used only inside functions
// := cannot be used declare const
// Constants can be character, string, boolean, or numeric values.

const PI = 3.14

func main() {
	// ---------------- 1 Hello World ------------------
	fmt.Println("Hello, World")

	// ---------------- 2 Variables and Types ----------
	var a int // variable declaration

	// multiple variables can be declared at once using var
	var p, q int

	fmt.Printf("p = %d; q = %d\n", p, q) // by default zero value of the type is stored in undefined variables

	fmt.Println(a) // using undefined variables does not throw an error
	// the above line prints out 0 which is the default value of int

	a = 1 // variable definition

	fmt.Printf("a = %d ; type = %T\n", a, a)

	// go supports automatic type inference if a variable is defined in the same line as it is declared
	var r = 2 // the type of r is infered based on the RHS

	fmt.Println("r = ", r)

	// the below syntax is available only inside functions
	b := 2 // variable declaration and definition with automatic type inference

	fmt.Printf("b = %d ; type = %T\n", b, b)

	// Go supports constants of character, string, boolean, and numeric values.
	// const definition can either be inside a function body or outside
	const constNum int = 10

	const constFloat = 1.234

	// Constant expressions perform arithmetic with arbitrary precision.

	const smallNum = 100000000

	const bigNum = 3e20 / smallNum

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println("bigNum of type int64 = ", int64(bigNum))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(bigNum))

	// primtive types in go
	// numeric - int(int8, int16, int32, int64)(signed); uint(uint8, uint16, uint32, uint64)(unsigned) ; float (float32, float64)
	// string
	// there is no char in go ?
	// bool
	// byte

	var uA uint // declaring and not using a variable leads to a compiler error
	uA = 1
	fmt.Println(uA)

	var fA float32
	fA = 1.234
	fmt.Println(fA)

	fA1 := 2.3434 // default float type is float64
	fmt.Println(fA1)

	// default integer type is int
	// default unsigned integer type is uint

	var sA string
	sA = "Hello"
	fmt.Println(sA)

	var b1 bool
	b1 = true
	fmt.Println(b1)

	fmt.Printf("b1 = %v, type of b1 is %T\n", b1, b1)

	// --------------- 3 Operators ---------------
	// arithemtic
	num1 := 1
	num2 := 3
	fNum1 := 1.231

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)

	// its a compiler error to do arithmetic ops between two different integer types
	// to do this we have explicitly cast one of the operands

	fmt.Println(fNum1 + float64(num2))
	fmt.Println(fNum1 - float64(num2))
	fmt.Println(fNum1 * float64(num2))
	fmt.Println(fNum1 / float64(num2))

	// comparison operator
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)

	// similar to arithmetic operators in comparison operators also explicit casting is necessary

	fmt.Println(fNum1 > float64(num2))
	fmt.Println(fNum1 < float64(num2))
	fmt.Println(fNum1 >= float64(num2))
	fmt.Println(fNum1 <= float64(num2))
	fmt.Println(fNum1 == float64(num2))
	fmt.Println(fNum1 != float64(num2))

	// boolean operators
	b2 := false
	fmt.Println(b1 && b2)
	fmt.Println(b1 || b2)
	fmt.Println(!b1)

	// bitwise operators
	fmt.Println(num1 & num2)
	fmt.Println(num1 | num2)
	fmt.Println(num1 ^ num2)
	fmt.Println(^num2)

	// ^ is a overloaded operator in go ?
	// increment and decrement operator

	num1++
	fmt.Println()

	// go does not have a ternary operator

	// ------------------- 4 Conditionals and Loops -------------------
	if num1%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num1%3 == 0 {
		fmt.Println("multiple of 3")
	} else if num1%5 == 0 {
		fmt.Println("multiple of 5")
	} else if num1%7 == 0 {
		fmt.Println("multiple of 7")
	} else {
		fmt.Println("neither a multiple of 3, 5, or 7")
	}

	// if - else if - else blocks can optionally precede with a statement

	if num5 := 10; num5%2 == 0 {
		fmt.Println("multiple of 2")
	} else if num5%7 == 0 {
		fmt.Println("multiple of 7")
	} else {
		fmt.Println("neither a multiple 2 or 7")
	}

	// num5 is valid only in the current if - else if - else block it won't be available outside
	// fmt.Println(num5)

	// break statements in the cases is not necessary in go
	switch num2 {
	case 2:
		{
			fmt.Println("case - 2")
		}

	case 3:
		{
			fmt.Println("case - 3")
		}

	case 4:
		{
			fmt.Println("case - 4")
		}

	default: // default is optional
		{
			fmt.Println("default case ")
		}
	}

	// to handle many cases with only code block
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		{
			fmt.Println("Weekend")
		}

	default:
		{
			fmt.Println("Weekday")
		}
	}

	// like if - else if - else blocks a optional statement can be added before the switch selector
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		{
			fmt.Println("Weekend")
			fmt.Println(today)
		}

	default:
		{
			fmt.Println("Weekday")
			fmt.Println(today)
		}
	}

	// switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	t := time.Now()

	switch {
	case t.Hour() < 12:
		{
			fmt.Println("Good Morning")
		}

	case t.Hour() >= 12 && t.Hour() < 4:
		{
			fmt.Println("Good Afternoon")
		}

	case t.Hour() >= 4 && t.Hour() < 7:
		{
			fmt.Println("Good Night")
		}

	default:
		{
			fmt.Println("something is wrong...")
		}
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i any) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case float32:
			fmt.Println("float32")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("unknown type")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI(1.213)
	whatAmI(float32(1.213))

	// loops
	// there is only for loops in go
	// go does not have while or do - while loops

	// using the for loop as the traditional while loop
	i := 0
	for i < 10 {
		fmt.Printf("%d ", i)
		i += 1
	}

	fmt.Println()

	// traditional C style for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// range
	// range iterates over elements in a variety of built-in data structures.
	for i := range 5 {
		fmt.Println(i)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

	i = 0
	for {
		if i%7 == 0 {
			break
		} else if i%5 == 0 {
			continue // continue is used skip the current iteration
		}

		fmt.Printf("%d ", i)
		i += 1
	}

	fmt.Println()
	//------------------------------- 5 Strings and Runes  -------------------------------

	//  A Go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
	s := "Hello"
	fmt.Printf("s = %s ; type = %T", s, s)

	// to concatenate two strings
	s = "Hello, " + " World"
	fmt.Printf("s = %s ; type = %T\n", s, s)

	// strings are immutable in go
	// Go string literals are UTF-8 encoded text.
	// range on strings iterates over Unicode code points. The first value is the starting byte index of the rune and the second the rune itself.
	// a string in go is equivalent to []byte
	// Indexing into a string produces the raw byte values at each index. This loop generates the hex values of all the bytes that constitute the code points in s.
	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	for index, char := range s {
		fmt.Printf("%d -> (%c, %d, %#U)\n", index, char, char, char)
	}

	// utf8.DecodeRuneInString ????

	// to get the length of a string
	fmt.Println("length of s = ", len(s))

	// to count the number of runes in a string
	// some forgein language characters can span multiple bytes (?)
	fmt.Println(utf8.RuneCountInString(s))

	// Values enclosed in single quotes are rune literals.
	var r1 rune = 'r'

	fmt.Printf("%c , %d, %#U, %T\n", r1, r1, r1, r1)

	//  We can compare a rune value to a rune literal directly.
	fmt.Println(r1 == 'r')

	// rune ???

	// -------------------------------- 6 Arrays----------------------------------
	// array declaration
	// The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s
	var arr [5]int

	fmt.Println(arr)

	// We can set a value at an index using the array[index] = value syntax
	for i := range 5 {
		arr[i] = (i * 2)
	}

	fmt.Println(arr)

	// to get a value from the array use array[index] syntax.
	for i := range 5 {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Println()

	// to get the length of a array
	fmt.Println("length of arr = ", len(arr))

	// to declare and define a array using in one line
	arr1 := [6]int{1, 2, 3, 4} // rest of the elements have the 0 value of the type
	fmt.Println(arr1)

	// in arr1 definition we can let the go compiler infer the length
	arr2 := [...]int{3, 6, 9, 12, 15}
	fmt.Println(arr2, len(arr2))

	// If you specify the index with :, the elements in between will be zeroed.
	arr3 := [...]int{4, 3: 12, 16, 24}
	fmt.Println(arr3, len(arr3))

	// multidimensional arrays
	var twoDArr [2][3]int

	for i := range 2 {
		for j := range 3 {
			twoDArr[i][j] = i + j
		}
	}

	fmt.Println(twoDArr)

	// to declare and define multi dimensional arrays in one line
	// for some reason [...][...] is not working
	twoDimArr := [...][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(twoDimArr)

	// using range to iterate over a array
	for index, val := range arr3 {
		fmt.Printf("%d -> %d ; ", index, val)
	}

	fmt.Println()

	// -------------------------------- 7 Slices ----------------------------------
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.
	var slc []int

	// a slice is a (heap allocated ?) dynamic array ?
	fmt.Println(slc, len(slc))

	// to create a slice with non-zero length, use the builtin make.
	slc = make([]int, 3)

	fmt.Println(slc, len(slc), cap(slc)) // cap(slc) gives the capacity of the slice which is equal to its length by default

	// to create a slc of lenght 0 and capacity 10
	slc = make([]int, 0, 10)
	fmt.Println(slc, len(slc), cap(slc))

	// to add elements to the end of the array
	slc = append(slc, 1)

	fmt.Println(slc, len(slc), cap(slc))

	for i := range 5 {
		slc = append(slc, i+1)
	}

	fmt.Println(slc, len(slc), cap(slc))

	// to create a copy of a slice
	slc1 := make([]int, len(slc))
	copy(slc1, slc)

	slc = append(slc, 7)

	fmt.Println(slc, slc1)

	// Slices support a “slice” operator with the syntax slice[low:high].
	slc2 := slc[2:] // from index 2(included) to the end
	fmt.Println(slc2, len(slc2), cap(slc2))

	slc2 = slc[:4] // from start to index 4(excluded)
	fmt.Println(slc2, len(slc2), cap(slc2))

	slc2 = slc[1:5]
	fmt.Println(slc2, len(slc2), cap(slc2))

	// to declare and define a slice in one line
	slc2 = []int{4, 8, 12, 16, 20, 24}
	fmt.Println(slc2, len(slc2), cap(slc2))

	// The slices package contains a number of useful utility functions for slices.
	fmt.Println(slices.BinarySearch(slc2, 16))
	fmt.Println(slices.BinarySearch(slc2, 28))

	// to check two slices have the same data
	fmt.Println(slices.Equal(slc, slc2))

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoDSlc := make([][]int, 3)

	for i := range 3 {
		currSlcLength := i + 2
		twoDSlc[i] = make([]int, i+currSlcLength)
		for j := range currSlcLength {
			twoDSlc[i][j] = 2 + 2*(i+j)
		}
	}

	fmt.Println(twoDSlc)

	// iterating over a slice using  range
	for index, val := range slc {
		fmt.Printf("%d -> %d ; ", index, val)
	}

	fmt.Println()
	// ------------------------------- 8 Maps -------------------------------------

	// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

	// to create a map
	mp := make(map[string]int)

	// to add elements to the map
	mp["A"] = 1
	mp["B"] = 2
	mp["C"] = 3

	fmt.Println(mp)

	// runes in go ???
	mp1 := make(map[rune]int)

	mp1['A'] = 1
	mp1['B'] = 2
	mp1['C'] = 3

	fmt.Println(mp1)

	// to get a value for a key with name[key].
	fmt.Println(mp1['A'])

	// If the key doesn’t exist, the zero value of the value type is returned.
	fmt.Println(mp1['Z'])

	// to get number of key value pairs a map is holding
	fmt.Println(len(mp1))

	// to delete a key value pair
	delete(mp1, 'C')
	fmt.Println(mp1, len(mp1))

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(mp1)
	fmt.Println(mp1)

	// name[key] syntax returns a second bool value which indicates if the key is present in the map or not
	val, isKeyPresent := mp["A"]
	fmt.Println(val, isKeyPresent)

	// You can also declare and initialize a new map in the same line with this syntax.
	mp = map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	mp2 := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	// The maps package contains a number of useful utility functions for maps.
	// to check two maps contain the same data
	fmt.Println(maps.Equal(mp, mp2))

	// iterate over the key - value pairs of a map using range
	for key, value := range mp2 {
		fmt.Printf("%s -> %d ; ", key, value)
	}

	fmt.Println()

	// to iterate just over the keys
	for key := range mp2 {
		fmt.Printf("%s -> %d ; ", key, mp2[key])
	}

	fmt.Println()

	// ---------------------------------------- 9 Functions ---------------------------------------------
	// function call
	c := add(1, 2)
	fmt.Println(c)

	// go does not allow variable shadowing i.e. redefinition of variables is not allowed
	c = mul(3, 4)
	fmt.Println(c)

	fmt.Println(isEven(2))
	fmt.Println(isEven(3))

	fmt.Println(evenOrOdd(2))
	fmt.Println(evenOrOdd(3))

	// to ignore some values returned by the function

	num4, _ := div(3, 2)

	fmt.Println(num4)

	// Variadic functions can be called with any number of trailing arguments.
	fmt.Println(addn(1, 2, 3))
	fmt.Println(addn(1, 2, 3, 4, 5))

	// we can also pass in a slice for a function which takes in variable number of arguments
	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(addn(nums...))

	// go supports anonymous functions
	sub := func(a, b int) int { return a - b }

	fmt.Println(sub(3, 4))

	// using anonymous functions we can create closure
	fmt.Println("----------------------------------------------")

	// We call intSeq, assigning the result (a function) to curr. This function value captures its own i value, which will be updated each time we call curr.
	curr := intSeq()
	for i := range 5 {
		// curr := intSeq() // if curr is declared here everytime we get result of  curr() ti be 1
		fmt.Println(i, curr())
	}
	fmt.Println("----------------------------------------------")

	// go supports recursive functions
	fmt.Println(factorial(10))

	// tail recursive functions
	fmt.Println(factorialT(10))

	// anonymous functions can also be recursive

	var fib func(int) int // we have to declare it like this when a anonymous function is supposed to be recursive
	// := cannot be used to declare anonymous functions
	// if  := is used then the recurisve part won't know what you are referring to
	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	// ----------------------------------- 10 Pointers ----------------------------------------
	num5 := 3

	passByValue(num5)
	fmt.Println("num5 = ", num5)

	passByPointer(&num5) // & gives the address of the variable i.e a pointer to the variable
	fmt.Println("num5 = ", num5)

	// --------------------------------- 11 User Defined Types -------------------------------
	pr := *newPerson("A", 20)

	fmt.Println(pr)
	fmt.Printf("%T\n", pr)
	// Access struct fields with a dot.
	fmt.Println(pr.name, pr.age)

	pr1 := newPerson("B", 21)
	// can also use dots with struct pointers - the pointers are automatically dereferenced.
	fmt.Println(pr1.name, pr1.age)

	fmt.Println(pr1)
	fmt.Printf("%T\n", pr1)

	// Structs are mutable.
	pr1.age += 1
	fmt.Println(pr1)

	// If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type.
	v := struct {
		x int
		y int
		z int
	}{
		1, 2, 3,
	}

	fmt.Println(v)

	p1 := Point{1, 2}
	p2 := Point{3, 4}

	// Omitted fields will be zero-valued.
	p3 := Point{x: 4} // p3 := Point{4} - this is throwing a error

	fmt.Println("p3 = ", p3)

	// Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

	fmt.Println(p1.add(&p2))

	p1.printPoint()

	p1.inc()
	p2.inc()
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	fmt.Println("distance between p1 and p2 = ", p1.euclidean_distance(&p2))

	// Go supports embedding of structs and interfaces to express a more seamless composition of types. This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.

	st1 := newStudent("C", 20, "XYZ", 14)
	st1.PrintStudent()

	// Since Student embeds Person, the methods of Person also become methods of a Student. Here we invoke a method that was embedded from Person directly on st1.
	st1.PrintPerson()

	// Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a Student now implements the describer interface because it embeds Person.
	st1.describe()
	Describe(st1)

	// ---------------------------------------- 9 Interfaces -------------------------------
	//	Interfaces are named collections of method signatures.

	rc1 := Rectangle{2.3, 3.4}

	PrintShape(rc1)

	s1 := Square{3.5}

	PrintShape(s1)

	// ----------------------------------- 10 Enums -----------------------------------------
	// Enumerated types (enums) are a special case of sum types. An enum is a type that has a fixed number of possible values, each with a distinct name. Go doesn’t have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.
	fmt.Println(StateIdle)
	fmt.Println(StateConnected)
	fmt.Println(StateError)
	fmt.Println(StateRetrying)

	// --------------------------------- 11 Generics -----------------------------------------
	// Starting with version 1.18, Go has added support for generics, also known as type parameters.
	// When invoking generic functions, we can often rely on type inference. Note that we don’t have to specify the types for generic functions (?)

	sll := SLL[int]{}

	sll.push(1)
	sll.push(2)
	sll.push(3)

	sll.printList()

	// -------------------------------- 12 Iterators --------------------------------------
	// 	Starting with version 1.23, Go has added support for iterators, which lets us range over pretty much anything

	for curr := range sll.All() {
		fmt.Printf("%d -> ", curr)
	}

	fmt.Println("nil")

	// Packages like slices have a number of useful functions to work with iterators. For example, Collect takes any iterator and collects all its values into a slice.
	slc3 := slices.Collect(sll.All())
	fmt.Println(slc3)

	for n := range getFib() {
		if n > 20 {
			break // Once the loop hits break or an early return, the yield function passed to the iterator will return false.
		}

		fmt.Printf("%d ", n)
	}

	fmt.Println()

	// ------------------------------------- 13 Errors ---------------------------------------------------
	// In Go it’s idiomatic to communicate errors via an explicit, separate return value.

	fmt.Println(someErrFunc(4))
	fmt.Println(someErrFunc(10))

	fmt.Println(someErrorFunc1(4))
	fmt.Println(someErrorFunc1(7))
	fmt.Println(someErrorFunc1(10))

	// It’s idiomatic to use an inline error check in the if line.
	if i, err := someErrorFunc1(10); err == nil {
		fmt.Println(i, ": no error")
	}

	if i, err := someErrorFunc1(4); err != nil {
		fmt.Println("-------------------------------------")
		fmt.Println(i)
		fmt.Println("error : ", err)
		fmt.Println(errors.Is(UnLuckyNumberError, err)) // to check if a error is of certain error type ?
		fmt.Println("-------------------------------------")
	}

	fmt.Println(someErrorFunc2(40))
	fmt.Println(someErrorFunc2(3))

	// errors.As is a more advanced version of errors.Is. It checks that a given error (or any error in its chain) matches a specific error type and converts to a value of that type, returning true. If there’s no match, it returns false.

	_, err := someErrorFunc2(40)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("conversion failed")
	}

	// ----------------------------------------- 14 Go Routines ---------------------------------------
	// A goroutine is a lightweight thread of execution.
	// To invoke this function in a goroutine, use go f(s).
	go goFunc1()
	go goFunc2()

	time.Sleep(5000 * time.Millisecond)

	// You can also start a goroutine for an anonymous function call.

	fmt.Println("--------------------------------------------------------")

	go func() {
		for i := range 3 {
			fmt.Println("func - 1 ; task - ", i)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		for i := range 3 {
			fmt.Println("func - 2 ;task - ", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(5000 * time.Millisecond)

	// ---------------------------------- 15 Channels ---------------------------------

	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
	// creating a new channel
	ch := make(chan string)

	go func() { ch <- "ping" }() // channel <- "value" sending messages to the channel

	msg := <-ch // <- channel receiving a message from the channel

	fmt.Println(msg)

	// By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

	ch1 := make(chan string, 2) // channel of strings buffering up to 2 values.

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.

	ch1 <- "msg1"
	ch1 <- "msg2"

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

// func function-name(<arg1 type>) return-type { body }
// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
func add(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func mul(a, b int) (c int) {
	c = a * b
	return // naked return
}

func isEven(num int) bool {
	return num%2 == 0
}

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func newtonSquareRoot(num int) float64 {
	x := float64(num)
	n := float64(num)
	iter := 0
	for {
		if (x*x-n <= 0) || (iter > 10000) {
			break
		}

		x = (x + (n / x)) / 2.0
	}

	return x
}

func isPrime(num int) bool {
	maxNum := int(newtonSquareRoot(num))

	for i := 2; i <= maxNum; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero error")
	}

	return a / b, nil
}

// variadic functions = functions which can take in any number of arguments
func addn(nums ...int) int {

	// Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range

	res := 0
	for _, val := range nums {
		res += val
	}

	return res
}

// closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// recursion
func factorial(n int) int {
	if n <= 0 {
		return 1
	}

	return n * factorial(n-1)
}

// tail recursive functions
func factorialT(n int) int {
	return factorialHelperFunc(n, 1)
}

func factorialHelperFunc(n int, acc int) int {
	if n <= 0 {
		return acc
	}

	return factorialHelperFunc(n-1, acc*n)
}

func passByValue(a int) {
	a += 1
	fmt.Println("a = ", a)
}

func passByPointer(a *int) { // accepting a pointer to a value of int type
	*a += 1                 // de-referencing a pointer and mutating it mutates the actual value at that address
	fmt.Println("a = ", *a) // de-referencing a pointer
}

// ---------------------------------------------- UDT -------------------------------------------
type Point struct {
	x int
	y int
}

// Go supports methods defined on struct types.
// the add method receives a pointer to the object of Point type
// Methods can be defined for either pointer or value receiver types
func (self *Point) add(other *Point) Point {
	return Point{self.x + other.x, self.y + other.y}
}

func (self *Point) inc() {
	self.x += 1
	self.y += 1
}

func (self *Point) euclidean_distance(other *Point) float64 {
	return math.Sqrt(float64((other.x-self.x)*(other.x-self.x) + (other.y-self.y)*(other.y-self.y)))
}

// the below method receives a value and not a pointer
func (self Point) printPoint() {
	fmt.Println("Point => ", self)
}

type Person struct {
	name string
	age  int
}

// like a constructor but completely optional
// Go is a garbage collected language; you can safely return a pointer to a local variable - it will only be cleaned up by the garbage collector when there are no active references to it.
// It’s idiomatic to encapsulate new struct creation in constructor functions
func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age} // An & prefix yields a pointer to the struct.
}

func (p Person) PrintPerson() {
	fmt.Println("name = ", p.name, " ; age = ", p.age)
}

// -------------------------------------- Interfaces -------------------------------------------------------
// interfaces like traits in Rust
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

// if the below function take a pointer to the rectangle some how Shape interface is not being implemented by Rectangle ??
func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type Square struct {
	sideLength float64
}

func (s Square) Area() float64 {
	return s.sideLength * s.sideLength
}

func (s Square) Perimeter() float64 {
	return 4 * s.sideLength
}

func PrintShape(s Shape) {
	fmt.Println("-----------------------------------------------")
	fmt.Println(s)

	//to check the actual type received
	if _, ok := s.(Rectangle); ok {
		fmt.Println("Rectangle")
	}

	if _, ok := s.(Square); ok {
		fmt.Println("Square")
	}

	fmt.Println("area = ", s.Area())
	fmt.Println("perimeter = ", s.Perimeter())
	fmt.Println("-----------------------------------------------")
}

// struct embedding (like dependency injection in java ?)
type Student struct {
	Person      //An embedding looks like a field without a name.
	institution string
	grade       int
}

// When creating structs with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
func newStudent(name string, age int, institution string, grade int) *Student {
	return &Student{
		Person: Person{
			name: name,
			age:  age,
		},
		institution: institution,
		grade:       grade,
	}
}

// We can access the base’s fields directly on co, e.g. co.num.
// Alternatively, we can spell out the full path using the embedded type name.
func (s Student) PrintStudent() {
	fmt.Println("Name = ", s.name, " ; Age = ", s.Person.age, " ; Institution = ", s.institution, "; Grade = ", s.grade)
}

type describer interface {
	describe() string
}

func Describe(d describer) {
	fmt.Println(d.describe())
}

func (p Person) describe() string {
	return fmt.Sprintf("name = %s ; age = %d", p.name, p.age)
}

// -------------------------------------------------------------------------------------------------
// The possible values for ServerState are defined as constants. The special keyword iota generates successive constant values automatically; in this case 0, 1, 2 and so on.

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the fmt.Stringer interface, values of ServerState can be printed out or converted to strings.
var serverStateToString = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (s ServerState) String() string {
	return serverStateToString[s]
}

// ------------------------------------------ Generics -------------------------------------------

type SLLNode[T comparable] struct {
	data T
	next *SLLNode[T]
}

func (node *SLLNode[T]) String() string {
	return fmt.Sprintf("%v", node.data) // In Go, the format specifier %v is used to print the value of a variable in its default format. ?
}

type SLL[T comparable] struct {
	head, tail *SLLNode[T]
	length     int
}

func (s *SLL[T]) push(newElement T) {
	newNode := SLLNode[T]{data: newElement, next: nil}
	if s.head == nil {
		s.head = &newNode
		s.tail = s.head
	} else {
		s.tail.next = &newNode
		s.tail = s.tail.next
	}

	s.length++
}

func (s *SLL[T]) printList() {
	curr := s.head
	for curr != nil {
		fmt.Printf("%v -> ", curr)
		curr = curr.next
	}

	fmt.Println("nil")
}

// iterator
// All returns an iterator, which in Go is a function with a special signature.
func (sll *SLL[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		// The iterator function takes another function as a parameter, called yield by convention (but the name can be arbitrary). It will call yield for every element we want to iterate over, and note yield’s return value for a potential early termination.

		for curr := sll.head; curr != nil; curr = curr.next {
			if !yield(curr.data) {
				return
			}
		}
	}
}

// Iteration doesn’t require an underlying data structure, and doesn’t even have to be finite! Here’s a function returning an iterator over Fibonacci numbers: it keeps running as long as yield keeps returning true.
func getFib() iter.Seq[int] { // ???
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}

			a, b = b, a+b
		}
	}
}

// ------------------------------------------ Errors ----------------------------------------
// In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java, Python and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

// By convention, errors are the last return value and have type error, a built-in interface.

func someErrFunc(n int) (int, error) {
	if n == 4 {
		return -1, errors.New("bad number") // errors.New constructs a basic error value with the given error message.
	}

	return (n * 4), nil // A nil value in the error position indicates that there was no error.
}

// A sentinel error is a predeclared variable that is used to signify a specific error condition.
var UnLuckyNumberError = errors.New("unlucky number")
var BadNumberError = errors.New("bad number")

func someErrorFunc1(n int) (int, error) {
	if n%4 == 0 {
		return -1, UnLuckyNumberError
	}

	if n%7 == 0 {
		return -1, fmt.Errorf("bad number : %w", BadNumberError) // % w ???
	}

	return n * 3, nil
}

// errors.As ???

// It’s possible to define custom error types by implementing the Error() method on them.

// A custom error type usually has the suffix “Error”.
type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d ; %s\n", e.arg, e.message)
}

func someErrorFunc2(n int) (int, error) {
	if n%4 == 0 {
		return -1, &argError{arg: n, message: "unlucky number"}
	}

	return n * 3, nil
}

// ------------------------------------ go threads ----------------------------------
func goFunc1() {
	for i := range 3 {
		fmt.Println("func - 1 ; task - ", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func goFunc2() {
	for i := range 3 {
		fmt.Println("func - 2 ; task - ", i)
		time.Sleep(1000 * time.Millisecond)
	}
}
