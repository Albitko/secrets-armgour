package entity

// CliSecrets - config for auth data stored in .token
type CliSecrets struct {
	UserName string `json:"user-name"`
	Key      string `json:"key"`
}
