get:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega
	go get github.com/jmcvetta/napping
	go get qiniupkg.com/api.v7

test:
	ginkgo -r -v 

clean:
	go clean -i ./...
