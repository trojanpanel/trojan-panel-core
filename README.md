# Trojan Panel Core

Trojan Panel核心

# 支持的节点类型

1. Xray
2. Trojan Go
3. Hysteria

默认数据处理：

1. 读取/写入account中quota、download、upload
2. 读取/写入users中api_port、password、download、upload。password需要进行加盐对称加密，upload、download单位是byte

主要逻辑：

1. api实时更新（应用同步至数据库）：set users.download = ?,users.upload = ? where users.api_port = ? and users.password = ?
2. api实时更新（数据库同步至应用）
    - 删除场景 条件：account.download + account.upload >= account.quota and account.quota >= 0
    - 添加场景 如果存在则不操作，如果不存在则添加：account.download + account.upload <
      account.quota || account.quota < 0
3. 重启场景：遍历users，根据配置文件启动相关的应用，并调api写入用户信息
4. 禁用用户：set account.quota = 0
5. 重设用户流量：api重设用户流量
6. 定时更新account.download,account.upload

# 编译命令

[compile.bat](./compile.bat)

## 版本对应关系

| Trojan Panel | Trojan Panel Core | Xray   | Trojan Go | Hysteria |
|--------------|-------------------|--------|-----------|----------|
| v1.2.0       | v1.2.0            | v1.5.5 | v1.10.6   | v1.1.0   |

# Community

- Telegram Channel: [Trojan Panel](https://t.me/TrojanPanel)
- Telegram Group: [Trojan Panel交流群](https://t.me/TrojanPanelGroup)

# Thanks

- [trojan-gfw](https://github.com/trojan-gfw/trojan)
- [trojan-go](https://github.com/p4gefau1t/trojan-go)
- [hysteria](https://github.com/HyNetwork/hysteria)
