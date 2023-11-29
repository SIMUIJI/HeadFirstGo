package calendar

import "errors"

// 해당 구조체는 외부로 노출이 가능함
//type Date struct {
//	Year int
//	Month int
//	Day int
//}

// // 해당 구조체의 변수는 외부로 노출이 불가함 그래서 구조체의 변수를 외부에서는 선언이 불가함
type Date struct {
	year  int
	month int
	day   int
}

// setter
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// getter
func (d *Date) Year() int {
	return d.year
}
