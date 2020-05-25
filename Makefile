.PHONY: front
front:
	cd ./front && npm run build

.PHONY: backend
backend:
	go build -v -o blocklator cmd/blocklator.go

.PHONY: docker
docker:
	docker build -f ./docker/Dockerfile -t tinywell/blocklator .	

.PHONY: start
start: front backend
	./blocklator server

.PHONY: clean
clean:
	rm -rf ./front/dist blocklator