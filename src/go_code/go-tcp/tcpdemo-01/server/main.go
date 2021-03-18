package main

import (
	"net"
	"fmt"
)

func HandleRequest(conn net.Conn){
	fmt.Printf("客户端%v已链接。\n", conn.RemoteAddr())
	for {	// 一直等待接收消息
		var b = make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil{
			fmt.Printf("客户端退出\n",)
			return
		}
		fmt.Print(string(b[:n])+"\n")
	}
	
	
}

func main(){
	// 监听
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("sever已开始监听, 等待客户端链接......\n")
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println(err)
			// 如果出错，继续等待下一个链接请求，不能break
			continue
		}
		
		go HandleRequest(conn)


	}
}