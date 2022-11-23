# Trojan Panel Core

Trojan Panel核心

# 支持的节点类型

1. Xray
2. Trojan Go
3. Hysteria
4. NaiveProxy

默认数据处理：

1. 读取/写入account中username,pass、hash、quota、download、upload、ip_limit、download_speed_limit、upload_speed_limit
   pass、hash需要进行hash，quota、upload、download、download_speed_limit、upload_speed_limit单位是byte

主要逻辑：

1. api实时更新（数据库同步至应用）
    - 删除场景 条件：account.quota = 0 or account.download + account.upload >= account.quota and account.quota > 0
    - 添加场景 如果存在则不操作，如果不存在则添加：account.download + account.upload <
      account.quota or account.quota < 0
2. 定时更新account.download,account.upload
3. account.quota=0，则禁用用户

# 建表语句示例

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

## 版本对应关系

| Trojan Panel | Trojan Panel Core | Xray   | Trojan Go | Hysteria | Caddy  |
|--------------|-------------------|--------|-----------|----------|--------|
| v1.2.0       | v1.2.0            | v1.6.0 | v0.10.6   | v1.2.1   | v2.6.2 |

# 编译命令

[compile.bat](./compile.bat)

# Author

[@jonssonyan](https://twitter.com/jonssonyan)

# Community

- Telegram Channel: [Trojan Panel](https://t.me/TrojanPanel)
- Telegram Group: [Trojan Panel交流群](https://t.me/TrojanPanelGroup)

# Thanks

- [trojan-gfw](https://github.com/trojan-gfw/trojan)
- [trojan-go](https://github.com/p4gefau1t/trojan-go)
- [hysteria](https://github.com/HyNetwork/hysteria)
