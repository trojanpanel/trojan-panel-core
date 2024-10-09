::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build -o build/xray-linux-386 -trimpath -ldflags "-s -w -buildid=" ./main
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/xray-linux-amd64 -trimpath -ldflags "-s -w -buildid=" ./main
::Linux armv6
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
go build -o build/xray-linux-armv6 -trimpath -ldflags "-s -w -buildid=" ./main
::Linux armv7
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o build/xray-linux-armv7 -trimpath -ldflags "-s -w -buildid=" ./main
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o build/xray-linux-arm64 -trimpath -ldflags "-s -w -buildid=" ./main
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
go build -o build/xray-linux-ppc64le -trimpath -ldflags "-s -w -buildid=" ./main
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
go build -o build/xray-linux-s390x -trimpath -ldflags "-s -w -buildid=" ./main