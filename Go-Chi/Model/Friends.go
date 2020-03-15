package Model

type Friends struct {
	Success bool
	Friends []string `json: "friends"`
	Count int
}
