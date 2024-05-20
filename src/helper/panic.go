package helper

func ErrorIfPanic(err error) {
	if err != nil {
		panic(err)
	}
}
