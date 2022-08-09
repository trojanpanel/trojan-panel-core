# Trojan Panel Core

Trojan Panel核心

# 支持的节点类型

只会读取/写入password、quota、download、upload。password需要进行加盐对称加密，quota、upload、download单位是byte

1. Xray
2. Trojan Go
3. Hysteria

主要逻辑：

1. api实时更新（应用同步至数据库）：download、upload
2. api实时更新（数据库至应用）：根据password实时更新，删除：download + upload >= quota && quota >= 0；查询如果存在则不操作，如果不存在则添加：download + upload <
   quota || quota < 0
3. 禁用用户：set quota = 0,download = 0,upload = 0
4. 重设用户流量：set download = 0,upload = 0 然后api删除用户

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
