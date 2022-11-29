package naiveproxy

type naiveProxyApi struct {
	apiPort uint
}

// NewNaiveProxyApi 初始化NaiveProxy Api
func NewNaiveProxyApi(apiPort uint) *naiveProxyApi {
	return &naiveProxyApi{
		apiPort: apiPort,
	}
}
