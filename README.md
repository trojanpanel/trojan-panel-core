## Trojan Panel Core

Trojan Panel Core

## Supported node types

1. Xray
2. Trojan Go
3. Hysteria
4. NaiveProxy

default data processing：

1. 读取/写入account中username,pass、hash、quota、download、upload、ip_limit、download_speed_limit、upload_speed_limit
   pass、hash需要进行hash，quota、upload、download、download_speed_limit、upload_speed_limit单位是byte

main logic：

1. api实时更新（数据库同步至应用）有效账户：account.quota < 0 or account.download + account.upload < account.quota
2. 定时更新account.download,account.upload
3. account.quota=0，则禁用用户

## Create database tables

```sql
create table trojan_panel_db.account
(
    id                   bigint(10) unsigned auto_increment comment '自增主键'
        primary key,
    username             varchar(64) default '' not null comment '登录用户名',
    pass                 varchar(64) default '' not null comment '登录密码',
    hash                 varchar(64) default '' not null comment 'pass的hash',
    quota                bigint      default 0  not null comment '配额 单位/byte',
    download             bigint unsigned default 0 not null comment '下载 单位/byte',
    upload               bigint unsigned default 0 not null comment '上传 单位/byte',
    ip_limit             tinyint(2) unsigned default 3 not null comment '限制IP设备数',
    upload_speed_limit   bigint unsigned default 0 not null comment '上传限速 单位/byte',
    download_speed_limit bigint unsigned default 0 not null comment '下载限速 单位/byte',
);
```

## Version relationship

| Trojan Panel Core | Xray   | Trojan Go | Hysteria | Caddy（NaiveProxy） |
|-------------------|--------|-----------|----------|-------------------|
| v2.0.3            | v1.7.5 | v0.10.6   | v1.3.3   | v2.6.4            |
| v2.0.4            | v1.7.5 | v0.10.6   | v1.3.3   | v2.6.4            |
| v2.1.0            | v1.8.0 | v0.10.6   | v1.3.4   | v2.6.4            |
| v2.1.1            | v1.8.0 | v0.10.6   | v1.3.4   | v2.6.4            |
| v2.1.2            | v1.8.0 | v0.10.6   | v1.3.4   | v2.6.4            |

## Dependencies

router->api->middleware->app->service/dao->core

# Compile

[compile.bat](./compile.bat)

## Author

[jonssonyan](https://github.com/jonssonyan)

## Community

Telegram Channel: [Trojan Panel](https://t.me/TrojanPanel)

## Support

- [trojan](https://github.com/trojan-gfw/trojan)
- [trojan-go](https://github.com/p4gefau1t/trojan-go)
- [Xray-core](https://github.com/XTLS/Xray-core)
- [hysteria](https://github.com/HyNetwork/hysteria)
- [naiveproxy](https://github.com/klzgrad/naiveproxy)
