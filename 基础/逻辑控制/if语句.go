package main

import ("fmt"
         "math")

func sqrt(x float64) string  {

	fmt.Println(x)
	//if语句不需要（）但是内容还是需要{}的啊
	if x<0 {
		fmt.Println(x)
		return sqrt(-x)+"i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}


func main()  {
	fmt.Println(sqrt(2),sqrt(-3))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}