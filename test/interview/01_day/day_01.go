package main

import "fmt"

func main() {

	//defer_call()

	//ms := MyString("SeekLoad")
	//r := Rect{5.0, 4.0}
	//
	//explain(ms)
	//explain(r)

	//day02()

	//day10()
}

func day101() {

}

type person struct {
	name string
}

func day10() {

	//var ch chan int
	//
	////ch := make(chan int)
	//ch <- 1
	//s := <-ch
	//fmt.Println(s)

	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p]) // 从map中取出一个不存在的值，会返回该类型的零值，int 的零值是 0

	i := []int{5, 6, 7}
	hello2(i) // 可变参数其实就是切片

	fmt.Println(i[0])
}

func hello2(num []int) {
	num[0] = 19 // 此时可变参数就是切片，而且直接会将切片传入函数，并不会创建新的切片，所以，对切片的修改就是会改变切片内的值
}

func hello() []string {
	return nil
}

func GetValue() interface{} {
	return 1
}

func day09() {
	h := hello // 是将Hello() 整个函数赋值给了H,其实是把函数所在的地址返回了，所有不会是nil， hello() 才是正常的函数调用，这时候才是nil
	fmt.Println(h)
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	i := GetValue()
	// 只有interface 类型才能进行 i.(type) 判断
	switch i.(type) {
	case int:
		fmt.Println("int")
	}

}

const (
	x = iota // 0
	_        //1
	y        //2
	z = "zz" // 3 "zz"
	k        // 4 "zz"
	p = iota //5
)

func day08() string {
	fmt.Println(x, y, z, k, p)

	//var x = nil
	//var x interface{} = nil
	//var x string
	//x = ""
	//
	//var x error = nil
	return ""
}

type user struct {
	name string
	age  int
}

type MyInt1 int
type MyInt2 = int

func day07() {
	u := user{name: "xiongdun", age: 11}

	(&u).name = "sss"
	fmt.Println(u.name)
	//fmt.Println((&u).name)
	//fmt.Println((*u).name)

	//var i int = 0
	//var i1 MyInt1 = i // 这编译不通过，大概意思就是说 int 类型不能赋值于MyInt1
	// ，因为MyInt1 虽说是基于int而来的类型，但是是一个全新的类型，go是强类型的语言，所有不能编译通过。不过可以使用强制类型转换 var i1 Myint = MyInt1(1)
	// var i2 MyInt2 = i // 编译通过，因为使用 = 符号表示的是将MyInt2 定义为了int 类型的别名，
	// 只是别名而已，本质还是int 所以可以编译通过
}

func day062() {
	m1 := make(map[string]string)

	m1["a"] = "1"

	m2 := make(map[string]string)
	m2["a"] = "1"

	//(map can only be compared to nil
	if m1 == nil {

	}

	if m2 == nil {

	}
}

// invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
// slice, map ,function 不能被比较
func day06() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "xiongdun"}

	sn2 := struct {
		age  int
		name string
	}{name: "xiongdun", age: 11}

	if sn1 == sn2 {
		fmt.Println("sn1==sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//sm2 := struct {
	//	age int
	//	m map[string]string
	//}{age:11, m: map[string]string{"a": "1"}}
	//
	//if sm1 != nil {
	//	fmt.Println("sm1==sm2")
	//
	//}
}

func day05() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}

	s1 = append(s1, s2...)

	fmt.Println(s1)
}

func day03() {
	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)

	s2 := make([]int, 0)
	int2s := append(s2, 1, 2, 3, 4)
	fmt.Println(int2s)

}

func funcMui(x, y int) (sum int, err error) {
	return x * y, nil
}

// new and make difference
// new 返回的是指针，适用于值类型，数组，结构体等
// make 返回的是经初始化后的类型T的引用，只适用于slice, map channel

func day02() {
	slice := []int{0, 1, 2, 3}

	// 取得指针地址上得值，因为只有一个指向value 得指针每次都把该地址上得值替换掉了，所以所有
	m := make(map[int]*int)

	for key, value := range slice {

		val := value
		m[key] = &val

		fmt.Println(&val)
	}

	for key, value := range m {
		fmt.Println(key, "->", *value)
	}
}

// defer 的执行顺序是后进先出，当出现panic 语句的时候，
// 会先按照defer 的后进先出的顺序执行，最后才会执行panic
func defer_call() {
	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

type MyString string
type Rect struct {
	width  float32
	height float32
}

func explain(i interface{}) {
	fmt.Printf("type of s is %T\n", i)
	fmt.Printf("value of s is %v\n", i)
}
