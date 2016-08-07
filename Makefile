include config.mk

build:
	go build
	GOOS=linux GOARCH=amd64 go build -o diskusageanalyzer.amd64

run: diskusageanalyzer
	./diskusageanalyzer "/Users/stefan/private/"

deploy:
	scp ./diskusageanalyzer.amd64 $(UPSTREAM):diskusageanalyzer

rmt_run:
	ssh $(UPSTREAM) "./diskusageanalyzer /var"
