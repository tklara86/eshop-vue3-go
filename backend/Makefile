run:
	echo "Run triggered"
	echo "Golang"
	go run api/main.go

build:
	env GOOS=linux go build -o bin/api api/main.go
	env GOOS=linux go build -o bin/hello testLambda/main.go
deploy: build
	serverless deploy --aws-profile tomasz
