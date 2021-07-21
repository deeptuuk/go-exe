package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("name = ", this.Name)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	//当前this 是Hero的一个拷贝
	this.Name = newName
}

func main() {

	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show()

	hero.SetName("test")
	hero.Show()
}
