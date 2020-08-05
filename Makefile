build:
	CGO_ENABLED=1 go build -o bass cmd/main.go 
clean:
	rm baas
