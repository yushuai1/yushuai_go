package 接口

import (
	"fmt"
	//"yushuai/基础/函数"
)

type English struct {
	Name string
	Age int
}

//func main() {
//
//	var english 函数.People
//	english = &English{Name:"asdasd",age:15}
//
//	english.Eat(5)
//	english.Sleep(10)
//	english.Work(121)
//
//}

func (s *English) Sleep(e int) bool {

	s.Age = s.Age * e

	fmt.Printf("English %s Sleep, age = %d\r\n", s.Name, s.Age)

	return true

}

func (s *English) Eat(e int) bool {

	s.Age = s.Age + e

	fmt.Printf("English %s Eat, age = %d\r\n", s.Name, s.Age)

	return true

}

func (s *English) Work(e int) bool {

	if s.Age > e {

		s.Age = s.Age - e

		fmt.Printf("English %s Work, age = %d\r\n", s.Name, s.Age)

		return true

	}

	fmt.Printf("English %s can not Work, age = %d\r\n", s.Name, s.Age)

	return false

}
