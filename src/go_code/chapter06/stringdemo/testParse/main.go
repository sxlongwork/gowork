package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"flag"
)

type Student struct {
	Name string	`xml:"name"`
	Age int	`xml:"age"`
}
func xmlParse(){
	p := Student{"jenny", 19}
	var data []byte
	data, err := xml.MarshalIndent(p, "", "    ");
	if  err == nil {
		fmt.Println(string(data))
	}
	 p2 := new(Student)
	 if err := xml.Unmarshal(data, p2); err == nil{
		fmt.Println(p2)
	 }
}

func argsParse(){
	fmt.Println(os.Args)
}

func flagParse(){
	n := flag.String("name", "", "username")
	var num int
	flag.IntVar(&num, "number", 0, "input number")
	
	flag.Parse()
	fmt.Println(*n)
	fmt.Println(num)
}


func main(){
	// xmlParse()
	// argsParse()
	flagParse()
}