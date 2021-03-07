all: windows linux macos
windows:
    GOOS=windows go build -o ./bin/win_network_gopher ./src/main.go
linux:
    GOOS=linux  go build -o ./bin/linux_network_gopher ./src/main.go
macos:
    GOOS=darwin go build -o ./bin/macos_network_gopher ./src/main.go