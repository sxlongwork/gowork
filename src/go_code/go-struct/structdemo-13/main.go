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

func (g Goods) test(){
	fmt.Println("Goods test",g.Name)
}

func (b Book) test(){
	fmt.Println("Book test",b.Name)
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
}
