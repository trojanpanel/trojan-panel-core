package bo

type NaiveProxyConfig struct {
	Admin   TypeMessage `json:"admin"`
	Logging TypeMessage `json:"logging"`
	Apps    Apps        `json:"apps"`
}

type Apps struct {
	Http Http `json:"http"`
	Tls  Tls  `json:"tls"`
}

type Http struct {
	Servers Servers `json:"servers"`
}

type Servers struct {
	Srv0 Srv0 `json:"srv0"`
}

type Srv0 struct {
	Listen                []string    `json:"listen"`
	Routes                []Route     `json:"routes"`
	TlsConnectionPolicies TypeMessage `json:"tls_connection_policies"`
	AutomaticHttps        TypeMessage `json:"automatic_https"`
}

type Tls struct {
	TlsCertificate TlsCertificate `json:"certificates"`
}

type TlsCertificate struct {
	LoadFiles []LoadFile `json:"load_files"`
}

type LoadFile struct {
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
}

type Route struct {
	Handle []Handle `json:"handle"`
}
type Handle struct {
	Handler      TypeMessage   `json:"handler"`
	HandleRoutes []TypeMessage `json:"routes"`
}

type RouteHandle struct {
	Handle []HandleAuth `json:"handle"`
}

type HandleAuth struct {
	AuthPassDeprecated string      `json:"auth_pass_deprecated"`
	AuthUserDeprecated string      `json:"auth_user_deprecated"`
	Handler            TypeMessage `json:"handler"`
	HideIp             TypeMessage `json:"hide_ip"`
	HideVia            TypeMessage `json:"hide_via"`
	ProbeResistance    TypeMessage `json:"probe_resistance"`
}
