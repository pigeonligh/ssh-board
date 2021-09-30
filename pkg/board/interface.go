package board

type Board interface {
	Name() string
	PlayOrDie(string) // username
}
