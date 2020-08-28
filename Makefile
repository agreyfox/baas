build:
	CGO_ENABLED=1 go build -o baas cmd/main.go 
sync:
	scp baas moactools@jeff:/home/moactools/
clean:
	rm baas
