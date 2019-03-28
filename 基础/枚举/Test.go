package main

import "fmt"

const(
	AE_PARAMETER = iota + 101
	AE_AEECSS_TOKEN = iota + 101
	AE_INTERFACE_ACCESS
	AE_ERROR_LOG
	AE_SUCCEE
	AE_EXCEPTION
	AE_FAILED
)

func main()  {

	fmt.Println(AE_AEECSS_TOKEN)
}
