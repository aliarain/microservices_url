package shortner

type MainService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
