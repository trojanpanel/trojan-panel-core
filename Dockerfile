FROM alpine:3.15
LABEL maintainer="jonsosnyan <https://jonssonyan.com>"
RUN mkdir -p /tpdata/trojan-panel-core/
WORKDIR /tpdata/trojan-panel-core/
ENV mariadb_ip=127.0.0.1 \
    mariadb_port=3306 \
    mariadb_user=root \
    mariadb_pas=123456 \
    database=trojan_panel_db \
    account_table=account \
    redis_host=127.0.0.1 \
    redis_port=6379 \
    redis_pass=123456 \
    crt_path=/tpdata/trojan-panel-core/cert/trojan-panel-core.crt \
    key_path=/tpdata/trojan-panel-core/cert/trojan-panel-core.key
ARG TARGETOS
ARG TARGETARCH
COPY build/trojan-panel-core-${TARGETOS}-${TARGETARCH} trojan-panel-core
ARG trojan_panel_core_version=v1.0.0
ENV TROJAN_PANEL_CORE_VERSION=${trojan_panel_core_version}
ARG BASE_URL=https://github.com/trojanpanel/install-script/releases/download/${trojan_panel_core_version}
ADD ${BASE_URL}/xray-${TARGETOS}-${TARGETARCH} bin/xray/xray-${TARGETOS}-${TARGETARCH}
ADD ${BASE_URL}/trojan-go-${TARGETOS}-${TARGETARCH} bin/trojango/trojan-go-${TARGETOS}-${TARGETARCH}
ADD ${BASE_URL}/hysteria-${TARGETOS}-${TARGETARCH} bin/hysteria/hysteria-${TARGETOS}-${TARGETARCH}
RUN chmod 777 bin/xray/xray-${TARGETOS}-${TARGETARCH}
RUN chmod 777 bin/trojango/trojan-go-${TARGETOS}-${TARGETARCH}
RUN chmod 777 bin/hysteria/hysteria-${TARGETOS}-${TARGETARCH}
# 国内环境开启以下注释 设置apk国内镜像
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
    -account-table=${account_table} \
    -redisHost=${redis_host} \
    -redisPort=${redis_port} \
    -redisPassword=${redis_pass} \
    -crt-path=${crt_path} \
    -key-path=${key_path}