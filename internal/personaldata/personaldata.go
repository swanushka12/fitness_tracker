package personaldata

import "fmt"

type Personal struct {
	Name string // имя пользователя
	Weight float64 // вес в кг
	Height float64 // рост в м
}

// вывод данных пользователя на экран
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}