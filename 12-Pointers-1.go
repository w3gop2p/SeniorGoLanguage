package main

type BigData struct {
	// 1 gb of memory
	// ..
	// ..
	// ..
}

func doSomethingWithData(data *BigData) {
	// manipulate the data inside this function
}

func main() {
	data := &BigData{} // 1gb

	for i := 0; i < 10000; i++ {
		// 8 bytes of pointer is copied into the function, this is why working with pointer is more performant in matter of timing
		doSomethingWithData(data)
	}
}
