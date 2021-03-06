package vars

var Pass = PassType{
	HTTP: ":80",
	HTTPS: ":443",
}

type PassType struct {
	AlreadyRegistered string `json:"alreadyreg"` // Path to which to 302 after attempting to reregister
	CaptchaSecret     []byte `json:"captcha"`    // Google reCAPTCHA secret data
	DataBase          string `json:"db"`         // Path to the main sqlite3 database
	HTTP              string `json:"http"`       // Address on which to listen to HTTP traffic
	HTTPS             string `json:"https"`      // Address on which to listen to HTTPS traffic
	Registered        string `json:"registered"` // Path to which to 302 after registering
	TLSCertData       []byte `json:"cert"`       // TLS certificate and key file data
	TLSKeyData        []byte `json:"key"`        // TLS certificate and key file data

	CaptchaThresholdRegister float64 `json:"captcha_t_register"` // Fail to register below this score

	FeatureUserProfiles bool `json:"ft_user_profiles"` // Feature toggle for User Profiles
}
