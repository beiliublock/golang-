## 字典的操作和约束
1. 回顾：
   <br/>上次我们了解了，Go语言container包中自带的list和ring，明白了它们之间的同与不同；了解了list是如何
   平衡“延迟初始化”的，同时也对比了数组与链表。同时，它们也都是针对单一元素的容器。
2. 前导：
   <br/>我们今天要讲的字典(map)却不同，它存储的是键值对的集合。
   >什么是键值对？它是从英文key-value pair直译过来的一个词。顾名思义，一个键值对就代表了一对键和值。
   注意，一个“键”和一个“值”分别代表了一个从属于某一类型的独立值，把它们两个捆绑在一起就是一个键
   值对了。

   在Go语言规范中，应该是为了避免歧义，他们将键值对换了一种称呼，叫做：“键 - 元素对”。

### Q：为什么字典的键类型会受到约束？
Go语言的字典类型本质上是一个散列表（hash table），在这个实现中，键和元素的最大不同在于，键的类型是
受限的，而元素却可以是任意类型的。如果要探究限制的原因，我们就先要了解散列表中最重要的一个过程：映射。
我们可以把键理解为元素的一个索引，而我们可以通过这个索引查找与它成对的那个元素。

键和元素的这种对应关系，在数学里被称为“映射”，这也是“map”这个词的本意，哈希表的映射过程就存在于对
键 - 元素的增、删、改、查的操作之中。

    aMap := map[string]int {
        "one":1,
        "two":2,
        "three":3,
    }
    k := "two"
    if v,ok := aMap[k];ok {
        fmt.Printf("The element of key %q:%d\n",k,v)
    }else {
        fmt.Println("Not found")
    }
比如，我们要在哈希表中查找与某个键值对应的那个元素值，那么我们需要先把键值作为参数传给这个哈希表。
哈希表会先用哈希函数（hash function）把键值转换为哈希值。哈希值通常是一个无符号整数，一个哈希表会
持有一定数量的桶(bucket)，这些哈希桶会均匀的存储其哈希表收纳的键 - 元素对。

因此，哈希表会先用这个键哈希值的低几位定位到一个哈希桶，然后再去这个哈希桶中，查找这个键。由于，
键 - 元素对总是被捆绑在一起存储的，所以一旦找到了键，就一定能找到其对应的元素值，随后，哈希表就会
把相应的元素值作为结果返回。

只要这个键 - 元素对存在哈希表中就一定会被查找到，因为哈希表增、删、改操作时，其底层的实现与查的操作
如出一辙。在Go语言的字典中，每一个键值都是由它的哈希值代表的。也就是说，字典不会独立存储任何键的值，
但会独立存储他们的哈希值。

### Q：字典的键类型不能是哪些类型？
Go语言字典的键类型不可以是函数类型、字典类型和切片类型。

**问题解析一：**
 1. Go语言规范规定，在键值类型的值之间必须可以施加判等操作符（==和!=）。
 2. 如果键的类型是接口类型的，那么键值的实际类型也不能是上述三种类型，否则在程序运行过程中会引发panic。

**问题解析二：**
 + 为什么键类型必须支持判等操作？我们上面说过，Go语言一旦定位到了某一个哈希桶，那么就会试图在这个桶
 中查找键值。具体是怎么查找的呢？
 + 首先，每个哈希桶都会把自己包含的所有键的哈希存起来。Go语言会用被查找键的哈希与这些哈希值逐个比较，
 看看是否相等。如果一个相等的都没有，那么就说明这个桶中 没有要查找的键值，这时Go语言就会立刻返回接果
 + 如果有相等的，那么就再用键值本身去对比一次。为什么还要对比？因为不同的值的哈希值可能是相同的，我们
 把这种情况，叫做“散列冲突”
 + 所以，只有键的哈希值和键值都相等，才能说明查找到了匹配的键 - 元素对。

#### 知识扩展
**问题1：应该优先考虑哪些类型作为字典的键类型？**
 + 求哈希和判等操作的速度越快，对应的类型就越适合作为键类型
 + 不建议使用高级数据类型作为字典的键类型
 + 在基本类型中，优先选择数值类型和指针类型
 + 具体来说，需要结合实际情况，来适当的选择字典类型的键值类型

**问题2：在值为nil的字典上执行读操作会成功吗，那写操作呢？**
 + 由于字典是引用类型，所以当我们仅声明而不初始化一个字典类型的变量的时候，它的值会是nil
 + 这个问题虽然简单，但我们必须铭记于心，因为这涉及程序运行时的稳定性
 + 除了添加键 - 元素对，我们在一个值为nil的字典上做任何操作都不会引发错误。但是，当我们
 试图在一个值为nil的字典中添加键 - 元素对的时候，Go语言的运行时系统就会立即抛出一个panic,
 可以参考一下demo2.go的代码

**我们看一下sync包中的代码，发现其也有一个Map类型的结构体**

    这里同样贴一下代码：
    type Map struct {
    	mu Mutex
    	read atomic.Value // readOnly
    	misses int
    }
这里给一下官方文档的链接，感兴趣的可以了解一下：[sync.Map](http://docs.studygolang.com/pkg/sync/#Map)