package talosconfig

type Talosconfig struct {
	Contexts map[string]Taloscontext
	Context  string `json:"context"`
}

type Taloscontext struct {
	Ca        string   `json:"ca"`
	Crt       string   `json:"crt"`
	Key       string   `json:"key"`
	Endpoints []string `json:"endpoints"`
	Nodes     []string `json:"nodes"`
}
