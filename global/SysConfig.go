package global

type SysConfig struct {
	ServiceConfig serviceConfig `toml:"ServiceConfig"`
}

type serviceConfig struct {
	Name string `toml:"name"`
	DisplayName string `toml:"displayName"`
	Description string `toml:"description"`
}