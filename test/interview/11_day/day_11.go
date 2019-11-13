package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

func main() {

	day45()
}

// 不能使用cap 获取map 的容量

func day45() {
	fmt.Println(int(0x11))
}

func day46() {
	x := []int{1, 2}

	_ = x
}

func day47() {
	const x = 124
	const y = 123.4

	fmt.Println(x)

}

func day48() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)

	}
	fmt.Println(x)
}

func day49() {
	fmt.Println(^3)
}

// := 不能用于赋值给结构体内的字段
//

type People struct {
	Name string `json:"name"`
}

func day50() {
	js := `{
		"name" : "seekload"
	}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(p)
}

func (n N) test() {
	fmt.Println(n)
}

func day54() {
	var n N = 10
	fmt.Println(10)

	n++
	f1 := N.test
	f1(n)

	n++

	f2 := (*N).test
	f2(&n)

}

type N int

func (n N) value() {
	n++
	fmt.Printf("v:%v, %v \n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%v, %v \n", *n, *n)
}

func day55() {
	var a N = 25
	p := &a
	p.value()
	p.pointer()
}

func day56() {
	//nil := 123

	fmt.Println(nil)

	var _ map[string]int = nil

	var x int = -128
	var y = x / -1
	fmt.Println(y)
	//day30()
}

func f2(n int) (r int) {
	defer func() {
		r += n
		recover()
		fmt.Println(r)
	}()

	fmt.Println(r)

	var f2 func()

	defer f2()
	fmt.Println(r)
	f2 = func() {
		r += 2
		fmt.Println(r)
	}
	fmt.Println(r)
	return n + 1
}

func day30() {
	fmt.Println(f2(3))
}

func day29() {
	v := []int{1, 2, 3}
	for i := range v { // 此时尽管在for循环中一直在新增 v 切片的长度，但是并不会影响到原切片本身的循环
		fmt.Println(len(v))
		v = append(v, i)
		fmt.Println(v)
	}

	var m = [...]int{1, 2, 3}
	for i, v := range m { // 此时的 i, v 定义出来之后在for循环中会一直重复使用，而不是重新定义，所有它或打印出它最后一次被重用的值
		// 改正方法当然就是让重用编程重新定义，就可以完美解决
		key := i
		value := v
		go func() {
			fmt.Println(key, value)
		}()

		go func(x, y int) {
			fmt.Println(x, y)
		}(i, v)
	}
	time.Sleep(time.Second * 3)
}

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	fmt.Println(*p)
}

func day28() {
	// fmt.Println([...]int{1} == [2]int{1}) // 使用... 会自动推断数组大小，此时为1， 而后面的数组写死了2 所以类型不一致，不能比较大小
	fmt.Println([1]int{1} == [1]int{1}) // 数组就可以这样 进行比较，
	// 但是slice can only be compared to nil 切片只能与 nil 进行比较，如文章这样比较是不行的

	p, err := foo() // 此时p 与 全局的p 不在同一个作用域下，全局的p 就会被 局部变量p 屏蔽，
	// 所有在bar中的全局变量中读取值是，还是空的，所以回报nil runtime error 空指针引用错误

	if err != nil {
		fmt.Println(err)
		return
	}

	bar()
	fmt.Println(*p)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": &Math{2, 3},
}

func day27() {
	fmt.Println(South)

	m["foo"].x = 4

	fmt.Println(m["foo"].x)
}

func day26() {
	fmt.Println(a, b, c, d)
}

// defer 会按倒序执行，但是它所包含的函数中的函数 并不会，而是会直接执行
//

const (
	a = iota
	b = iota
)

const (
	name = "name"
	c    = iota
	d    = iota
)

func day23() {

	s1 := []int{1, 2, 3}
	s2 := s1[1:] // 使用这种方式产生的新切片会与原切片使用同一个底层数据结构，所有对s2修改会影响到s1
	s2[1] = 4

	fmt.Println(s1)
	fmt.Println(&s1)
	fmt.Println(&s2)
	s2 = append(s2, 5, 6, 7) // 但是通过append 方式重新返回的一个切片是一个全新的切片，对s2修改不会影响s1 // 底层数组也不会同一个了

	s2[2] = 4
	fmt.Println(s1)

	//var s string = nil // go 语言中 string 不能赋值为 nil ，也不能与nil 比较

	if a := 1; false {

	} else if b := 2; false {

	} else {
		fmt.Println(a, b)
	}
}

// 永远不需要使用 *interface{} 来接收 指针类型的值，因为 interface 本身已经就可以接收任何类型的值，包括指针类型
//

type S struct {
	m string
}

func f() *S {
	return &S{m: "xiongun"}
}

func day21() {
	p := f()
	fmt.Println(p.m)
}

type Person struct {
	age int
}

func day191() {
	person := &Person{28}

	defer fmt.Println(person.age)

	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}

func fa() (r int) { // 1
	defer func() {
		r++
	}()
	return 0
}

func fb() (r int) { // 5
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func fc() (r int) { //1 这里修改的值是 匿名喊出自己作用域内的变量，只是名字叫r 而已其实和外面函数中的r 一点关系都没有
	fmt.Println(r)
	defer func(i int) {
		i = i + 5
		fmt.Println(i)
	}(r)
	return 1
}

func day19() {
	fmt.Println(fa())
	fmt.Println(fb())
	fmt.Println(fc())

}

type Shape interface {
	Area() float32
}

type Object interface {
	Perimeter() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * (c.radius * c.radius)
}

func day181() {
	var s Shape = Circle{3}
	value1, ok1 := s.(Shape)
	value2, ok2 := s.(Object) // circle 根本没有实现 object 的接口，所有无法等到它的值，但不会编译报错，会返回nil 零值，flase

	fmt.Println(value1, ok1)
	fmt.Println(value2, ok2)
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}
func (w Work) ShowB() int {
	return w.i + 20
}

func (w Work) ShowC() int {
	return w.i + 30
}

func day18() {
	var x int

	x, y := f1()

	fmt.Println(x, y)

	var a A = Work{3}
	s := a.(Work) // 类型断言可以用户来获取接口的底层值，通常的语法是i.(type) i 必须是接口类型，这样就可以获得该结构体类型的所有实现方法

	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
	fmt.Println(s.ShowC())
}

func f1() (int, int) {
	return 1, 2
}

func day17() {
	s := [3]int{1, 2, 3}

	a := s[0:2]

	b := s[0:]                     // i 省略，默认为0
	fmt.Println(len(b), cap(b), b) //
	fmt.Println(len(s), cap(s))    //
	fmt.Println(len(a), cap(a), a) // s[0] <= a value < s[2] len 2-0= 2  cap = 3 - 0= 3
}

func day16() {

	var s1 []int     //  nil 切片 表示定义出来，为赋值// 表示一个不存在的切片
	var s2 = []int{} // 空切片，定义出来了，只是值是空的 // 空切片与 nil 不相等，表示一个空的集合

	fmt.Println(&s1)
	fmt.Println(s2)
	if s2 == nil {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(string(65))
}

func incr(p *int) int {
	*p++
	return *p
}

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func day15() {
	s1 := "hello"

	//s1[0] = 'x' // 字符串是只读的，不能被修改

	fmt.Println(s1)

	p := 1
	incr(&p)
	fmt.Println(p)

	add(1, 2)
	add(1, 2, 3)
	//add([]int{1, 2, 4})
	add([]int{1, 2, 4}...) // 可变参数

}

func hello(i int) {
	fmt.Println(i)
}

func day14() {
	i := 5
	defer hello(i) // defer语句虽说会在最后执行，但是在运行到该位置时，会先保存此时的副本，而此时 i= 5

	i += 10
	fmt.Println(i)
}

func day13() {

}

type people struct {
}

func (p *people) showB() {
	fmt.Println("showB")
}

func (p *people) showA() {
	fmt.Println("showA")
	p.showB()
}

type teacher struct {
	people
}

var str string

func (t *teacher) showB() {
	fmt.Println("teacher show b")
}

func day12() {

	i := -5
	j := +5

	fmt.Printf("%+d %+d", i, j)

	t := teacher{}
	t.showB()
}

func day11() {

	m := make(map[int]int)

	delete(m, m[1]) // 删除不存在的键时，并不会报错
	//ch := make(chan int)

	is := [5]int{1, 1, 2, 4, 5}

	fmt.Println(cap(is))
}
