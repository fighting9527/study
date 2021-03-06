package main

import "fmt"

/**
Go语言的整数类型一共有10个。

    其中计算架构相关的整数类型有两个，即：有符号的整数类型int和无符号的整数类型uint。

    顺便提一下，有符号的整数类型会使用最高位的比特（bit）表示整数的正负。显然，这会对它能表示的整数的范围有一定的损耗（使其缩小）。而无符号的整数类型会使用所有的比特位来表示数值。如此类型的值均为正数。这也是用“无符号的”来形容它们的原因。

    言归正传，为什么说这两个整数类型是计算架构相关的呢？这是因为，在不同的计算架构的计算机之上，它们体现的宽度是不同的。宽度即指存储一个某类型的值所需要的空间。空间的单位可以是比特，也可以是字节（byte）。请看下表。



    我想你应该已经能够悟到它们的对应关系了。

    除了这两个计算架构相关的整数类型之外，还有8个可以显式表达自身宽度的整数类型。如下表所示。


 */

func main() {
	// 变量声明和赋值语句，由关键字var、变量名num、变量类型uint64、特殊标记=，以及值10组成。
	var num uint64 = 65535

	// 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
	size := (8)

	// 打印函数调用语句。在这里用于描述一个uint64类型的变量所需占用的比特数。
	// 这里用到了字符串的格式化函数。
	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为 %d 个字节。\n", num, size)
}
