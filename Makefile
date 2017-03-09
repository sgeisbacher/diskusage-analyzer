build:
	GOOS=darwin GOARCH=amd64 go build -o bin/diskusage-analyzer.darwin64
	GOOS=linux GOARCH=amd64 go build -o bin/diskusage-analyzer.linux64

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

test:
	go get -t -v ./...
	go test -v -race ./...
