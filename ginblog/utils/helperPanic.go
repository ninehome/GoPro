package utils

/* some helper functions */
func trueOrPanic(exp bool, what interface{}) {
	if exp == false {
		panic(what)
	}
}

func falseOrPanic(exp bool, what interface{}) {
	if exp == true {
		panic(what)
	}
}

func okOrPanic1(err error) {
	if err != nil {
		panic(err)
	}
}

func okOrPanic2(dontcare interface{}, err error) (interface{}, error) {
	if err != nil {
		panic(err)
	}
	return dontcare, nil
}
