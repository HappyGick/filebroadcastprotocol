client:
	go build -o bin/client/client ./cmd/client/main.go

server:
	go build -o bin/server/server ./cmd/server/main.go

run_client: client
	./bin/client/client

run_server: server
	./bin/server/server

clean:
	rm -rf bin