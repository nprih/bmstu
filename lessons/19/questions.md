Как отобразить хозяина заказа при выводе всех заказов? (Заказы являются вложенной структурой пользователя, 
точнее слайсом структур заказов)

Я привык в PHP (и не только) иметь доступ к свойствам и методам 
(если они, конечно, объявлены, как доступные откуда угодно) вложенных объектов.

Например:
Object_user имеет поля/свойства: id, name, mail, age, balance, orders (которое, в свою очередь является object_orders)
Object_order имеет поля: id, length, width, name, status

Чтобы заполучить свойство length (объекта Object_order) имея в распоряжении только Object_user,
мне достаточно написать следующую конструкцию (на PHP, упрощенно, конечно): Object_user->Object_order->length
По такому же принципу доступны и методы вложенных объектов (Object_user->Object_order->setLength($someLength)

А возможен ли подобный подход в Go?
ЗЫ: В JS, если не путаю, подобный принцип тоже присутствует, там только вместо "->" используется ".", как и в Go, 
как и во всех C-подобных языках.

type Order struct {
id     int
length int
width  int
name   string
status string
}
type User struct {
id      int
name    string
mail    string
age     int
balance int
orders  []Order
}
func main() {
newUser := User{
id:      1,
name:    "User",
mail:    "user@user.ee",
age:     12,
balance: 13,
orders: []Order{
{
id:     2,
name:   "Phone",
status: "In base",
},
},
}
fmt.Println(newUser.orders[0].name)
}