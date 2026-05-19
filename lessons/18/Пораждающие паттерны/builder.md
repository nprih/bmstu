type House struct {
windowType string
doorType   string
floorCount int
hasGarage  bool
hasGarden  bool
}
// Builder - интерфейс строителя
type HouseBuilder interface {
setWindowType()
setDoorType()
setFloorCount()
setGarage()
setGarden()
build() House
}
// ConcreteBuilder - реализация строителя
type ModernHouseBuilder struct {
house House
}
func NewModernHouseBuilder() *ModernHouseBuilder {
return &ModernHouseBuilder{house: House{}}
}
func (b *ModernHouseBuilder) setWindowType() {
b.house.windowType = "Панорамные окна"
}
func (b *ModernHouseBuilder) setDoorType() {
b.house.doorType = "Стеклянная дверь"
}
func (b *ModernHouseBuilder) setFloorCount() {
b.house.floorCount = 3
}
func (b *ModernHouseBuilder) setGarage() {
b.house.hasGarage = true
}
func (b *ModernHouseBuilder) setGarden() {
b.house.hasGarden = false
}
func (b *ModernHouseBuilder) build() House {
return b.house
}
// Другой конкретный строитель
type CountryHouseBuilder struct {
house House
}
func NewCountryHouseBuilder() *CountryHouseBuilder {
return &CountryHouseBuilder{house: House{}}
}

func (b *CountryHouseBuilder) setWindowType() {
b.house.windowType = "Деревянные окна"
}
func (b *CountryHouseBuilder) setDoorType() {
b.house.doorType = "Деревянная дверь"
}
func (b *CountryHouseBuilder) setFloorCount() {
b.house.floorCount = 1
}
func (b *CountryHouseBuilder) setGarage() {
b.house.hasGarage = false
}
func (b *CountryHouseBuilder) setGarden() {
b.house.hasGarden = true
}
func (b *CountryHouseBuilder) build() House {
return b.house
}
// Client code
func main() {

// Строим современный дом
modernBuilder := NewModernHouseBuilder()
modernBuilder.setWindowType()
modernBuilder.setDoorType()
modernBuilder.setFloorCount()
modernBuilder.setGarage()
modernHouse := modernBuilder.build()

fmt.Println("Современный дом:")
fmt.Printf("Окна: %s\n", modernHouse.windowType)
fmt.Printf("Дверь: %s\n", modernHouse.doorType)
fmt.Printf("Этажи: %d\n", modernHouse.floorCount)
fmt.Printf("Гараж: %t\n", modernHouse.hasGarage)
fmt.Printf("Сад: %t\n\n", modernHouse.hasGarden)
    
// Строим загородный дом
countryBuilder := NewCountryHouseBuilder()
countryBuilder.setWindowType()
countryBuilder.setDoorType()
countryBuilder.setFloorCount()
countryBuilder.setGarden()
countryHouse := countryBuilder.build()
    
fmt.Println("Загородный дом:")
fmt.Printf("Окна: %s\n", countryHouse.windowType)
fmt.Printf("Дверь: %s\n", countryHouse.doorType)
fmt.Printf("Этажи: %d\n", countryHouse.floorCount)
fmt.Printf("Гараж: %t\n", countryHouse.hasGarage)
fmt.Printf("Сад: %t\n", countryHouse.hasGarden)
}