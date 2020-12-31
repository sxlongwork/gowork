package main
import (
	"fmt"
)

type Student struct{
	Name string
}

func judgeType(args... interface{}) {
	for _, v := range args {
		switch v.(type) {	//这里的type是固定写法，一种语法
		case bool:
			fmt.Printf("%v是bool类型\n", v)
		case string:
			fmt.Printf("%v是string类型\n", v)
		case int, int16, int32, int64:
			fmt.Printf("%v是整数类型\n", v)
		case float32:
			fmt.Printf("%v是float32类型\n", v)
		case float64:
			fmt.Printf("%v是float64类型\n", v)
		case nil:
			fmt.Println("是空类型")
		case Student:
			fmt.Println("Student类型", v)
		case *Student:
			if s, ok := v.(*Student); ok{
				fmt.Println("*Student类型", *s)
			}
			
		default:
			fmt.Printf("输入的类型不对\n")
		}
	
	}
}

func main(){
	var stu1 = Student{"tom"}
	judgeType(false, 3, 4.4, stu1, &stu1)
}