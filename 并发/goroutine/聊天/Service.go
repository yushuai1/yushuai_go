package yuaa

import (
	"net"
	"time"
	"log"
	"fmt"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err);
	}
}


//单独处理客户端的请求
func clientHandle(conn net.Conn) {
	//设置当客户端3分钟内无数据请求时，自动关闭conn
	conn.SetReadDeadline(time.Now().Add(time.Minute * 3));
	defer func() {
		fmt.Println(conn.RemoteAddr().String(),"客户端关闭了连接！")
		conn.Close();
	}()

	//循环的处理客户的请求
	for {
		data := make([]byte, 4);
		//从conn中读取数据
		conn.Read(data);
		lens:=BytesToUInt32(data)
		//发送给客户端的数据
		datastring := make([]byte, lens);
		conn.Read(datastring)

		hj:=BytesToStruct(datastring)
		hj.Name = "nidaye"

		//发送数据
		byes := StructToBytes(hj)
		//datastring = append(datastring,byes... )
		conn.Write(UInt32ToBytes(len(byes)));
		conn.Write(byes)
	}
}

func StartMain() {
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8084");
	chkError(err);
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr);
	chkError(err2);
	for {
		conn, err3 := tcplisten.Accept();

		if err3 != nil {
			continue;
		}else {
			fmt.Println("客户端连接成功！",conn.RemoteAddr())
		}
		go clientHandle(conn);
	}
}