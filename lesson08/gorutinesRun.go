package lesson08

const n int = 3
const m int = 1

var chErr = make(chan int, m)
var chRun = make(chan int, n)

func GRun(f []func() error) {
	for i, r := range f {
		if len(chErr) != cap(chErr) {
			go run(r, i)
		}
	}
}

func run(rf func() error, v int) {
	chRun <- v
	ok := rf()
	if ok != nil {
		if len(chErr) == cap(chErr) {
			_, ok1 := <-chErr
			if !ok1 {
				close(chErr)
			}
			_, ok2 := <-chRun
			if !ok2 {
				close(chRun)
			}
		} else {
			chErr <- v
		}
	} else {
		<-chRun
	}
}
