package personaldata

import "fmt"

type Personal struct {
	// TODO: добавить поля
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Printf("Имя:%s\nВес:%.2f\nРост:%.2f\n", p.Name, p.Weight, p.Height)
}
