buildall:
	make clean
	go build -o bin/ .
	GOOS=windows GOARCH=amd64 go build -o bin/ .

clean:
	rm -rf bin/