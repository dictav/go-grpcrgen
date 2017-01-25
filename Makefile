grpcrgen:
	go generate
	go build -o grpcrgen

clean:
	rm -rf test_data
	rm grpcrgen

test: test_data
	go test

test_data:
	flatc --go --grpc test_data.fbs
