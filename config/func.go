package config

func (c *config) GetConfig() map[string]interface{} {
	return c.conf.Config
}
