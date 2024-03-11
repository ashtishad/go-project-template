run:
	export API_PORT=8000 \
	export DB_USER=ash \
	export DB_PASSWD=strong_password \
	export DB_HOST=127.0.0.1 \
	export DB_PORT=5432 \
	export DB_NAME=datalake \
&& go run main.go
test:
	go test -v ./...
race:
	go run -race .
lint:
	golangci-lint run
