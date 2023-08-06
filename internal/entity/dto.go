package entity

// UserCredentials - full user credentials data
type UserCredentials struct {
	ServiceName     string
	ServiceLogin    string
	ServicePassword string
	Meta            string
}

// UserText - full user text data
type UserText struct {
	Title string
	Body  string
	Meta  string
}

// UserAuth - user auth data for login and register
type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// UserCard - user card full data stored in service
type UserCard struct {
	CardHolder         string
	CardNumber         string
	CardValidityPeriod string
	CvcCode            string
	Meta               string
}

// UserBinary - user binary full data stored in service
type UserBinary struct {
	Title      string
	B64Content string
	Meta       string
}

// CutCard - shorten user card data stored in service
// for list command representation
type CutCard struct {
	Id         uint32 `json:"id"`
	CardNumber string `json:"card-number"`
	Meta       string `json:"meta"`
}

// CutText - shorten user text data stored in service
// for list command representation
type CutText struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

// CutBinary - shorten user binary data stored in service
// for list command representation
type CutBinary struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

// CutCredentials - shorten user credentials data stored in service
// for list command representation
type CutCredentials struct {
	Id          uint32 `json:"id"`
	ServiceName string `json:"service-name"`
	Meta        string `json:"meta"`
}
