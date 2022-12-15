::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build -o build/hysteria-linux-386 -tags=gpl -trimpath ./app/cmd/
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/hysteria-linux-amd64 -tags=gpl -trimpath ./app/cmd/
::Linux armv6
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
go build -o build/hysteria-linux-armv6 -tags=gpl -trimpath ./app/cmd/
::Linux armv7
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o build/hysteria-linux-armv7 -tags=gpl -trimpath ./app/cmd/
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o build/hysteria-linux-arm64 -tags=gpl -trimpath ./app/cmd/
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
go build -o build/hysteria-linux-ppc64le -tags=gpl -trimpath ./app/cmd/
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
go build -o build/hysteria-linux-s390x -tags=gpl -trimpath ./app/cmd/