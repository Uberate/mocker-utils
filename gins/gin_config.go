package gins

// WebConfig is the gin server. The i18n use the gin server.
type WebConfig struct {
	Addr []string `json:"Addr" yaml:"Addr" mapstructure:"addr"`

	// Mod is the mod of gin
	//
	// options:
	// - debug
	// - release
	// - test
	Mod string `json:"mod" yaml:"modk" mapstructure:"mod"`
}
