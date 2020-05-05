compose:
	docker-compose up
build-image:
	sudo rm -rf postgres
	docker build -t alexdias/store .
push-image:
	docker push alexdias/store
test:
	go test -v ./...
lint:
	golint ./...
