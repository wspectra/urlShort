package store

type Store interface {
	GetInfo(find string) string
	PostInfo(info string) string
}
