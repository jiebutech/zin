test:
	go test ./...
run:
	go run main.go
config:
	ln -s ./conf/application.$(word 2,$(MAKECMDGOALS)).yaml ./application.yaml
wire:
	cd wire && wire gen
build:
	go build -o ./build/ .