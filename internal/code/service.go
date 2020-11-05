package code

type CodeService struct{}

func NewService() (*CodeService, error) {
	return &CodeService{}, nil
}