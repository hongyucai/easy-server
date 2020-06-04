package providers

type run interface {
	register()
}

type App struct {
	config string
	id int
}

func (tmp *App) register()  {

}