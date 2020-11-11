package context

type Context struct {
	Config func(Config, error)
}
