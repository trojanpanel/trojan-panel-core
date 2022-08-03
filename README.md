# Xray Manage

只会读取/写入password、quota、download、upload。password需要进行base64编码，quota、upload、download单位是byte

主要定时任务

1. 实时更新download、upload

2. 查询xray_user表，检查download + upload < quota，如果是，则授予连接，否则拒接，quota如果为负数则表示无限配额
