package main

import "fmt"

/**
Go语言-指针（续）
    我们在讲接口的时候说过，如果一个数据类型所拥有的方法集合中包含了某一个接口类型中的所有方法声明的实现，那么就可以说这个数据类型实现了那个接口类型。要获知一个数据类型都包含哪些方法并不难。但是要注意指针方法与值方法的区别。

    拥有指针方法Grow和Move的指针类型*Person是接口类型Animal的实现类型，但是它的基底类型Person却不是。这样的表象隐藏着另一条规则：一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。

    以上一小节练习题中的类型MyInt为例，如果Increase方法是它的指针方法且Decrease方法是它的值方法，那么*MyInt类型会拥有这两个方法，而MyInt类型仅拥有Decrease方法。再以Person类型为例。即使我们把Grow和Move都改为值方法，*Person类型也仍会是Animal接口的实现类型。另一方面，Grow和Move中只要有一个是指针方法，Person类型就不可能是Animal接口的实现类型。

    另外，还有一点需要大家注意，我们在基底类型的值上仍然可以调用它的指针方法。例如，若我们有一个Person类型的变量bp，则调用表达式bp.Grow()是合法的。这是因为，如果Go语言发现我们调用的Grow方法是bp的指针方法，那么它会把该调用表达式视为(&bp).Grow()。实际上，这时的bp.Grow()是(&bp).Grow()的速记法。
 */

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	name string
	age uint8
}

func (dog Dog) Name() string {
	return dog.name
}
func (dog Dog) Age() uint8 {
	return dog.age
}

func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
}
