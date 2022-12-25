package utils

func PanicOnErr(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
