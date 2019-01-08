package main

import "fmt"

type dog struct {
	name           string //名字
	scientificName string //学名
	category       string //动物学基本分类
}

func myNew(name, scientificName, category string) dog {
	return dog{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (d *dog) setName(name string) {
	d.name = name
}

func (d dog) setNameOfCopy(name string) {
	d.name = name
}
//注意方法名不能和字段名重名
func (d dog) Name() string {
	return d.name
}

func (d dog)ScientificName() string{
	return d.scientificName
}

func (d dog)Category() string{
	return d.category
}

func (d dog) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		d.scientificName, d.category, d.name)
}

func main() {
	d := myNew("little pig","Golden retriever","dog")
	d.setName("monster") //（&d).setName("monster")
	fmt.Printf("The dog:%s\n",d)

	d.setNameOfCopy("little pig")
	fmt.Printf("The dog:%s\n",d)

	type pet interface {
		setName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_,ok := interface{}(d).(pet)
	fmt.Printf("dog implements interface Pet: %v\n", ok)
	_,ok = interface{}(&d).(pet)
	fmt.Printf("*dog implements interface Pet: %v\n", ok)
	//Output:
	//The dog:Golden retriever (category: dog, name: "monster")
	//The dog:Golden retriever (category: dog, name: "monster")
	//dog implements interface Pet: false
	//*dog implements interface Pet: true
}


