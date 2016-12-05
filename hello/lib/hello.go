package lib

type Hello struct {
	Language string
}

func (f *Hello) Hello() string {
	return "Hello, " + f.Language

}