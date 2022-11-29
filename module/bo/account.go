package bo

type AccountUpdateBo struct {
	Pass     string
	Download int
	Upload   int
}

type AccountBo struct {
	Username string
	Pass     string
}
