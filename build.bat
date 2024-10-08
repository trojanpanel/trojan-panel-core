::Windows amd64
::SET CGO_ENABLED=0
::SET GOOS=windows
::SET GOARCH=amd64
::go build -o build/trojan-core-windows-amd64.exe -trimpath -ldflags "-s -w -buildid="
::Mac amd64
::SET CGO_ENABLED=0
::SET GOOS=darwin
::SET GOARCH=amd64
::go build -o build/trojan-core-darwin-amd64 -trimpath -ldflags "-s -w -buildid="
::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build -o build/trojan-core-linux-386 -trimpath -ldflags "-s -w -buildid="
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/trojan-core-linux-amd64 -trimpath -ldflags "-s -w -buildid="
::Linux armv6
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
go build -o build/trojan-core-linux-armv6 -trimpath -ldflags "-s -w -buildid="
::Linux armv7
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o build/trojan-core-linux-armv7 -trimpath -ldflags "-s -w -buildid="
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o build/trojan-core-linux-arm64 -trimpath -ldflags "-s -w -buildid="
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
go build -o build/trojan-core-linux-ppc64le -trimpath -ldflags "-s -w -buildid="
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
go build -o build/trojan-core-linux-s390x -trimpath -ldflags "-s -w -buildid="