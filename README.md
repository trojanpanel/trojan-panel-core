# Trojan Panel Core

Trojan Panel Core

## Supported Proxy

1. Xray
2. Trojan Go
3. Hysteria1/Hysteria2
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

## Create database table statement example

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

[Release Notes](https://github.com/trojanpanel/install-script/blob/main/README_ARCHIVE_ZH.md#%E5%8F%91%E8%A1%8C%E8%AF%B4%E6%98%8E)

## Prevent circular dependencies

router->api->middleware->app->service/dao->core

## Build

Windows: [build.bat](build.bat)

Linux: [build.sh](build.sh)

## Other

Telegram Channel: https://t.me/jonssonyan_channel

You can subscribe to my channel on YouTube: https://www.youtube.com/@jonssonyan

## Support

- [trojan](https://github.com/trojan-gfw/trojan)
- [trojan-go](https://github.com/p4gefau1t/trojan-go)
- [Xray-core](https://github.com/XTLS/Xray-core)
- [hysteria](https://github.com/HyNetwork/hysteria)
- [naiveproxy](https://github.com/klzgrad/naiveproxy)