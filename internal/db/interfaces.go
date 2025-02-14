package db

type NameProvider interface {
	Provide() string
}
