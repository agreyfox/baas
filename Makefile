build:
	CGO_ENABLED=1 go build -o baas cmd/main.go 
sync:
	scp baas moactools@jeff:/home/moactools/
sq:
	scp baas root@47.116.97.219:/root/baas/new
sy:
	scp -i ~/block/doc/moacfoundation_svr3 baas root@47.92.242.197:/home/baas/new
clean:
	rm baas
