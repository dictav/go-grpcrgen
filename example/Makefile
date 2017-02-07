all: myrouter test

myrouter: myservice
	grpcrgen -o myrouter myservice

myservice:
	flatc --go --grpc -I fbs fbs/*

test:
	go test

clean:
	rm -rf myservice
	rm -rf myrouter
