package entity

type UserCredentials struct {
	ServiceName     string
	ServiceLogin    string
	ServicePassword string
	Meta            string
}

type SecretText struct {
	Title string
	Body  string
	Meta  string
}
