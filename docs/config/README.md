# 配置字段说明

```
{
  "databaseIp": "127.0.0.1", // 数据库IP
  "databasePort": 9507, // 数据库端口
  "databaseUser": "root", // 数据库用户名
  "databasePass": "123456", // 数据库密码
  "databaseName": "trojan_panel_db", // 数据库名称
  "accountTable": "account", // 账户表表名
  "redisHost": "127.0.0.1", // Redis IP
  "redisPort": 6378, // Redis 端口
  "redisPass": "123456", // Redis 密码
  "crtPath": "/tpdata/trojan-core/cert/trojan-core.crt", // crt 证书路径
  "keyPath": "/tpdata/trojan-core/cert/trojan-core.key", // key 证书路径
  "grpcPort": 8100, // gRPC 端口
  "serverPort": 8082 // Web 端口
}
```