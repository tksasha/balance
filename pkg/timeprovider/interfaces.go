package timeprovider

type TimeProvider interface {
	IsCurrentMonth(month int) bool
	IsCurrentYear(year int) bool
	CurrentYear() int
}
