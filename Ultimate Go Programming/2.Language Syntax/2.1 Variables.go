package __Language_Syntax

import "fmt"

var QUOTE21 = []string{
	"1.Type is everything, it determine the meaning of bytes",
}

func Demo() {
	var a int //声明不初始化,默认为0值
	var b string
	var c float64
	var d bool
	fmt.Printf("var a int \t %T [%v] \n", a, a)
	fmt.Printf("var b string \t %T [%v] \n", b, b)
	fmt.Printf("var c float64 \t %T [%v] \n", c, c)
	fmt.Printf("var d bool \t %T [%v] \n", d, d)
	aa := 10 //声明且初始化
	bb := "hello"
	cc := 3.14159
	dd := true
	fmt.Printf("aa := 10 \t %T [%v] \n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v] \n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v] \n", cc, cc)
	fmt.Printf("dd := true \t %T [%v] \n", dd, dd)
}
func WatchOut() {
	//Go有强制类型转换transform，但是没有隐式类型转换casting，除非被转换对象没有显示声明类型
	//casting是原地转换，极容易造成内存越界异常
	//Go选择显示声明转换类型，得到的新值额外开辟一块内存空间，不会对原值有影响
	var a int
	var b int64 = a        //a显式声明为int 不能被显示转换为int64
	var b int64 = int64(a) //a可以被强制类型转换
	b = 10                 // 10没有声明类型，可以隐式转换成int64
}
