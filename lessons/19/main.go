package main

/*
POST /api/v1/users
Host: example.com
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Accept: application/json

	{
	  "name": "Alice",
	  "email": "alice@example.com"
	}

HTTP/1.1 201 Created
Content-Type: application/json
Location: /api/v1/users/42

	{
	  "id": 42,
	  "name": "Alice",
	  "email": "alice@example.com",
	  "created_at": "2024-01-15T10:30:00Z"
	}
*/
func main() {

}
