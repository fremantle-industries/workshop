default: test build.compose up

configure:
	./tools/configure

test:
	go test ./...

build.compose:
	docker compose build

build.docker:
	docker build . -t fremantle-industries/workshop

build.bin:
	go build -o ./bin ./main.go

publish.docker:
	docker push fremantle-industries/workshop

k8s.apply:
	kubectl apply -f deploy/k8s

up:
	docker compose up

down:
	docker compose down
