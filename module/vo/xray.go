package vo

type XrayStatsVo struct {
	Name     string `json:"name"`
	Download int    `json:"download"`
	Upload   int    `json:"upload"`
}
