Сущности: user, order
User {
id int
name string
email string
age int
balance int
order []order
}
Order {
id int
userId int
name string
status string //[на складе, в ПВЗ]
}
API документация
GET /api/v1/users - get all users
GET /api/v1/orders - get all orders
POST /api/v1/users - create new user
POST /api/v1/orders - create new order
GET /api/v1/users/{id}/orders - get all orders for currnet user
PATCH /api/v1/orders/{id} {status: "на складе/в ПВЗ"} //no auth
* Обработка ошибок по rest правилам со статус-кодами
** Получить всех пользователей, у которых
в ПВЗ есть товары