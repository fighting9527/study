package main

import "fmt"

/**
任何Go语言源码文件都由若干个程序实体组成的。在Go语言中，变量、常量、函数、结构体和接口被统称为“程序实体”，而它们的名字被统称为“标识符”。

    标识符可以是任何Unicode编码可以表示的字母字符、数字以及下划线“_”。不过，首字母不能是数字或下划线。

    注意：在Go语言中，我们对程序实体的访问权限控制只能通过它们的名字来实现。名字首字母为大写的程序实体可以被任何代码包中的代码访问到。而名字首字母为小写的程序实体则只能被同一个代码包中的代码所访问。

    Go语言还规定了一些特定的字符序列。它们被称为“关键字”。编程人员不能把关键字作为标识符。Go语言的关键字如下表：

http://img.mukewang.com/55360c390001980911940300.jpg


 */

func main()  {
	fmt.Println("hello")

}