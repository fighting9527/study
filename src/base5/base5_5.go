package main

import "fmt"

/**
Go语言-通道类型
    通道（Channel）是Go语言中一种非常独特的数据结构。它可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的。相比之下，我们之前介绍的那些数据类型都不是并发安全的。这一点需要特别注意。

    Goroutine（也称为Go程序）可以被看做是承载可被并发执行的代码块的载体。它们由Go语言的运行时系统调度，并依托操作系统线程（又称内核线程）来并发地执行其中的代码块。至于怎样编写这样的代码块以及怎样驱动这样的代码块执行，我们先按下不表。

    通道类型的表示方法很简单，仅由两部分组成，如下：

chan T
    在这个类型字面量中，左边是代表通道类型的关键字chan，而右边则是一个可变的部分，即代表该通道类型允许传递的数据的类型（或称通道的元素类型）。这两部分之间需要以空格分隔。

    与其它的数据类型不同，我们无法表示一个通道类型的值。因此，我们也无法用字面量来为通道类型的变量赋值。我们只能通过调用内建函数make来达到目的。make函数可接受两个参数。第一个参数是代表了将被初始化的值的类型的字面量（比如chan int），而第二个参数则是值的长度。例如，若我们想要初始化一个长度为5且元素类型为int的通道值，则需要这样写：

make(chan int, 5)
    顺便说一句，实际上make函数也可以被用来初始化切片类型或字典类型的值。


    确切地说，通道值的长度应该被称为其缓存的尺寸。换句话说，它代表着通道值中可以暂存的数据的个数。注意，暂存在通道值中的数据是先进先出的，即：越早被放入（或称发送）到通道值的数据会越先被取出（或称接收）。

    下面，我们声明一个通道类型的变量，并为其赋值：

ch1 := make(chan string, 5)
    这样一来，我们就可以使用接收操作符<-向通道值发送数据了。当然，也可以使用它从通道值接收数据。例如，如果我们要向通道ch1发送字符串"value1"，那么应该这样做：

ch1 <- "value1"
    另一方面，我们若想从ch1那里接收字符串，则要这样：

<- ch1
    这时，我们可以直接把接收到的字符串赋给一个变量，如：

value := <- ch1
    与针对字典值的索引表达式一样，针对通道值的接收操作也可以有第二个结果值。请看下面的示例：

value, ok := <- ch1

    这样做的目的同样是为了消除与零值有关的歧义。这里的变量ok的值同样是bool类型的。它代表了通道值的状态，true代表通道值有效，而false则代表通道值已无效（或称已关闭）。更深层次的原因是，如果在接收操作进行之前或过程中通道值被关闭了，则接收操作会立即结束并返回一个该通道值的元素类型的零值。按照上面的第一种写法，我们无从判断接收到零值的原因是什么。不过，有了第二个结果值之后，这种判断就好做了。

    说到关闭通道值，我们可以通过调用内建函数close来达到目的，就像这样：

close(ch1)
    请注意，对通道值的重复关闭会引发运行时恐慌。这会使程序崩溃。所以一定要避免这种情况的发生。另外，在通道值有效的前提下，针对它的发送操作会在通道值已满（其中缓存的数据的个数已等于它的长度）时被阻塞。而向一个已被关闭的通道值发送数据会引发运行时恐慌。另一方面，针对有效通道值的接收操作会在它已空（其中没有缓存任何数据）时被阻塞。除此之外，还有几条与通道的发送和接收操作有关的规则。不过在这里我们记住上面这三条就可以了。

  最后，与切片和字典类型相同，通道类型属于引用类型。它的零值即为nil。Go语言-通道类型
    通道（Channel）是Go语言中一种非常独特的数据结构。它可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的。相比之下，我们之前介绍的那些数据类型都不是并发安全的。这一点需要特别注意。

    Goroutine（也称为Go程序）可以被看做是承载可被并发执行的代码块的载体。它们由Go语言的运行时系统调度，并依托操作系统线程（又称内核线程）来并发地执行其中的代码块。至于怎样编写这样的代码块以及怎样驱动这样的代码块执行，我们先按下不表。

    通道类型的表示方法很简单，仅由两部分组成，如下：

chan T
    在这个类型字面量中，左边是代表通道类型的关键字chan，而右边则是一个可变的部分，即代表该通道类型允许传递的数据的类型（或称通道的元素类型）。这两部分之间需要以空格分隔。

    与其它的数据类型不同，我们无法表示一个通道类型的值。因此，我们也无法用字面量来为通道类型的变量赋值。我们只能通过调用内建函数make来达到目的。make函数可接受两个参数。第一个参数是代表了将被初始化的值的类型的字面量（比如chan int），而第二个参数则是值的长度。例如，若我们想要初始化一个长度为5且元素类型为int的通道值，则需要这样写：

make(chan int, 5)
    顺便说一句，实际上make函数也可以被用来初始化切片类型或字典类型的值。


    确切地说，通道值的长度应该被称为其缓存的尺寸。换句话说，它代表着通道值中可以暂存的数据的个数。注意，暂存在通道值中的数据是先进先出的，即：越早被放入（或称发送）到通道值的数据会越先被取出（或称接收）。

    下面，我们声明一个通道类型的变量，并为其赋值：

ch1 := make(chan string, 5)
    这样一来，我们就可以使用接收操作符<-向通道值发送数据了。当然，也可以使用它从通道值接收数据。例如，如果我们要向通道ch1发送字符串"value1"，那么应该这样做：

ch1 <- "value1"
    另一方面，我们若想从ch1那里接收字符串，则要这样：

<- ch1
    这时，我们可以直接把接收到的字符串赋给一个变量，如：

value := <- ch1
    与针对字典值的索引表达式一样，针对通道值的接收操作也可以有第二个结果值。请看下面的示例：

value, ok := <- ch1

    这样做的目的同样是为了消除与零值有关的歧义。这里的变量ok的值同样是bool类型的。它代表了通道值的状态，true代表通道值有效，而false则代表通道值已无效（或称已关闭）。更深层次的原因是，如果在接收操作进行之前或过程中通道值被关闭了，则接收操作会立即结束并返回一个该通道值的元素类型的零值。按照上面的第一种写法，我们无从判断接收到零值的原因是什么。不过，有了第二个结果值之后，这种判断就好做了。

    说到关闭通道值，我们可以通过调用内建函数close来达到目的，就像这样：

close(ch1)
    请注意，对通道值的重复关闭会引发运行时恐慌。这会使程序崩溃。所以一定要避免这种情况的发生。另外，在通道值有效的前提下，针对它的发送操作会在通道值已满（其中缓存的数据的个数已等于它的长度）时被阻塞。而向一个已被关闭的通道值发送数据会引发运行时恐慌。另一方面，针对有效通道值的接收操作会在它已空（其中没有缓存任何数据）时被阻塞。除此之外，还有几条与通道的发送和接收操作有关的规则。不过在这里我们记住上面这三条就可以了。

  最后，与切片和字典类型相同，通道类型属于引用类型。它的零值即为nil。
 */

func main() {
	ch2 := make(chan string, 1)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- ("已到达！")
	}()
	var value string = "数据"
	value = value + (<- ch2)
	fmt.Println(value)
}
