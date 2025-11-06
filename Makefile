lint:
	golangci-lint run ./...

fmtimport:
	goimports -w .

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

run:
	go run ./cmd/elec-trade-data-tw/main.go

build:
	go build -o bin/elec-trade-data-tw ./cmd/elec-trade-data-tw/main.go

runb:
	./bin/elec-trade-data-tw

output:
	GOOS=linux GOARCH=amd64 go build -o ./bin/elec-trade-data-tw-linux-amd64 ./cmd/elec-trade-data-tw
