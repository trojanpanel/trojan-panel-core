::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build -o build/trojan-go-linux/386 -trimpath -ldflags "-s -w -buildid="
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/trojan-go-linux/amd64 -trimpath -ldflags "-s -w -buildid="
::Linux arm
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
go build -o build/trojan-go-linux/arm -trimpath -ldflags "-s -w -buildid="
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o build/trojan-go-linux/arm64 -trimpath -ldflags "-s -w -buildid="
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
go build -o build/trojan-go-linux/ppc64le -trimpath -ldflags "-s -w -buildid="
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
go build -o build/trojan-go-linux/s390x -trimpath -ldflags "-s -w -buildid="