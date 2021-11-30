/**
 * @author fire9
 * @email fire9@me.com
 * @create date 2021-11-30 15:55:28
 * @modify date 2021-11-30 15:55:28
 * @desc 说明一下这个是多行注释
 */
package main

import "fmt"

const name string = "Hank" // 常量必须赋值

var age int = 44

var (
	num      int
	flo      float64
	isMan    bool
	nickName string
)

func sayHello() {
	println("Hello, " + name + "!")
	fmt.Printf("I'm %d years old.\n", age)
}

func zeroValue() {
	fmt.Printf("整数的零值是 %d\n", num)
	fmt.Printf("浮点数的零值是 %f\n", flo)
	fmt.Printf("布尔的零值是 %v\n", isMan)
	fmt.Printf("字符串的零值是 %q\n", nickName)
}

func equalType() {
	var num1 int32
	var num2 rune
	if num1 == num2 {
		fmt.Println("rune是int32的别名")
	}
}

func main() {
	sayHello()
	msg := "Good 再见!"
	fmt.Println(msg)
	zeroValue()
	equalType()
}
