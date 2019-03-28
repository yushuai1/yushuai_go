package main

import "fmt"

func main() {

	writeData :=[]string{"yu","shuai","shi","yi","ge","hao","ren"}
	memoryChan := make(chan string, 3)

	for _,msg := range writeData {
		select {
		case memoryChan <- msg:
			fmt.Println(msg,"-----------------")
		default:
			fmt.Println(msg,"*****************")
		}
	}
}
