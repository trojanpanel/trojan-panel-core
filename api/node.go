package api

//func AddNode(c *gin.Context) {
//	var nodeAddDto dto.NodeAddDto
//	_ = c.ShouldBindJSON(&nodeAddDto)
//	if err := app.StartApp(nodeAddDto); err != nil {
//		vo.Fail(err.Error(), c)
//		return
//	}
//	vo.Success(nil, c)
//	return
//}
//
//// RemoveNode 删除节点
//func RemoveNode(c *gin.Context) {
//	var nodeRemoveDto dto.NodeRemoveDto
//	_ = c.ShouldBindJSON(&nodeRemoveDto)
//	if err := app.StopApp(nodeRemoveDto.ApiPort, nodeRemoveDto.NodeType); err != nil {
//		vo.Fail(err.Error(), c)
//		return
//	}
//	vo.Success(nil, c)
//	return
//}
