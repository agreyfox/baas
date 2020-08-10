build:
	CGO_ENABLED=1 go build -o baas cmd/main.go 
clean:
	rm baas
