package main

import "fmt"

// 示例1。
// AnimalCategory 代表动物分类学中的基本分类法。
type animalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac animalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

//示例2
type animal struct {
	scientificName   string  //学名
	animalCategory //动物基本分类
}

//示例3
type cat struct {
	name string
	animal
}

//该方法会“屏蔽”掉嵌入字段中的同名方法
func (c cat) String() string {
	return fmt.Sprintf("%s (category:%s,name:%q)",c.scientificName,
		c.animal.animalCategory,c.name)
}

func main() {
	//示例1
	category := animalCategory{species:"cat"}
	fmt.Printf("The animal category:%s\n",category)

	//示例2
	a := animal{
		scientificName:"American Shorthair",
		animalCategory:category,
	}
	fmt.Printf("The animal:%s\n",a)

	// 示例3。
	c := cat{
		name:   "little pig",
		animal: a,
	}
	//链式调用被屏蔽掉的同名方法
	fmt.Printf("The cat: %s\n", c.animal.String())
	fmt.Printf("The cat: %s\n", c)
	//Output:
	//The animal category:cat
	//The animal:cat
	//The cat: cat
	//The cat: American Shorthair (category:cat,name:"little pig")
}

