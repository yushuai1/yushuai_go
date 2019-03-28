package main

import (
	"fmt"
	//"net"
	//"os"
	"strings"
	//"crypto/aes"
	"strconv"
	"net"
	"os"
)


func GetIntranetIp() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ip:", ipnet.IP.String())
			}

		}
	}
}

func GroupIdToArray(content string) []int32 {
	stringArray := strings.Split(content, "#")
	array := make([]int32, len(stringArray))
	for k, v := range stringArray {
		b, _ := strconv.Atoi(v)
		array[k] = int32(b)
	}

	return array
}


func GetPulicIP() string {
	conn, _ := net.Dial("udp", "1.5.7.5:555")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}
func main() {

	//GetIntranetIp()
	fmt.Println()
	fmt.Println(GetPulicIP())
	//var yu string = "456#123#566"
	//sm:=GroupIdToArray(yu)
	//for k, v := range sm {
	//	fmt.Println(k, v)
	//}
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {


		// 检查ip地址判断是否回环地址
		ipnet, ok := address.(*net.IPNet)

		if  ok && !ipnet.IP.IsLoopback(){

			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}
