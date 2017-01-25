all: grpcrgen test

grpcrgen: myservice
	grpcrgen

myservice:
	flatbuffers/flatc --go --grpc -I fbs fbs/*

test:
	go test

clean:
	rm -rf myservice
	rm -rf grpcrgen
