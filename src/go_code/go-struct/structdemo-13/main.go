package main
import "fmt"

type A struct{
	Name string
}

type Goods struct{
	Name string
	Price float64
	A
}

type Book struct{
	Goods
	Publish string
	// Name string
}

type Computer struct{
	Goods
}

func (g Goods) test(){
	fmt.Println("Goods test",g.Name)
	fmt.Printf("g的数据类型是%T\n", g)	//g的数据类型是main.Goods，虽然是Book类型的实例调用，但是g的类型仍然是Goods
}

func (b Book) test(){
	fmt.Println("Book test",b.Name)
	fmt.Printf("b的数据类型是%T\n", b)
}

func main(){

	//定义方式一
	// var book = Book{Goods{"c++", 56.5}, "2009"}
	//定义方式二
	var book = Book{
		Goods:Goods{
			Name: "golang",
			Price: 92.3,
		},
		Publish: "2013",
		// Name: "book1",
	}
	fmt.Println(book)

	// fmt.Println(book.Goods.Name,book.Goods.Price)	//可以通过 实例.匿名结构体名称.字段访问匿名结构体中的属性
	fmt.Println(book.Name,book.Goods.Name,book.Price)	//如果实例中不存在与匿名结构体同名的字段，可以通过 实例.匿名结构 体字段名访问

	book.test()	

	var com = Computer{
		Goods{	//当Computer结构体中只有一个匿名结构体时，初始化时可以直接写:匿名结构体名{...}，如果还有其他有名字段，则必须写成:匿名结构体名:匿名结构体{...}
			Name: "huawei",
			Price: 12323.4,
			A:A{	//Goods结构体中有匿名结构体A，还有其他字段Name等，这里初始化时需写成：匿名结构体名:匿名结构体{...}
				Name: "v1",
			},
		},
	}
	fmt.Println(com.Name,com.A.Name)	//huwei,v1
}
