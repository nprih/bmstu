github.com/gtank/cryptopasta

openssl genrsa -out rootCA.key 2048
openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024 -out rootCA.pem

go get github.com/golang-jwt/jwt/v5

curl http://localhost:8085/users
curl http://localhost:8085/secret
curl -X POST http://localhost:8085/login -H "Content-Type: application/json" -d '{"name":"Sasha", "password":"654321"}'

curl http://localhost:8085/secret -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkIjoxNzgxMTk0NjU4LCJleHBpcmVkIjoxNzgxMjgxMDU4LCJuYW1lIjoiU2FzaGEifQ.bskLdk_CBkI85EE_Wk7qfDzvaWQ_z2bHAi9yEzWFTEw"