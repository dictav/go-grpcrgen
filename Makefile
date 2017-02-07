grpcrgen:
	go generate
	go build -o grpcrgen

clean:
	rm -rf test_data
	rm grpcrgen

test: test_data
	go test -v -race
	go test -v -race ./cmd/...
	go test -v -race ./example


test_data:
	flatc --go --grpc test_data.fbs
