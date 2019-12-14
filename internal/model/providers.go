package model

// Providers represents a provider configuration
type Providers struct {
	Image  []PrdImage  `yaml:"image,omitempty" json:",omitempty"`
	Docker []PrdDocker `yaml:"docker,omitempty" json:",omitempty"`
}

// PrdImage holds image provider configuration
type PrdImage Image

// PrdDocker holds docker provider configuration
type PrdDocker struct {
	ID             string `yaml:"id,omitempty" json:",omitempty"`
	Endpoint       string `yaml:"endpoint,omitempty" json:",omitempty"`
	ApiVersion     string `yaml:"api_version,omitempty" json:",omitempty"`
	TLSCertsPath   string `yaml:"tls_certs_path,omitempty" json:",omitempty"`
	TLSVerify      bool   `yaml:"tls_verify,omitempty" json:",omitempty"`
	SwarmMode      bool   `yaml:"swarm_mode,omitempty" json:",omitempty"`
	WatchByDefault bool   `yaml:"watch_by_default,omitempty" json:",omitempty"`
	WatchStopped   bool   `yaml:"watch_stopped,omitempty" json:",omitempty"`
}
