
package main

import (
	"errors"
	"fmt"
	"time"
)

// ===================== 1. 常量 & iota 枚举 =====================
const (
	_ = iota
	RED
	GREEN
	BLUE
	PI    = 3.1415926
	NAME  = "Go全语法示例"
)

// 类型别名
type MyInt int
type Str string

// ===================== 2. 结构体 & 方法 =====================
// 定义结构体
type Person struct {
	Name string
	Age  int
	// 小写私有
	addr string
}

// 结构体方法 值接收者
func (p Person) SayHi() {
	fmt.Printf("我是 %s，年龄 %d\n", p.Name, p.Age)
}

// 结构体方法 指针接收者
func (p *Person) SetAge(age int) {
	p.Age = age
}

// ===================== 3. 接口 =====================
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪汪"
}

// ===================== 4. 函数、多返回值、错误返回 =====================
func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 可变参数
func sum(nums ...int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

// 匿名函数 & 闭包
func closureDemo() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// ===================== 5. defer、panic、recover =====================
func panicDemo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()
	panic("手动触发panic")
}

// ===================== 6. 协程 & Channel =====================
func worker(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	// ===================== 变量定义 =====================
	var a int = 10
	var b string = "hello"
	c := 3.14
	var d bool = true
	fmt.Println("变量：", a, b, c, d)

	// 基础数据类型
	var (
		ui  uint  = 100
		i64 int64 = 123456
		by  byte  = 'A'
		r   rune  = '中'
	)
	fmt.Println("基础类型：", ui, i64, by, r)

	// ===================== 数组 & 切片 =====================
	// 数组
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("数组：", arr)

	// 切片
	s := []int{10, 20, 30}
	s = append(s, 40)
	fmt.Println("切片：", s)

	// ===================== Map =====================
	m := make(map[string]int)
	m["语文"] = 90
	m["数学"] = 95
	fmt.Println("Map：", m)

	// ===================== 指针 =====================
	x := 100
	px := &x
	fmt.Printf("指针地址：%p，值：%d\n", px, *px)

	// ===================== 流程控制 if / for / switch =====================
	// if
	if a > 5 {
		fmt.Println("if 条件成立")
	} else {
		fmt.Println("if 条件不成立")
	}

	// for 三种写法
	for i := 0; i < 2; i++ {
		fmt.Println("for循环", i)
	}

	k := 0
	for k < 2 {
		fmt.Printlnfor 条件循环, k
		k++
	}

	// 无限循环
	// for {}

	// switch
	num := 2
	switch num {
	case 1:
		fmt.Println("case1")
	case 2:
		fmt.Println("case2")
	default:
		fmt.Println("default")
	}

	// ===================== 结构体使用 =====================
	p := Person{Name: "张三", Age: 20, addr: "北京"}
	p.SayHi()
	p.SetAge(25)
	p.SayHi()

	// ===================== 接口使用 =====================
	var spt Speaker = Dog{}
	fmt.Println("接口调用：", spt.Speak())

	// ===================== 函数调用 =====================
	fmt.Println("add(3,5) =", add(3, 5))
	res, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("div结果：", res)
	}
	fmt.Println("可变参数sum：", sum(1, 2, 3, 4))

	// 闭包
	f := closureDemo()
	fmt.Println("闭包1：", f())
	fmt.Println("闭包2：", f())

	// ===================== defer 延迟调用 =====================
	defer fmt.Println("main 最后执行defer")

	// ===================== panic recover =====================
	panicDemo()

	// ===================== 协程 Channel =====================
	ch := make(chan int)
	go worker(ch)
	for v := range ch {
		fmt.Println("channel 接收：", v)
	}

	// ===================== 常量 iota =====================
	fmt.Println("常量iota：", RED, GREEN, BLUE)
}
