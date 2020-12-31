package main
import (
	"fmt"
)

type Usb interface{
	start()
	stop()
}

type Phone struct{
	Name string
}

type Camera struct{
	Name string
}

type Computer struct{

}

func (p Phone) start(){
	fmt.Println(p.Name,"--phone开始--")
}

func (p Phone) stop(){
	fmt.Println(p.Name, "--phone停止--")
}

func (p Phone) call(){
	fmt.Println(p.Name, "------phone在打电话------")
}

func (c Camera) start(){
	fmt.Println(c.Name,"--Camera开始--")
}

func (c Camera) stop(){
	fmt.Println(c.Name,"--Camera停止--")
}

// 根据传入参数的不同，调用不同实例的方法；
//	 传入Phone实例，调用Phone实例的方法；
//   传入Camera实例，调用camera实例的方法
func (com Computer) work(usb Usb){	
	//通过usb来调用start和stop方法
	usb.start()
	usb.stop()
}

func main(){
	var p1 = Phone{"huwei mate"}
	var ca = Camera{"dor"}

	var computer Computer
	computer.work(p1)
	computer.work(ca)
	fmt.Println("-----------------------------")

	var arr  [3]Usb
	var p2 = Phone{"荣耀9X"}
	arr[0] = p1
	arr[1] = p2
	arr[2] = ca
	for _, v := range arr{
		// 类型断言，如果v是Phone实例，则ok为true，v.(Phone)转化为了Phone类型变量p；如果ok为false,继续执行后面的代码
		if p, ok := v.(Phone); ok{
			p.call()
		}
		v.start()
		v.stop()

	}
	fmt.Println("-----------------------------")

	var str string
	// var str int	// 有类型断言，即使定义str为int类型，程序也可以正常执行
	var inter interface{}	//空接口可以接收所有类型变量
	inter = str

	// var newStr string
	//如果inter不是指向string类型，则此处执行时会报panic退出，后面的代码也不会执行了
	// newStr = inter.(string)	

	// 添加类型断言，即使转换失败，也可以继续执行后续的代码
	if newStr, ok := inter.(string); ok{
		fmt.Println(len(newStr))
	} else{
		fmt.Println("继续执行后续代码...")
	}
	


}