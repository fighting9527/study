package main

import (
	"fmt"
	"time"
)

/**
Go语言-通道的更多种类
    我们在上一节所说的通道，实际上只是Go语言中的通道的一种。它被称为带缓冲的通道，或简称为缓冲通道。

    通道有带缓冲和非缓冲之分。我们已经说过，缓冲通道中可以缓存N个数据。我们在初始化一个通道值的时候必须指定这个N。相对的，非缓冲通道不会缓存任何数据。发送方在向通道值发送数据的时候会立即被阻塞，直到有某一个接收方已从该通道值中接收了这条数据。非缓冲的通道值的初始化方法如下：

make(chan int, 0)
    注意，在这里，给予make函数的第二个参数值是0。

    除了上述分类方法，我们还可以以数据在通道中的传输方向为依据来划分通道。默认情况下，通道都是双向的，即双向通道。如果数据只能在通道中单向传输，那么该通道就被称作单向通道。我们在初始化一个通道值的时候不能指定它为单向。但是，在编写类型声明的时候，我们却是可以这样做的。例如：

type Receiver <-chan int
    类型Receiver代表了一个只可从中接收数据的单向通道类型。这样的通道也被称为接收通道。在关键字chan左边的接收操作符<-形象地表示出了数据的流向。相对应的，如果我们想声明一个发送通道类型，那么应该这样：

type Sender chan<- int
    这次<-被放在了chan的右边，并且“箭头”直指“通道”。想必不用多说你也能明白了。我们可以把一个双向通道值赋予上述类型的变量，就像这样：

var myChannel = make(chan int, 3)
var sender Sender = myChannel
var receiver Receiver = myChannel
    但是，反之则是不行的。像下面这样的代码是通不过编译的：

var myChannel1 chan int = sender
    单向通道的主要作用是约束程序对通道值的使用方式。比如，我们调用一个函数时给予它一个发送通道作为参数，以此来约束它只能向该通道发送数据。又比如，一个函数将一个接收通道作为结果返回，以此来约束调用该函数的代码只能从这个通道中接收数据。这属于API设计的范畴。因此我们在这里仅了解一下即可。
 */

type Sender chan <- int
type Receiver <- chan int

func main() {
	var myChannel = make(chan int, 0)
	var number int = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <- receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}
