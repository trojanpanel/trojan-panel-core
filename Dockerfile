FROM alpine:3.15
LABEL maintainer="jonsosnyan <https://jonssonyan.com>"
WORKDIR /tpdata/trojan-panel-core/
ENV mariadb_ip=127.0.0.1 \
    mariadb_port=9507 \
    mariadb_user=root \
    mariadb_pas=123456 \
    database=trojan_panel_db \
    account_table=account \
    redis_host=127.0.0.1 \
    redis_port=6378 \
    redis_pass=123456 \
    crt_path=/tpdata/trojan-panel-core/cert/trojan-panel-core.crt \
    key_path=/tpdata/trojan-panel-core/cert/trojan-panel-core.key \
    grpc_port=8100 \
    server_port=8082 \
    TZ=Asia/Shanghai \
    GIN_MODE=release
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
COPY build/trojan-panel-core-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} trojan-panel-core
COPY build/xray-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} bin/xray/xray
COPY build/trojan-go-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} bin/trojango/trojan-go
COPY build/hysteria-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} bin/hysteria/hysteria
COPY build/naiveproxy-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} bin/naiveproxy/naiveproxy
COPY build/hysteria2-${TARGETOS}-${TARGETARCH}${TARGETVARIANT} bin/hysteria2/hysteria2
RUN chmod 777 bin/xray/xray
RUN chmod 777 bin/trojango/trojan-go
RUN chmod 777 bin/hysteria/hysteria
RUN chmod 777 bin/naiveproxy/naiveproxy
RUN chmod 777 bin/hysteria2/hysteria2
# Set apk China mirror
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add bash tzdata ca-certificates && \
    rm -rf /var/cache/apk/*
ENTRYPOINT chmod 777 ./trojan-panel-core && \
    ./trojan-panel-core \
    -host=${mariadb_ip} \
    -port=${mariadb_port} \
    -user=${mariadb_user} \
    -password=${mariadb_pas} \
    -database=${database} \
    -accountTable=${account_table} \
    -redisHost=${redis_host} \
    -redisPort=${redis_port} \
    -redisPassword=${redis_pass} \
    -crtPath=${crt_path} \
    -keyPath=${key_path} \
    -grpcPort=${grpc_port} \
    -serverPort=${server_port}