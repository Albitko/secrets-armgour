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

type UserBinary struct {
	Title      string
	B64Content string
	Meta       string
}

type CutCard struct {
	Id         uint32 `json:"id"`
	CardNumber string `json:"card-number"`
	Meta       string `json:"meta"`
}

type CutText struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

type CutBinary struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

type CutCredentials struct {
	Id          uint32 `json:"id"`
	ServiceName string `json:"service-name"`
	Meta        string `json:"meta"`
}
