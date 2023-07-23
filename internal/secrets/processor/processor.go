package processor

type repository interface {
}

type processor struct {
	repo repository
}

func New(repo repository) *processor {
	return &processor{
		repo: repo,
	}
}
