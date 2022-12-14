package store

//go:generate mockgen  -source=store.go -destination=../mocks/mock.go
type Store interface {
	GetInfo(find string) (string, error)
	PostInfo(info string) (string, error)
}
