all: test
build: clean get
	go build
test: build todo metalinter
	GOOGLE_APPLICATION_CREDENTIALS=~/.ssh/sr.json  go test -v	
get:
	go get -v

metalinter:
	gometalinter .
       
clean:
	go clean  

todo:
	@echo Implement shoutdown script
	@echo - https://cloud.google.com/compute/docs/shutdownscript

