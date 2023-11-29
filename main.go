package main

type MyType string

func (m MyType) sayHi() {
	fmt.println("HI")
}

func main() {
	value := MyType("test")
	value.sayHi()
}
