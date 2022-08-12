BINARY_NAME= main.out

build:
	go build -o bin/tac main.go

run: go run main.go

docker-build:
	docker build --rm --tag $(BINARY_NAME) .

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go