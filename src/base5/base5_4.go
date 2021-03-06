package main

import "fmt"

/**
Go语言-字典类型
    Go语言的字典（Map）类型其实是哈希表（Hash Table）的一个实现。字典用于存储键-元素对（更通俗的说法是键-值对）的无序集合。注意，同一个字典中的每个键都是唯一的。如果我们在向字典中放入一个键值对的时候其中已经有相同的键的话，那么与此键关联的那个值会被新值替换。

    字典类型的字面量如下：

map[K]T
    其中，“K”意为键的类型，而“T”则代表元素（或称值）的类型。如果我们要描述一个键类型为int、值类型为string的字典类型的话，应该这样写：

map[int]string
    请注意，字典的键类型必须是可比较的，否则会引起错误。也就是说，它不能是切片、字典或函数类型。

    字典值的字面量表示法实际上与数组和切片的字面量表示法很相似。首先，最左边仍然是类型字面量，右边紧挨着由花括号包裹且有英文逗号分隔的键值对。每个键值对的键和值之间由英文冒号分隔。以字典类型map[int]string为例，它的值的字面量可以是这样的：

map[int]string{1: "a", 2: "b", 3: "c"}
    我们可以把这个值赋给一个变量：

mm := map[int]string{1: "a", 2: "b", 3: "c"}
    然后运用索引表达式取出字典中的值，就像这样：

b := mm[2]
    注意，在这里，我们放入方括号中的不再是索引值（实际上，字典中的键值对也没有索引），而是与我们要取出的值对应的那个键。在上例中变量b的值必是字符串"b"。当然，也可以利用索引表达式来赋值，比如这样：

mm[2] = b + "2"
    这使得字典mm中与键2对应的值变为了"b2"。现在我们再来向mm添加一个键值对：

mm[4] = ""
    之后，在从中取出与`4`和`5`对应的值：

d := mm[4]
e := mm[5]
    此时，变量d和e的值都会是多少呢？答案是都为""，即空字符串。对于变量d来说，由于在字典mm中与4对应的值就是""，所以索引表达式mm[4]的求值结果必为""。这理所应当。但是mm[5]的求值结果为什么也是空字符串呢？原因是，在Go语言中有这样一项规定，即：对于字典值来说，如果其中不存在索引表达式欲取出的键值对，那么就以它的值类型的空值（或称默认值）作为该索引表达式的求值结果。由于字符串类型的空值为""，所以mm[5]的求值结果即为""。

    在不知道mm的确切值的情况下，我们无法得知mm[5]的求值结果意味着什么？它意味着5对应的值就是一个空字符串？还是说mm中根本就没有键为5的键值对？这无所判别。为了解决这个问题，Go语言为我们提供了另外一个写法，即：

e, ok := mm[5]
    针对字典的索引表达式可以有两个求值结果。第二个求值结果是bool类型的。它用于表明字典值中是否存在指定的键值对。在上例中，变量ok必为false。因为mm中不存在以5为键的键值对。

    从字典中删除键值对的方法非常简单，仅仅是调用内建函数delete而已，就像这样：

delete(mm, 4)
    无论mm中是否存在以4为键的键值对，delete都会“无声”地执行完毕。我们用“有则删除，无则不做”可以很好地概括它的行为。

    最后，与切片类型相同，字典类型属于引用类型。它的零值即为nil。
 */
func main() {
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
	mm2["scala"] = 25
	mm2["erlang"] = 50
	mm2["python"] = 0
	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])
	e, ok := mm2["golang"]
	fmt.Println(e, ok)
}
