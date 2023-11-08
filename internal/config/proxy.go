package config

type Proxy struct {
	// 是否开启
	Enabled *bool `default:"true" json:"enabled" yaml:"enabled" xml:"enabled" toml:"enabled"`
	// 主机
	Host string `json:"host" yaml:"host" xml:"host" toml:"host" validate:"required"`
	// 代理类型
	// nolint: lll
	Scheme string `default:"scheme" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme" validate:"required,oneof=socks4 socks5 http https"`
	// 目标
	Target string `json:"target" yaml:"target" xml:"target" toml:"target"`
	// 排除
	Exclude string `json:"exclude" yaml:"exclude" xml:"exclude" toml:"exclude"`
	// 代理认证用户名
	Username string `json:"username" yaml:"username" xml:"username" toml:"username"`
	// 代理认证密码
	Password string `json:"password" yaml:"password" xml:"password" toml:"password"`
}
