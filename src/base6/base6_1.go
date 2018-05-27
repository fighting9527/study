package main

import (
	"sync/atomic"
	"strconv"
	"fmt"
)

/**
Go语言-函数
    在Go语言中，函数是一等（first-class）类型。这意味着，我们可以把函数作为值来传递和使用。函数代表着这样一个过程：它接受若干输入（参数），并经过一些步骤（语句）的执行之后再返回输出（结果）。特别的是，Go语言中的函数可以返回多个结果。

    函数类型的字面量由关键字func、由圆括号包裹参数声明列表、空格以及可以由圆括号包裹的结果声明列表组成。其中，参数声明列表中的单个参数声明之间是由英文逗号分隔的。每个参数声明由参数名称、空格和参数类型组成。参数声明列表中的参数名称是可以被统一省略的。结果声明列表的编写方式与此相同。结果声明列表中的结果名称也是可以被统一省略的。并且，在只有一个无名称的结果声明时还可以省略括号。示例如下：

func(input1 string ,input2 string) string
    这一类型字面量表示了一个接受两个字符串类型的参数且会返回一个字符串类型的结果的函数。如果我们在它的左边加入type关键字和一个标识符作为名称的话，那就变成了一个函数类型声明，就像这样：

type MyFunc func(input1 string ,input2 string) string
    函数值（或简称函数）的写法与此不完全相同。编写函数的时候需要先写关键字func和函数名称，后跟参数声明列表和结果声明列表，最后是由花括号包裹的语句列表。例如：

func myFunc(part1 string, part2 string) (result string) {
    result = part1 + part2
    return
}
    我们在这里用到了一个小技巧：如果结果声明是带名称的，那么它就相当于一个已被声明但未被显式赋值的变量。我们可以为它赋值且在return语句中省略掉需要返回的结果值。至于什么是return语句，我就不用多说了吧。显然，该函数还有一种更常规的写法：

func myFunc(part1 string, part2 string) string {
    return part1 + part2
}
    注意，函数myFunc是函数类型MyFunc的一个实现。实际上，只要一个函数的参数声明列表和结果声明列表中的数据类型的顺序和名称与某一个函数类型完全一致，前者就是后者的一个实现。请大家回顾上面的示例并深刻理解这句话。

    我们可以声明一个函数类型的变量，如：

var splice func(string, string) string // 等价于 var splice MyFunc
    然后把函数myFunc赋给它：

splice = myFunc
    如此一来，我们就可以在这个变量之上实施调用动作了：

splice("1", "2")
    实际上，这是一个调用表达式。它由代表函数的标识符（这里是splice）以及代表调用动作的、由圆括号包裹的参数值列表组成。

    如果你觉得上面对splice变量声明和赋值有些啰嗦，那么可以这样来简化它：

var splice = func(part1 string, part2 string) string {
    return part1 + part2
}
    在这个示例中，我们直接使用了一个匿名函数来初始化splice变量。顾名思义，匿名函数就是不带名称的函数值。匿名函数直接由函数类型字面量和由花括号包裹的语句列表组成。注意，这里的函数类型字面量中的参数名称是不能被忽略的。

    其实，我们还可以进一步简化——索性省去splice变量。既然我们可以在代表函数的变量上实施调用表达式，那么在匿名函数上肯定也是可行的。因为它们的本质是相同的。后者的示例如下：

var result = func(part1 string, part2 string) string {
    return part1 + part2
}("1", "2")
    可以看到，在这个匿名函数之后的即是代表调用动作的参数值列表。注意，这里的result变量的类型不是函数类型，而与后面的匿名函数的结果类型是相同的。

  最后，函数类型的零值是nil。这意味着，一个未被显式赋值的、函数类型的变量的值必为nil。
 */

// 员工ID生成器
type EmployeeIdGenerator func(company string, department string, sn uint32) string

// 默认公司名称
var company = "Gophers"

// 序列号
var sn uint32

// 生成员工ID
func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	// 这是一条 if 语句，我们会在下一章讲解它。
	// 若员工ID生成器不可用，则无法生成员工ID，应直接返回。
	if generator == nil {
		return "", false
	}
	// 使用代码包 sync/atomic 中提供的原子操作函数可以保证并发安全。
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，所以提供这样一个函数作为辅助。
func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

func main() {
	var generator EmployeeIdGenerator
	generator = func(company string, department string, sn uint32) string {
		return appendSn(company + "-" + department + "-", sn)
	}
	fmt.Println(generateId(generator, "RD"))
}
