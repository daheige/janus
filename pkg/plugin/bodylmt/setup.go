package bodylmt

import (
	"github.com/hellofresh/janus/pkg/plugin"
	"github.com/hellofresh/janus/pkg/proxy"
)

// Config represents the Body Limit configuration
type Config struct {
	Limit string `json:"limit"`
}

func init() {
	plugin.RegisterPlugin("body_limit", plugin.Plugin{
		Action: setupBodyLimit,
	})
}

func setupBodyLimit(def *proxy.RouterDefinition, rawConfig plugin.Config) error {
	var config Config
	err := plugin.Decode(rawConfig, &config)
	if err != nil {
		return err
	}

	def.AddMiddleware(NewBodyLimitMiddleware(config.Limit))
	return nil
}
