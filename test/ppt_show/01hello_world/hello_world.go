package main

import "fmt"

func init() {
	fmt.Println("22222")
}

func main() {

	fmt.Println("Hello World!")

	u := user{}
	u.name = "main"
	u.getName()

	fmt.Println(u.name)

}

type user struct {
	name string
}

func (u user) getName() {
	u.name = "11111"
}

func (u *user) getName1() {
	u.name = "sss"
}
