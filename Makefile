
all: deps gils.bin

deps:
	go get github.com/boutros/marc;
	cd js; yarn install; cd ..

gils.bin: main.go marc_ops.go
	go build -o gils.bin main.go marc_ops.go