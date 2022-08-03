# Xray Manage

只会读取/写入password、quota、download和upload字段。password必须通过SHA224散列，quota、upload、download 单位是byte

实时更新download、upload，quota如果为负数则表示无限配额

查询xray_user表，检查是否download + upload < quota; 如果是，则授予连接
