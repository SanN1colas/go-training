package __Language_Syntax

import "fmt"

type example struct {
	flag    bool    //1 byte
	counter int16   //2 bytes
	pi      float32 //4 bytes
}

func Demo() example {
	var e example
	//How many bytes that e1 was allocated?
	//1 + 2 + 4 ? No.
	//It depends on the arch of system. Under x64, each word is specified by 8 bytes(a word is a minimum unit that a CPU
	//will take from memory once)
	//1.for less fetch operations, system will try to align the structs as much as possible.
	//2.each struct will try not to cross two word for less operations, like we do not expect a 2-byte struct in two word like
	//[_______1][2_______]
	//So each struct will be aligned to a proper start location like
	//[1_22____],even the 2nd byte is waste, the 2-byte struct will still start from the 3rd pos
	//So It need 8 bytes
	//[1_224444]
	fmt.Printf("%v", e)

	//5 ways to declare a struct
	var e1 example
	e2 := example{}
	e3 := example{
		flag:    false,
		counter: 0,
		pi:      0,
	}
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{flag: false,
		counter: 0,
		pi:      0,
	}
	return example{
		flag:    false,
		counter: 0,
		pi:      0,
	}
}
func WatchOut() {
	type example struct {
		flag     bool    //1
		counter  int64   //8
		flag2    bool    //1
		counter2 int64   //8
		flag3    bool    //1
		counter3 int64   //8
		pi       float32 //4
	}
	e := example{}
	//TODO 1.How many memory bytes allocated for e?
	//[1_______][88888888][1_______][88888888][1_______][88888888][4444____]
	//We need extra 7*3 bytes to store e (the last 4 bytes could be used for others)
	//TODO 2.How to reduce the extra bytes?
	//Arrange the struct by its size from big to small
	type example2 struct {
		counter  int64   //8
		counter2 int64   //8
		counter3 int64   //8
		pi       float32 //4
		flag     bool    //1
		flag2    bool    //1
		flag3    bool    //1
	}
	e2 := example2{} //[88888888][88888888][88888888][4444111_] We need 0 extra bytes to store e2
	//TODO 3.Should we do the sorting for every struct?
	//No, it will damage the readiness, we satisfy some cost for more efficiency in readiness.
}
func WatchOut2() {
	//type conversion
	type Nico struct {
		Name string
		Age  int
		Sex  bool
	}
	type Mia struct {
		Name string
		Age  int
		Sex  bool
	}
	var mia Mia
	var nico Nico
	nico = mia       //TODO 1.nico可以被转换成Mia类型么？不可以，因为mia是显示声明成Mia type的struct，而不是Nico，即使二者结构完全相同，但也不可以直接赋值
	nico = Nico(mia) //TODO 2.可以强制类型转换，因为变量名和类型完全相同
	nipan := struct {
		Name string
		Age  int
		Sex  bool
	}{}
	nipan2 := struct {
		Age  int
		Sex  bool
		Name string
	}{}
	//TODO 3.nipan可以赋值给nico么？可以
	nico = nipan
	//TODO 4.nipan2可以赋值给nico么？不可以，顺序也不能变
	nico = nipan2
}
