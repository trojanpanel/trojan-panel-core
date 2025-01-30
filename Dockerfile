FROM alpine:3.15

LABEL maintainer="jonsosnyan <https://jonssonyan.com>"

WORKDIR /app

ENV TZ=Asia/Shanghai
ENV GIN_MODE=release
ENV TROJAN_PANEL_CORE_WEB_PORT=8082
ENV TROJAN_PANEL_CORE_GRPC_PORT=8083

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

COPY build/trojan-panel-core-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} trojan-panel-core

RUN apk update && apk add --no-cache bash tzdata ca-certificates nftables \
    && rm -rf /var/cache/apk/* \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && chmod +x /app/trojan-panel-core

CMD ["./trojan-panel-core"]