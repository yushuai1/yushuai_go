package main

import (
	"github.com/satori/go.uuid"
	"fmt"

	"strings"
)
func main() {

	u2, err :=  uuid.NewV4()

	if err==nil {

		fmt.Println(strings.Replace(u2.String(),"-","",-1))
	}


}
