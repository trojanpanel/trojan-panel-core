::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build -o build/trojan-go-linux-386 -tags "full"
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/trojan-go-linux-amd64 -tags "full"
::Linux armv6
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
go build -o build/trojan-go-linux-armv6 -tags "full"
::Linux armv7
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o build/trojan-go-linux-armv7 -tags "full"
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o build/trojan-go-linux-arm64 -tags "full"
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
go build -o build/trojan-go-linux-ppc64le -tags "full"
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
go build -o build/trojan-go-linux-s390x -tags "full"