package service

type RunConfig struct {
	Indexes []string `json:"indexes"`
	Tasks   []string `json:"tasks"`
	Start   uint64   `json:"start"`
	Stop    uint64   `json:"stop"`
	Step    uint64   `json:"step"`
}
