#!/bin/bash
export GOPATH=$(pwd)
go build -o release/hushtache main.go

# build folder
mkdir -p builds
rm builds/* || true

# build linux
env GOOS=linux GOARCH=amd64 go build -o release/linux64/hushtache main.go
chmod a+x release/linux64/hushtache
cp README.md release/linux64/README.md
cp LICENSE.md release/linux64/LICENSE.md
zip -r -j builds/linux64.zip release/linux64/*

# build linux
env GOOS=linux GOARCH=386 go build -o release/linux32/hushtache main.go
chmod a+x release/linux32/hushtache
cp README.md release/linux32/README.md
cp LICENSE.md release/linux32/LICENSE.md
zip -r -j builds/linux32.zip release/linux32/*

# build mac
env GOOS=darwin GOARCH=amd64 go build -o release/darwin64/hushtache main.go
chmod a+x release/darwin64/hushtache
cp README.md release/darwin64/README.md
cp LICENSE.md release/darwin64/LICENSE.md
zip -r -j builds/darwin64.zip release/darwin64/*

# build mac
env GOOS=darwin GOARCH=386 go build -o release/darwin32/hushtache main.go
chmod a+x release/darwin32/hushtache
cp README.md release/darwin32/README.md
cp LICENSE.md release/darwin32/LICENSE.md
zip -r -j builds/darwin32.zip release/darwin32/*

# build windows
env GOOS=windows GOARCH=amd64 go build -o release/windows64/hushtache.exe main.go
chmod a+x release/windows64/hushtache
cp README.md release/windows64/README.md
cp LICENSE.md release/windows64/LICENSE.md
zip -r -j builds/windows64.zip release/windows64/*

# build windows
env GOOS=windows GOARCH=386 go build -o release/windows32/hushtache.exe main.go
chmod a+x release/windows32/hushtache
cp README.md release/windows32/README.md
cp LICENSE.md release/windows32/LICENSE.md
zip -r -j builds/windows32.zip release/windows32/*