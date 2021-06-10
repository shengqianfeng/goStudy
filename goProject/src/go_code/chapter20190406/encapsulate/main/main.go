package main
import (
	"fmt"
	"go_code/chapter20190406/encapsulate/model"
)

//封装快速入门
func main()  {
	
	p:= model.NewPerson("smith")

	fmt.Println(p)//&{smith 0 0} 返回一个p实例的指针

	p.SetSal(5000)
	p.SetAge(18)

	fmt.Println(p)//&{smith 18 5000}
}