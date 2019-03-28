package 接口

import (

	"fmt"

	//"yushuai/基础/函数"
)


//
//func main (){
//
//	var ming 函数.People
//
//	ming = &Chinese{Name : "ming", Energy : 10}
//
//	fmt.Println(ming)
//
//	ming.Eat(5)
//
//	ming.Sleep(3)
//
//	ming.Work(-8)
//
//	fmt.Println(ming)
//
//}





type Chinese struct{

	Name string

	Energy int

}

func (s *Chinese) Sleep(e int) bool {

	s.Energy = s.Energy + e

	fmt.Printf("Chinese %s Sleep, Energy = %d\r\n", s.Name, s.Energy)

	return true

}

func (s *Chinese) Eat(e int) bool {

	s.Energy = s.Energy + e

	fmt.Printf("Chinese %s Eat, Energy = %d\r\n", s.Name, s.Energy)

	return true

}

func (s *Chinese) Work(e int) bool {

	if s.Energy > e {

		s.Energy = s.Energy - e

		fmt.Printf("Chinese %s Work, Energy = %d\r\n", s.Name, s.Energy)

		return true

	}

	fmt.Printf("Chinese %s can not Work, Energy = %d\r\n", s.Name, s.Energy)

	return false

}
