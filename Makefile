compose:
	docker-compose up
build-image:
	docker build -t alexdias/store .
push-image:
	docker push alexdias/store
test:
	go test -v ./...
lint:
	golint ./...	
