package sender

type httpAPI interface {
}
type sender struct {
	api httpAPI
}

func New(api httpAPI) *sender {
	return &sender{
		api: api,
	}
}
