https://github.com/golang-course/lesson13

go test
go test -v
go test -cover

go test -coverprofile=coverage.out
go tool cover -html=coverage.out - создает файл html в /tmp (FireFox не открывает)

go tool cover -html=coverage.out -o coverage.html - создает файл html в текущей папке

go test -bench=.
go test -bench=BenchmarkSimple
go test -bench=. -benchmem
go test -bench=. -benchtime=100x