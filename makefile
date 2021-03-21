all: windows linux macos
windows:
	GOOS=windows go build -o ./bin/win_network_gopher.exe ./src
linux:
	GOOS=linux go build -o ./bin/lin_network_gopher ./src
macos:
	GOOS=darwin go build -o ./bin/mac_network_gopher ./src
clean:
	rm -rf ./bin
