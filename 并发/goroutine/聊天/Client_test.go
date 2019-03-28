package yuaa

import (
	"testing"
	"fmt"
)

type Myt struct {
	Age int
	Name string
}

func TestClient_NewClient(t *testing.T) {

	cli :=new(Client)
	cli.NewClient()

	bn :=&My{
		Age:10,Name:"shuai",
	}

	data := StructToBytes(bn)

	re :=cli.Sender(data)

	fg:=BytesToStruct(re)

	fmt.Println(fg)


}