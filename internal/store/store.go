package store

type Store interface {
	GetInfo(find string) (string, error)
	PostInfo(info string) (string, error)
}
