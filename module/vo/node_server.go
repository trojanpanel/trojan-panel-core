package vo

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type NodeServerGroupVo struct {
	CpuUsed  float64 `json:"cpuUsed"`
	MemUsed  float64 `json:"memUsed"`
	DiskUsed float64 `json:"diskUsed"`
}

var messageInfo = make([]protoimpl.MessageInfo, 1)

func (n *NodeServerGroupVo) ProtoReflect() protoreflect.Message {
	mi := &messageInfo[0]
	if protoimpl.UnsafeEnabled && n != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(n))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(n)
}
