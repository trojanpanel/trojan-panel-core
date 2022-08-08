# Trojan Panel Core

Trojan Panel核心

# 支持的节点类型

只会读取/写入password、quota、download、upload。password需要进行加盐对称加密，quota、upload、download单位是byte

1. Xray
2. Trojan Go
3. Hysteria

主要方法：

1. 实时更新download、upload字段
2. 实时调用api根据password实时更新数据库中的用户至应用，条件：download + upload < quota
3. 重设用户：调用api删除用户，再调用api添加用户

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
