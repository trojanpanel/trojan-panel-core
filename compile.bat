go install mvdan.cc/garble@latest
::Windows amd64
::SET CGO_ENABLED=0
::SET GOOS=windows
::SET GOARCH=amd64
::go build -ldflags="-H windowsgui -s -w" -o build/trojan-panel-core-win-amd64.exe
::Mac amd64
::SET CGO_ENABLED=0
::SET GOOS=darwin
::SET GOARCH=amd64
::go build -ldflags "-s -w" -o build/trojan-panel-core-mac-amd64
::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
garble -literals build -o build/trojan-panel-core-linux/386
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
garble -literals build -o build/trojan-panel-core-linux/amd64
::Linux arm
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
garble -literals build -o build/trojan-panel-core-linux/arm/v6
garble -literals build -o build/trojan-panel-core-linux/arm/v7
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
garble -literals build -o build/trojan-panel-core-linux/arm64
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
garble -literals build -o build/trojan-panel-core-linux/ppc64le
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
garble -literals build -o build/trojan-panel-core-linux/s390x