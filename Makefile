get:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega
	go get github.com/jmcvetta/napping

test:
	cd testbiz
	go test

clean:
	cd testbiz;go clean -i ./...
