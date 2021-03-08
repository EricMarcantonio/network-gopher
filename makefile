all: windows linux macos
clean:		./src
	rm -rf ./bin
windows:	./src
    GOOS=windows go build -o ./bin/win_network_gopher ./src
linux:		./src
    GOOS=linux go build -o ./bin/linux_network_gopher ./src
macos:		./src
    GOOS=darwin go build -o ./bin/macos_network_gopher ./src