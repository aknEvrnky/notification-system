package api

type Application struct {
}

func (a *Application) GetVersion() string {
	return "0.0.1-dev"
}

func NewApplication() *Application {
	return &Application{}
}
