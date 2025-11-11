package utils

func Assert(cond bool, message string) {
	if !cond {
		panic(message)
	}
}
