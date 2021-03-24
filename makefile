all: windows linux macos
windows:
	GOOS=windows go build -o ./bin/win_network_gopher.exe ./main.go
linux:
	GOOS=linux go build -o ./bin/lin_network_gopher ./main.go
macos:
	GOOS=darwin go build -o ./bin/mac_network_gopher ./main.go
clean:
	rm -rf ./bin
