// Интерфейс продукта
type IGun interface {
Shoot()
}
// Конкретный продукт 1
type Ak47 struct{}
func (a *Ak47) Shoot() {
fmt.Println("AK47: Тра-та-та-та!")
}
// Конкретный продукт 2
type Musket struct{}
func (m *Musket) Shoot() {
fmt.Println("Мушкет: Ба-бах!")
}
// Интерфейс создателя (фабрики)
type IGunFactory interface {
CreateGun() IGun
}
// Конкретная фабрика 1
type Ak47Factory struct{}
func (f *Ak47Factory) CreateGun() IGun {
return &Ak47{}
}
// Конкретная фабрика 2
type MusketFactory struct{}
func (f *MusketFactory) CreateGun() IGun {
return &Musket{}
}
// Клиентский код
func main() {
// Создаем фабрику AK47
ak47Factory := &Ak47Factory{}
ak47 := ak47Factory.CreateGun()
ak47.Shoot() // Вывод: AK47: Тра-та-та-та!
// Создаем фабрику мушкета
musketFactory := &MusketFactory{}
musket := musketFactory.CreateGun()
musket.Shoot() // Вывод: Мушкет: Ба-бах!
}