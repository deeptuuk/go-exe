package main

import "fmt"

//如果类名首字母大写，则其他包也能使用这个类
type Hero struct {
	//类的属性首字母大写，则对外可以访问
	Name  string
	Ad    int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("name = ", this.Name)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//当前this 是当前对象
	this.Name = newName
}

func main() {

	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}
	hero.Show()

	hero.SetName("test")
	hero.Show()
}
