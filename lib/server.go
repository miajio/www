package lib

type HttpServerParam struct {
	JwtKey    string `json:"jwtKey" toml:"jwtKey" yaml:"jwtKey" xml:"jwtKey"`
	Port      string `json:"port" toml:"port" yaml:"port" xml:"port"`
	UseHttps  bool   `json:"useHttps" toml:"useHttps" yaml:"useHttps" xml:"useHttps"`
	HttpsKey  string `json:"httpsKey" toml:"httpsKey" yaml:"httpsKey" xml:"httpsKey"`
	HttpsPem  string `json:"httpsPem" toml:"httpsPem" yaml:"httpsPem" xml:"httpsPem"`
	HttpsHost string `json:"httpsHost" toml:"httpsHost" yaml:"httpsHost" xml:"httpsHost"`
}
