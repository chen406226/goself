package hello

import (
	"fmt"
	"strconv"
)
/*
	var e, f = 123, "hello"
	//这种不带声明格式的只能在函数体中出现
	//g, h := 123, "hello"
------------
const 也可以定义枚举
import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)
abc 3 16
---------
=======
func main() {
  _,numb,strs := numbers() //只获取函数返回值的后两个
  fmt.Println(numb,strs)
}

//一个可以返回多个值的函数
func numbers()(int,int,string){
  a , b , c := 1 , 2 , "str"
  return a,b,c
}
=======

*/

func show(){
	/*
	1. 布尔型
	2. 数字类型
	3. 字符串类型
	4. 派生类型 (指针类型Pointer)(数组类型)（接口类型interface）
	(结构化类型struct)（Channel类型）（函数类型）（切片类型）（Map类型）
	一，默认值 var a int,string,bool
	数值类型 0 布尔类型false 字符串 ""
	nil:  类型
	var a pointer
	var a *int
	var a []int
	var a map[string] int
	var a chan int
	var a func(string) int
	var a error //error是接口

	声明变量  可以用  key := value 这个key是新生成的不能有旧的
	//类型相同多个变量, 非全局变量
	var vname1, vname2, vname3 type
	vname1, vname2, vname3 = v1, v2, v3

	var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

	vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
	var a, b int
	var c string
	a, b, c = 5, 7, "abc"

	a, b, c := 5, 7, "abc"

	// 这种因式分解关键字的写法一般用于声明全局变量
	var (
	    vname1 v_type1
	    vname2 v_type2
	)
	// const 定义常量（布尔，数字，字符串）不能定义引用型iota特殊用处初始未0
	 const (
	            a = iota   //0
	            b          //1
	            c          //2
	            d = "ha"   //独立值，iota += 1
	            e          //"ha"   iota += 1
	            f = 100    //iota +=1
	            g          //100  iota +=1
	            h = iota   //7,恢复计数
	            i          //8
	    )

	其他运算符 & 返回变量存储地址 * 指针变量 &a 是地址*&a是值

	// 条件语句   if   switch   select
	//for 循环有三种形式 for init; condition; post { }
	类似while    for condition { }
	for {}
	for 循环的range格式可以对slice(切片),集合map,数组(array),通道(channel),字符串进行迭代循环格式：返回key,value对
	for key or _, value := range oldMap {
		newMap[key] = value
	}
	    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	    for i, c := range "go" {
	        fmt.Println(i, c)
	    }
		0 103
		1 111
	for 的控制语句有三个   break   continue  goto将控制转移到被标记的语句
	goto 不在for里面也可以使用
	LOOP: if a < 15 {
		a = a + 1
		fmt.Printf("a的值为 : %d\n", a)
		goto LOOP
	}

	//函数func max(num1, num2 int) int {}
	//func max(num1, num2 int) (int,string) {}
	//传参 1，  值传递，传入值不会随函数执行而改变，基础数据类型，
	2， 引用传递： 再调用函数时讲实际参数的地址传递到函数中，那么再函数中对参数所进行的修改，将影响到实际参数。
	// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 方法，一个方法就是包含了接收者的函数，接收者可以是明明类型活着结构体类型的一个值活着是一个指针。
	所有给定类型的方法属于改类型的方法集
	// 定义结构体
	type Circle struct {
		radius float64
	}

	func main() {
		var c1 Circle
		c1.radius = 10.00
		fmt.Println("圆的面积 = ", c1.getArea())
	}

	//该 method 属于 Circle 类型对象中的方法
	func (c Circle) getArea() float64 {
		//c.radius 即为 Circle 类型对象中的属性
		return 3.14 * c.radius * c.radius
	}

	//数组
	Go 语言数组声明需要指定元素类型及元素个数  var variable_name [SIZE] variable_type
	var balance [10] float32
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	读取数组中的元素， 可以通过索引类读取ep:  var salary float32 = balance[9]
	
	//多维数组var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	var three [5][10][4]int
	a = [3][4]int{
	 	{0, 1, 2, 3} ,   /  第一行索引为 0 /
		{4, 5, 6, 7} ,   /  第二行索引为 1 /
		{8, 9, 10, 11},   / 第三行索引为 2 /
	}
	val := a[2][3]
	//Go中数组长度是不可变的，Slice切片是灵活的数组，可以追加元素，再追加时可能使切片的容量增大
	var identifier []type
	切片不需要说明长度，或使用make()函数来创建切片
	var slice1 []type = make([]type, len)
	也可以简写为
	slice1 := make([]type, len)
	也可以指定容量，其中capacity为可选参数
	make([]T, length, capacity)
	s :=[] int {1,2,3 }

	var numbers []int len=0 cap=0 slice=[] 切片是空的
	   if(numbers == nil){
	      fmt.Printf("切片是空的")
	   }
	len()获得长度  cap()函数 append() 增加切片容量 copy拷贝切片




		//指针
	var  ptr *int
	ptr 的值为 : 0
	指针判断
	if(ptr != nil)     / ptr 不是空指针 /
	if(ptr == nil)    / ptr 是空指针 /

	// 使用函数
	fmt.Println(getSquareRoot(9))

	============================结构体
	结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。
	type struct_variable_type struct {
	   member definition
	   member definition
	   ...
	   member definition
	}
	一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
	variable_name := structure_variable_type {value1, value2...valuen}
	或
	variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
	type Books struct {
	   title string
	   author string
	   subject string
	   book_id int
	}
	func main() {

	    // 创建一个新的结构体
	    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	    // 也可以使用 key => value 格式
	    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	    // 忽略的字段为 0 或 空
	   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	   var Book1 Books        //声明Book1为Books类型

		Book1.title = "Go语言"f
		Book1.author = "www.runoob.com"
		Book1.subject = "Go 语言教程"
		Book1.book_id = 6495407

		// 打印 Book1 信息
		fmt.Printf( "Book 1 title : %s\n", Book1.title)
		fmt.Printf( "Book 1 author : %s\n", Book1.author)
		fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
		fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
		}
		//res
		{Go 语言 www.runoob.com Go 语言教程 6495407}
		{Go 语言 www.runoob.com Go 语言教程 6495407}
		{Go 语言 www.runoob.com  0}

	结构体是作为参数的值传递：
	如果想在函数里面改变结果体数据内容，需要传入指针：
	changeBook(&book1)
	func changeBook(book *Books) {
	    book.title = "book1_change"
	}

	====================
	Map集合
	/ 声明变量，默认 map 是 nil
	var map_variable map[key_data_type]value_data_type

	/ 使用 make 函数
	map_variable := make(map[key_data_type]value_data_type)
	如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

	    var countryCapitalMap map[string]string /创建集合
	countryCapitalMap = make(map[string]string)

	/ map插入key - value对,各个国家对应的首都
	countryCapitalMap [ "France" ] = "巴黎"
	capital, ok := countryCapitalMap [ "American" ]
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	// interface 接口 所有具有共性的方法定义在一起，任何其他类型只要实现这些方法就是实现了这个接口。
	/ 定义接口
	type interface_name interface {
		method_name1 [return_type]
		method_name2 [return_type]
		method_name3 [return_type]
		method_namen [return_type]
	}
	/定义结构体
	type struct_name struct {
	/variables
	}
	/ 实现接口方法
	func (struct_name_variable struct_name) method_name1() [return_type] {
	/方法实现
	}
	...
	func (struct_name_variable struct_name) method_namen() [return_type] {
	 方法实现
	}
		EP:

	value, ok = element.(T)
	if value, ok := element.(int); ok {}
	*/

}


func main() {

	







	//var _s bool = true
	v := 334
	fmt.Printf("Hello,world\n")
	//fmt.Scan(&v)
	fmt.Printf(strconv.Itoa(v))
	fmt.Printf("Hello,world\n")


	// Switch  也可以判断类型type-switch 来判断某个 interface 变量中实际存储的变量类型switch x.(type){
	//使用 fallthrough 会强制执行后面的《一条》 case 语句
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T",i)
		//fallthrough
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
	}
}
