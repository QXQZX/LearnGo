package m_interface

import (
	"fmt"
)

/*
关于实现接口：
1. 一个struct可以实现多个接口
2. 接口间可以相互组合
3. 接受者类型必须保持一致，要么全是指针接收者  要么全是值接收者
*/

type Cat struct {
	ID   int
	Name string
	Age  int
}

func (c *Cat) Eat() string {
	fmt.Println("cat eating")
	return ""
}

func (c *Cat) Run() string {
	fmt.Println("cat Running")
	return ""
}

type Dog struct {
	ID   int
	Name string
	Age  int
}

func (d *Dog) Eat() string {
	fmt.Println("dog eating")
	return ""
}

func (d Dog) Run() string {
	fmt.Println("dog Running")
	return ""
}

type mPig struct {
}

// 实现stringer 相当于java的toString 方法
func (m mPig) String() string {
	// fmt.Println(bp)会打印返回值
	return "{mPig:...}"
}

func (m mPig) Run() string {
	fmt.Println("mPig running")
	return ""
}

func (m mPig) Eat() string {
	fmt.Println("mPig eating")
	return ""
}

func (m mPig) Sleep() string {
	fmt.Println("mPig sleeping")
	return ""
}

func (m mPig) Drink() string {
	fmt.Println("mPig drinking")
	return ""
}

func AnimalDo(a Animal) {
	a.Eat()
	a.Run()
}

func PigDo(bp BigPig) {
	bp.Eat()
	bp.Run()
	bp.Drink()
	bp.Sleep()
}

func TestInterface() {
	cat := Cat{
		ID:   1,
		Name: "11",
		Age:  0,
	}

	dog := Dog{
		ID:   2,
		Name: "22",
		Age:  0,
	}

	var bp BigPig
	bp = mPig{}
	fmt.Printf("%T, %v\n", bp, bp)

	var p Pig
	p = bp
	fmt.Printf("%T, %v\n", p, p)

	switch v := p.(type) {
	case *mPig:
		fmt.Println("*mPig", v)
	case mPig:
		fmt.Println("mPig")
	default:
		break
	}

	/**
	指针类型可以调用  值接收器和指针接收器的函数；
	值类型也可以调用  值接收器和指针接收器的函数；

	类型实现接口中的函数 只要有一个函数用了指针接收器类型，之后就只能将该类型的地址赋值给该接口
	或者说，结构体变量不能赋给用指针接收器实现的接口
	*/

	//var dogg Animal = dog // dog实现了Animal接口中，Eat用指针接收器实现，Run用值接收器，所以只能将地址赋值给Animal接口

	AnimalDo(&cat) // 指针接收器必须传地址
	AnimalDo(&dog)
	AnimalDo(bp)

	fmt.Println(bp) // 实现了stringer 相当于java的toString 方法
}

//下面方法只适合于interface
//xx.(type)只能用于switch中
//xx.(int)强制转成int类型
func TypeAssert() {
	var q Animal = mPig{}
	// 范围大  ->  范围小的
	switch v := q.(type) {
	case mPig:
		fmt.Println("mPig", v)
	case *mPig:
		fmt.Println("*mPig", v)
	}

	//if s, ok := q.(int); ok {
	//	fmt.Println(s)
	//} else {
	//	fmt.Println("err")
	//}
}
