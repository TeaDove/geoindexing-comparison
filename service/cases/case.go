package cases

type RunConfig struct {
	Indexes     []string `json:"indexes"`
	Tasks       []string `json:"tasks"`
	AmountStart uint64   `json:"amount_start"`
	AmountEnd   uint64   `json:"amount_end"`
	AmountStep  uint64   `json:"amount_step"`
}
