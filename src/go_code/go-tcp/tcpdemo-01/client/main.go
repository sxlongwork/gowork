package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("链接成功, ",conn)

	for {
		// 从标准输入获取数据
		reader := bufio.NewReader(os.Stdin)
		// 读取数据直到读到字符'\n'
		line, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("read data err.")
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit"{
			return
		}
		conn.Write([]byte(line))
	}
}