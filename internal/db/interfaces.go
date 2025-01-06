package db

type DBNameProvider interface {
	Provide() string
}
