package entity

type UserCredentials struct {
	ServiceName     string
	ServiceLogin    string
	ServicePassword string
	Meta            string
}

type UserText struct {
	Title string
	Body  string
	Meta  string
}

type UserCard struct {
	CardHolder         string
	CardNumber         string
	CardValidityPeriod string
	CvcCode            string
	Meta               string
}
