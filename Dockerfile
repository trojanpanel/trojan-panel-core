FROM alpine:3.15

LABEL maintainer="jonsosnyan <https://jonssonyan.com>"

WORKDIR /tpdata/trojan-panel-core/

ENV TZ=Asia/Shanghai
ENV GIN_MODE=release

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

COPY build/trojan-panel-core-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} trojan-panel-core

# Set apk China mirror
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash tzdata ca-certificates nftables \
    && rm -rf /var/cache/apk/* \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && chmod +x /tpdata/trojan-panel-core/trojan-panel-core

CMD ["./tpdata/trojan-panel-core/trojan-panel-core"]