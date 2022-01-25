package structs

type ServerInfo struct {
	Machine string `json:"machine"`
	OS      string `json:"os,omitempty"`
	Node    string `json:"node,omitempty"`
	PSQL    string `json:"psql,omitempty"`
	Redis   string `json:"redis,omitempty"`
	CPU     struct {
		Model string `json:"model"`
		Cores int    `json:"cores"`
	} `json:"cpu"`
	Mem struct {
		Total int `json:"total"`
	} `json:"mem"`
	FS struct {
		Total int `json:"total"`
		Used  int `json:"used"`
	} `json:"fs"`
	Net *struct {
		Interface string `json:"interface"`
	} `json:"net,omitempty"`
}
