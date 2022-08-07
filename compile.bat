go install mvdan.cc/garble@latest
::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
garble build -o build/trojan-panel-core-linux/386 -trimpath -ldflags "-s -w -buildid="
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
garble build -o build/trojan-panel-core-linux/amd64 -trimpath -ldflags "-s -w -buildid="
::Linux arm
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
garble build -o build/trojan-panel-core-linux/arm/v6 -trimpath -ldflags "-s -w -buildid="
garble build -o build/trojan-panel-core-linux/arm/v7 -trimpath -ldflags "-s -w -buildid="
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
garble build -o build/trojan-panel-core-linux/arm64 -trimpath -ldflags "-s -w -buildid="
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
garble build -o build/trojan-panel-core-linux/ppc64le -trimpath -ldflags "-s -w -buildid="
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
garble build -o build/trojan-panel-core-linux/s390x -trimpath -ldflags "-s -w -buildid="