# Trojan Panel Core

Trojan Panel核心

# 支持的节点类型

## Xray

只会读取/写入password、quota、download、upload。password需要进行base64编码，quota、upload、download单位是byte

主要定时任务

1. 实时更新download、upload

2. 查询xray_user表，检查download + upload < quota，如果是，则授予连接，否则拒接，quota如果为负数则表示无限配额

## Trojan Go

## Hysteria

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
