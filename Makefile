
all: deps jsx gils.bin

deps: js/package.json
	go get github.com/boutros/marc;
	cd js; yarn install; cd ..

jsx:
	cd js; webpack; cd ..

gils.bin: main.go marc_ops.go
	go build -o gils.bin main.go marc_ops.go