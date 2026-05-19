package main

/*
1) GET - получить ресурс
2) HEAD - то же самое, что и get, но без тела - только заголовок
3) POST - создать новую запись
4) PUT - заменить запись на новую целиком
5) PATCH/UPDATE - частичная замена
6) DELETE - удалить запись
7) OPTION - узнать допустимые методы
*/

/*
[
	{1, Alise, 22}
	{2, James, 23}
	{3, Arnold, 21}
]
	REST API, значит чтобы достать Alice
	GET /users/1 -> {1, Alise, 22}
	PUT    /users/1 -> request body: {"id": 1,
									  "name": Alise Smith,
									  "age": 23}
	PATCH  /users/1 -> request body: {"age":23}
*/

func main() {

}
