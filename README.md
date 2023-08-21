# Trojan Panel Core

Trojan Panel Core

# Supported node types

1. Xray
2. Trojan Go
3. Hysteria
4. NaiveProxy

Default data processing：

1. Read/write username, pass, hash, quota, download, upload, ip_limit, download_speed_limit, upload_speed_limit in
   account. pass, hash needs to be hashed, quota, upload, download, download_speed_limit, upload_speed_limit unit is
   byte

Main logic：

1. API real-time update (database to application) valid account: account.quota < 0 or account.download +
   account.upload < account.quota
2. Regularly update account.download, account.upload
3. account.quota=0, the user is disabled

# Create database table statement example

```sql
create table trojan_panel_db.account
(
    id                   bigint(10) unsigned auto_increment comment 'auto increment primary key'
        primary key,
    username             varchar(64) default '' not null comment 'login username',
    pass                 varchar(64) default '' not null comment 'login password',
    hash                 varchar(64) default '' not null comment 'hash of pass',
    quota                bigint      default 0  not null comment 'quota unit/byte',
    download             bigint unsigned default 0 not null comment 'download unit/byte',
    upload               bigint unsigned default 0 not null comment 'upload unit/byte',
    ip_limit             tinyint(2) unsigned default 3 not null comment 'limit the number of IP devices',
    download_speed_limit bigint unsigned default 0 not null comment 'download speed limit unit/byte',
    upload_speed_limit   bigint unsigned default 0 not null comment 'upload speed limit unit/byte',
);
```

## Version relationship

| Trojan Panel Core | Xray   | Trojan Go | Hysteria | Caddy（NaiveProxy） |
|-------------------|--------|-----------|----------|-------------------|
| v2.1.1            | v1.8.0 | v0.10.6   | v1.3.4   | v2.6.4            |
| v2.1.2            | v1.8.0 | v0.10.6   | v1.3.4   | v2.6.4            |

## Prevent circular dependencies

router->api->middleware->app->service/dao->core

# Compile command

[compile.bat](./compile.bat)

# Author

[jonssonyan](https://github.com/jonssonyan)

# Community

- Telegram Channel: [Trojan Panel](https://t.me/TrojanPanel)

# Thanks

- [trojan-gfw](https://github.com/trojan-gfw/trojan)
- [trojan-go](https://github.com/p4gefau1t/trojan-go)
- [hysteria](https://github.com/HyNetwork/hysteria)
- [naiveproxy](https://github.com/klzgrad/naiveproxy)
