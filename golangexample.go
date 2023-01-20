package main

import (
	"fmt"
	// "math"
	// "time"
	// "unicode/utf8"
)

// const s string = "constant"

/*Functions*/

// func plus(a int, b int) int {
// 	return a + b
// }
// func plusPlus(a, b, c int) int {
// 	return a + b + c
// }

/*Multiple Return Values*/

// func vals() (int, int) {
// 	return 3, 7
// }

/*Variadic Functions*/

// func sum(nums ...int) {
// 	fmt.Println(nums, "")
// 	total := 0

// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }

/*Closures*/

// func intSeq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

/*Recursion*/

// func fact(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }

/*Pointers*/

// func zeroval(ival int) {
// 	ival = 0
// }

// func zeroptr(iptr *int) {
// 	*iptr = 0
// }

/*Strings and Runes*/

// func examineRune(r rune) {
// 	if r == 't' {
// 		fmt.Println("found tee")
// 	} else if r == 'ส' {
// 		fmt.Println("found so sua")
// 	}
// }

/*Structs*/

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string) *person {
// 	p := person{name: name}
// 	p.age = 42

// 	return &p
// }

/*Methods*/

// type rect struct {
// 	width, height int
// }

// func (r *rect) area() int {
// 	return r.width * r.height
// }

// func (r rect) perim() int {
// 	return 2*r.width + 2*r.height
// }

/*Interfaces*/

// type geometry interface {
// 	area() float64
// 	perim() float64
// }

// type rect struct {
// 	width, height float64
// }

// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (r rect) perim() float64 {
// 	return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

/*Struct Embedding*/

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

/*Generics*/

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {

	/*Values*/

	// fmt.Println("go" + "lang")

	// fmt.Println("1+1 =", 1+1)
	// fmt.Println("7.0/3.0 =", 7.0/3.0)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	/*Variables*/

	// var a = "initial"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)

	// var d = true
	// fmt.Println(d)

	// var e int
	// fmt.Println(e)

	// f := "apple"
	// fmt.Println(f)

	/*Constants*/

	// fmt.Println(s)

	// const n = 500000000

	// const d = 3e20 / n
	// fmt.Println(d)

	// fmt.Println(int64(d))

	// fmt.Println(math.Sin(n))

	/*For*/

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	/*If/Else*/

	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	/*Switch*/

	// i := 2
	// fmt.Println("Write ", i, " as")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }

	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	/*Arrays*/

	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])

	// fmt.Println("len:", len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	/*Slices*/

	// s := make([]string, 3)
	// fmt.Println("emp:", s)

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set:", s)
	// fmt.Println("get:", s[2])

	// fmt.Println("len:", len(s))

	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("apd: ", s)

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy:", c)

	// l := s[2:5]
	// fmt.Println("sl1:", l)

	// l = s[:5]
	// fmt.Println("sl2:", l)

	// l = s[2:]
	// fmt.Println("sl3:", l)

	// t := []string{"g", "h", "i"}
	// fmt.Println("dcl:", t)

	// twoD := make([][]int, 3)
	// for i := 0; i < 3; i++ {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }

	// fmt.Println("2d: ", twoD)

	/*Maps*/

	// m := make(map[string]int)

	// m["k1"] = 7
	// m["k2"] = 13

	// fmt.Println("map:", m)

	// v1 := m["k1"]
	// fmt.Println("v1: ", v1)

	// fmt.Println("len:", len(m))

	// delete(m, "k2")
	// fmt.Println("map:", m)

	// _, prs := m["k2"]
	// fmt.Println("prs:", prs)

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map: ", n)

	/*Range*/

	// nums := []int{2, 3, 4}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("sum:", sum)

	// for i, num := range nums {
	// 	if num == 3 {
	// 		fmt.Println("index:", i)
	// 	}
	// }

	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }

	// for k := range kvs {
	// 	fmt.Println("key: ", k)
	// }

	// for i, c := range "go" {
	// 	fmt.Println(i, c)
	// }

	/*Functions*/

	// res := plus(1, 2)
	// fmt.Println("1+2 = ", res)

	// res = plusPlus(1, 2, 3)
	// fmt.Println("1+2+3 = ", res)

	/*Multiple Return Values*/

	// a, b := vals()
	// fmt.Println(a)
	// fmt.Println(b)

	// _, c := vals()
	// fmt.Println(c)

	/*Variadic Functions*/

	// sum(1, 2)
	// sum(1, 2, 3)

	// nums := []int{1, 2, 3, 4}
	// sum(nums...)

	/*Closures*/

	// nextInt := intSeq()

	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// newInts := intSeq()
	// fmt.Println(newInts())

	/*Recursion*/

	// fmt.Println(fact(7))

	// var fib func(n int) int

	// fib = func(n int) int {
	// 	if n < 2 {
	// 		return n
	// 	}
	// 	return fib(n-1) + fib(n-2)
	// }

	// fmt.Println(fib(7))

	/*Pointers*/

	// i := 1
	// fmt.Println("initial:", i)

	// zeroval(i)
	// fmt.Println("zeroval:", i)

	// zeroptr(&i)
	// fmt.Println("zeroptr:", i)

	// fmt.Println("pointer:", &i)

	/*Strings and Runes*/

	// const s = "สวัสดี"

	// fmt.Println("Len:", len(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%x", s[i])
	// }
	// fmt.Println()

	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// for idx, runeValue := range s {
	// 	fmt.Printf("%#U starts at  %d\n", runeValue, idx)
	// }

	// fmt.Println("\nUsing DecodeRuneInString")
	// for i, w := 0, 0; i < len(s); i += w {
	// 	runeValue, width := utf8.DecodeRuneInString(s[i:])
	// 	fmt.Printf("%#U starts at %d\n", runeValue, i)
	// 	w = width

	// 	examineRune(runeValue)
	// }

	/*Structs*/

	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{name: "Alice", age: 30})

	// fmt.Println(person{name: "Fred"})

	// fmt.Println(&person{name: "Ann", age: 40})

	// fmt.Println(newPerson("Jon"))

	// s := person{name: "Sean", age: 50}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp.age)

	// sp.age = 51
	// fmt.Println(sp.age)

	/*Methods*/

	// r := rect{width: 10, height: 5}

	// fmt.Println("area: ", r.area())
	// fmt.Println("perim: ", r.perim())

	// rp := &r
	// fmt.Println("area: ", rp.area())
	// fmt.Println("perim: ", rp.perim())

	/*Interfaces*/

	// r := rect{width: 3, height: 4}
	// c := circle{radius: 5}

	// measure(r)
	// measure(c)

	/*Struct Embedding*/

	// co := container{
	// 	base: base{
	// 		num: 1,
	// 	},
	// 	str: "some name",
	// }

	// fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// fmt.Println("also num:", co.base.num)

	// fmt.Println("describe:", co.describe())

	// type describer interface {
	// 	describe() string
	// }

	// var d describer = co
	// fmt.Println("describer:", d.describe())

	/*Generics*/

	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
