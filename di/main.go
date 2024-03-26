package main

func main() {
	application := NewApplication(&Service{})
	application.Apply(1)
}

type OrderService interface {
	Apply(i int) error
}
type Service struct {
}

func (s Service) Apply(i int) error {
	return nil
}

type Application struct {
	os OrderService
}

func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}

func (a Application) Apply(i int) error {
	return a.os.Apply(i)
}
