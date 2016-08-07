build:
	go build
	GOOS=linux GOARCH=amd64 go build -o diskusageanalyzer.amd64

run: diskusageanalyzer
	./diskusageanalyzer "/Users/stefan/private/"

deploy:
	scp ./diskusageanalyzer.amd64 nclutz-docker001.ops.local.netconomy.net:diskusageanalyzer

rmt_run:
	ssh nclutz-docker001.ops.local.netconomy.net "./diskusageanalyzer /var"
