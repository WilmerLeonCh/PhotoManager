package utils

func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](val T, err error) T {
	Throw(err)
	return val
}
