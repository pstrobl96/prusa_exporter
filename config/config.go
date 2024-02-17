package config

// Configuration struct for the configuration file prusa.yml
type Configuration struct {
	Exporter struct {
		ScrapeTimeout  int    `yaml:"scrape_timeout"`
		ReloadInterval int    `yaml:"reload_interval"`
		LogLevel       string `yaml:"log_level"`
		Syslog         struct {
			ListenAddress string `yaml:"listen_address"`
			Enabled       bool   `yaml:"enabled"`
		} `yaml:"syslog"`
	} `yaml:"exporter"`
	Printers []Printers `yaml:"printers"`
}

// Printers struct containing the printer configuration
type Printers struct {
	Address   string `yaml:"address"`
	Username  string `yaml:"username,omitempty"`
	Password  string `yaml:"password,omitempty"`
	Apikey    string `yaml:"apikey,omitempty"`
	Name      string `yaml:"name,omitempty"`
	Type      string `yaml:"type,omitempty"`
	Reachable bool
}
