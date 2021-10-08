package board

type MenuItem interface {
	Name() string
	RunOrDie(string) // username
}
