package main

import "fmt"

//type MyType string
//
//// 메서드를 호출하고 있는 값을 메서드 리시버 라고 합니다.
//// 앞에 m은 리시버로 어떤 값을 받아올 것인지를 명시하는 곳 해당 MyType은 포인터로도 변경가능
//func (m MyType) sayHi() {
//	fmt.Println("HI")
//}
//
//func main() {
//	value := MyType("test")
//	value.sayHi() // value가 메서드 리시버
//}

type Liter float64
type Milliliters float64
type Gallons float64

func (l Liter) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func main() {
	soda := Liter(2)
	fmt.Printf("%0.3f liter equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}
