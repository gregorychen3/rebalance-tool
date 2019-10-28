build:
	GO111MODULE=on go build -o rebalance_tool

test:
	go test -count=1 -v ./...

