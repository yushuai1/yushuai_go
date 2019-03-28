package yuaa

import (
	"fmt"
	"net"
	"os"
	"time"
)

func (p *Client)Sender(sentData []byte) []byte{

	    data := make([]byte, 4);
        p.connCli.Write(UInt32ToBytes(len(sentData)))
		p.connCli.Write(sentData)
		//从conn中读取数据
		p.connCli.Read(data);
		lens:=BytesToUInt32(data)
		//发送给客户端的数据
		datastring := make([]byte, lens);
		p.connCli.Read(datastring)
		return datastring
}


type Client struct {
	connCli *net.TCPConn
}

func (p *Client)NewClient()  {
	server := "127.0.0.1:8084"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.SetNoDelay(true)
	conn.SetKeepAlive(true)
	p.connCli = conn
}

func main() {
	server := "127.0.0.1:8084"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.SetNoDelay(true)
	conn.SetKeepAlive(true)
	defer func() {

		conn.Close()
		fmt.Println("客户端关闭！")
	}()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")

	//go sender(conn)
	//go receive(conn)

	time.Sleep(time.Second*10)
}