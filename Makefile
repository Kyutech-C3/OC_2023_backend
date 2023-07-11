.PHONY run-app
run-app:
	docker compose up -d --build 


.PHONY go-test
go-test:
	go test ./...


.PHONY go-generate
go-generate:
	go generate ./...