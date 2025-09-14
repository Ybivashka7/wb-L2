package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

// вот так можно исправить функцию и будет выведено "ok"
// func test() error {
//
//		return nil
//	}
func test() *customError {

	return nil
}

func main() {
	var err error
	err = test()
	// при сравнении err не nil в у нее value равно nil, но сам err не nil там хранится тип
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
