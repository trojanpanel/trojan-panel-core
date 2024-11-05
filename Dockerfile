FROM alpine:3.15

LABEL maintainer="jonsosnyan <https://jonssonyan.com>"

WORKDIR /trojan-core

ENV TZ=Asia/Shanghai
ENV GIN_MODE=release

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

COPY build/trojan-core-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} trojan-core

RUN apk update && apk add --no-cache bash tzdata ca-certificates nftables \
    && rm -rf /var/cache/apk/* \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && chmod +x /trojan-core/trojan-core

CMD ["./trojan-core"]