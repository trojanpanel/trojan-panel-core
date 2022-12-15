go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
::Linux 386
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
xcaddy build --output build/naiveproxy-linux-386 --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
::Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
xcaddy build --output build/naiveproxy-linux-amd64 --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
::Linux armv6
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=6
xcaddy build --output build/naiveproxy-linux-armv6 --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
::Linux armv7
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
xcaddy build --output build/naiveproxy-linux-armv7 --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
::Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
xcaddy build --output build/naiveproxy-linux-arm64 --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
::Linux ppc64le
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=ppc64le
xcaddy build --output build/naiveproxy-linux-ppc64le --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive
::Linux s390x
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=s390x
xcaddy build --output build/naiveproxy-linux-s390x --with github.com/caddyserver/forwardproxy@caddy2=github.com/klzgrad/forwardproxy@naive