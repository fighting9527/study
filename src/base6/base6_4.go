package main

import "fmt"

/**
Go语言-指针
    我们在前面多次提到过指针及指针类型。例如，*Person是Person的指针类型。又例如，表达式&p的求值结果是p的指针。方法的接收者类型的不同会给方法的功能带来什么影响？该方法所属的类型又会因此发生哪些潜移默化的改变？现在，我们就来解答第一个问题。至于第二个问题，我会在下一小节予以解答。

    指针操作涉及到两个操作符——&和*。这两个操作符均有多个用途。但是当它们作为地址操作符出现时，前者的作用是取址，而后者的作用是取值。更通俗地讲，当地址操作符&被应用到一个值上时会取出指向该值的指针值，而当地址操作符*被应用到一个指针值上时会取出该指针指向的那个值。它们可以被视为相反的操作。

    除此之外，当*出现在一个类型之前（如*Person和*[3]string）时就不能被看做是操作符了，而应该被视为一个符号。如此组合而成的标识符所表达的含义是作为第二部分的那个类型的指针类型。我们也可以把其中的第二部分所代表的类型称为基底类型。例如，*[3]string是数组类型[3]string的指针类型，而[3]string是*[3]string的基底类型。

    好了，我们现在回过头去再看结构体类型Person。它及其两个方法的完整声明如下：

type Person struct {
    Name    string
    Gender  string
    Age     uint8
    Address string
}

func (person *Person) Grow() {
    person.Age++
}

func (person *Person) Move(newAddress string) string {
    old := person.Address
    person.Address = newAddress
    return old
}
    注意，Person的两个方法Grow和Move的接收者类型都是*Person，而不是Person。只要一个方法的接收者类型是其所属类型的指针类型而不是该类型本身，那么我就可以称该方法为一个指针方法。上面的Grow方法和Move方法都是Person类型的指针方法。

    相对的，如果一个方法的接收者类型就是其所属的类型本身，那么我们就可以把它叫做值方法。我们只要微调一下Grow方法的接收者类型就可以把它从指针方法变为值方法：

func (person Person) Grow() {
    person.Age++
}
    那指针方法和值方法到底有什么区别呢？我们在保留上述修改的前提下编写如下代码：

p := Person{"Robert", "Male", 33, "Beijing"}
p.Grow()
fmt.Printf("%v\n", p)
    这段代码被执行后，标准输出会打印出什么内容呢？直觉上，34会被打印出来，但是被打印出来的却是33。这是怎么回事呢？Grow方法的功能失效了？！

    解答这个问题需要引出一条定论：方法的接收者标识符所代表的是该方法当前所属的那个值的一个副本，而不是该值本身。例如，在上述代码中，Person类型的Grow方法的接收者标识符person代表的是p的值的一个拷贝，而不是p的值。我们在调用Grow方法的时候，Go语言会将p的值复制一份并将其作为此次调用的当前值。正因为如此，Grow方法中的person.Age++语句的执行会使这个副本的Age字段的值变为34，而p的Age字段的值却依然是33。这就是问题所在。

    只要我们把Grow变回指针方法就可以解决这个问题。原因是，这时的person代表的是p的值的指针的副本。指针的副本仍会指向p的值。另外，之所以选择表达式person.Age成立，是因为如果Go语言发现person是指针并且指向的那个值有Age字段，那么就会把该表达式视为(*person).Age。其实，这时的person.Age正是(*person).Age的速记法。
 */

type MyInt struct {
	n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
}

func main() {
	mi := MyInt{}
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)
}
