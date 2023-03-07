# env GOOS=target-OS GOARCH=target-architecture go build main.go names.go lastnames.go 
go build -o target/datagen.linux.x86_64 main.go names.go lastnames.go 
env GOOS=windows GOARCH=amd64 go build -o target/datagen.exe main.go names.go lastnames.go
env GOOS=linux GOARCH=386 go build -o target/datagen.linux.386 main.go names.go lastnames.go 
env GOOS=darwin GOARCH=arm64 go build -o target/datagem.darwin.arm64 main.go names.go lastnames.go