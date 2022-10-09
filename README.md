# Trojan Panel Core

Trojan Panel核心

# 支持的节点类型

1. Xray
2. Trojan Go
3. Hysteria

默认数据处理：

1. 读取/写入account中username、pass、quota、download、upload、ip_limit、download_speed_limit、upload_speed_limit
   pass需要进行加盐对称加密，quota、upload、download、download_speed_limit、upload_speed_limit单位是byte

主要逻辑：

1. api实时更新（数据库同步至应用）
    - 删除场景 条件：account.download + account.upload >= account.quota and account.quota >= 0
    - 添加场景 如果存在则不操作，如果不存在则添加：account.download + account.upload <
      account.quota || account.quota < 0
2. 定时更新account.download,account.upload

## 版本对应关系

| Trojan Panel | Trojan Panel Core | Xray   | Trojan Go | Hysteria |
|--------------|-------------------|--------|-----------|----------|
| v1.2.0       | v1.2.0            | v1.6.0 | v0.10.6   | v1.2.1   |

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
