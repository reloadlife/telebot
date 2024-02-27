package config

type parsedConf struct {
	Settings  settings               `yaml:"settings" json:"settings"`
	Bot       botConfig              `yaml:"bot" json:"bot"`
	Config    map[string]any         `yaml:"config" json:"config,omitempty"`
	Locales   []string               `yaml:"locales" json:"locales,omitempty"`
	Handles   []handle               `yaml:"handles" json:"handles,omitempty"`
	Buttons   map[string]ButtonMap   `yaml:"buttons" json:"buttons,omitempty"`
	Keyboards map[string]keyboardMap `yaml:"keyboards" json:"keyboards,omitempty"`
}
