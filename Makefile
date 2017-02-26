build:
	go build
	GOOS=linux GOARCH=amd64 go build -o diskusageanalyzer.amd64

deploy:
	scp ./diskusageanalyzer.amd64 $(UPSTREAM):diskusageanalyzer

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

