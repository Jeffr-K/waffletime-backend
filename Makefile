hello:
	echo "Main Function is started.."

build:
	go build -o cmd/main main.go

run:
	go run cmd/main.go