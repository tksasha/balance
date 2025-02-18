package valueobjects

type CurrentDateProvider interface {
	CurrentYear() int
	CurrentMonth() int
}
