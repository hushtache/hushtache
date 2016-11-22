#!/bin/bash
export GOPATH=$(pwd)
go build -o release/hushtache main.go

# build linux
env GOOS=linux GOARCH=amd64 go build -o release/linux64/hushtache main.go
cp README.md release/linux64/README.md
cp LICENSE.md release/linux64/LICENSE.md

# build linux
env GOOS=linux GOARCH=386 go build -o release/linux32/hushtache main.go
cp README.md release/linux32/README.md
cp LICENSE.md release/linux32/LICENSE.md

# build mac
env GOOS=darwin GOARCH=amd64 go build -o release/darwin64/hushtache main.go
cp README.md release/darwin64/README.md
cp LICENSE.md release/darwin64/LICENSE.md

# build mac
env GOOS=darwin GOARCH=386 go build -o release/darwin32/hushtache main.go
cp README.md release/darwin32/README.md
cp LICENSE.md release/darwin32/LICENSE.md

# build windows
env GOOS=windows GOARCH=amd64 go build -o release/windows64/hushtache.exe main.go
cp README.md release/windows64/README.md
cp LICENSE.md release/windows64/LICENSE.md

# build windows
env GOOS=windows GOARCH=386 go build -o release/windows32/hushtache.exe main.go
cp README.md release/windows32/README.md
cp LICENSE.md release/windows32/LICENSE.md