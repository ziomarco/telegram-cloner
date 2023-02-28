package http

type LoginResponse struct {
	Ok     bool `json:"ok""`
	Result struct {
		Token     string `json:"token""`
		AuthState string `json:"authorization_state""`
	} `json:"result""`
}

type ConfirmLoginResult struct {
	Ok     bool `json:"ok"`
	Result struct {
		AuthorizationState      string `json:"authorization_state"`
		Token                   string `json:"token"`
		Timeout                 int    `json:"timeout"`
		PasswordHint            string `json:"password_hint"`
		HasRecoveryEmailAddress bool   `json:"has_recovery_email_address"`
	} `json:"result"`
}
