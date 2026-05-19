package main
import (
"fmt"
"sync"
)
// Структура-одиночка
type Singleton struct {
value string
}
var (
instance *Singleton
once     sync.Once // Гарантирует однократное создание
)
// GetInstance возвращает единственный экземпляр Singleton
func GetInstance() *Singleton {
once.Do(func() {
instance = &Singleton{value: "Я единственный!"}
fmt.Println("Создан новый экземпляр Singleton")
})
return instance
}
func (s *Singleton) GetValue() string {
return s.value
}
func (s *Singleton) SetValue(v string) {
s.value = v
}
func main() {
// Попытка получить экземпляр несколько раз
s1 := GetInstance()
fmt.Println("s1:", s1.GetValue()) // Я единственный!
s2 := GetInstance()
fmt.Println("s2:", s2.GetValue()) // Я единственный!
// Проверим, что это один и тот же объект
s1.SetValue("Новое значение")
fmt.Println("s1 после изменения:", s1.GetValue()) // Новое значение
fmt.Println("s2 после изменения:", s2.GetValue()) // Новое значение (то же самое!)
// Проверим, что указатели одинаковые
fmt.Printf("s1 адрес: %p\n", s1) // Одинаковые адреса
fmt.Printf("s2 адрес: %p\n", s2) // Одинаковые адреса

fmt.Println("s1 == s2:", s1 == s2) // true
}