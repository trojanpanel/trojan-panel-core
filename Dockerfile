FROM alpine:3.15
LABEL maintainer="jonsosnyan <https://jonssonyan.com>"
RUN mkdir -p /tpdata/trojan-panel-core/
WORKDIR /tpdata/trojan-panel-core/
ENV mariadb_ip=trojan-panel-mariadb \
    mariadb_port=3306 \
    mariadb_user=root \
    mariadb_pas=123456 \
    database=trojan_panel_db \
    account_table=account \
    redis_host=trojan-panel-redis \
    redis_port=6379 \
    redis_pass=123456 \
    crt_path=/tpdata/trojan-panel-core/cert/trojan-panel-core.crt \
    key_path=/tpdata/trojan-panel-core/cert/trojan-panel-core.key
ARG TARGETPLATFORM
COPY build/trojan-panel-core-${TARGETPLATFORM} trojan-panel-core
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