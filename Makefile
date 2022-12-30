gen:
	kratos proto client src/proto/*.proto

s:
	go build -o server src/server/server.go

c:
	go build -o client src/client/client.go

winc:
	 CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o client.exe src/client/client.go

wins:
	 CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o server.exe src/client/server.go

clean:
	rm server server.exe client client.exe
