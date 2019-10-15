package tee

func Tee(in chan string) (<-chan string, <-chan string) {
	out1, out2 := make(chan string), make(chan string)
	go func() {
		defer close(out1)
		defer close(out2)
		for v := range in {
			out1, out2 := out1, out2
			for i := 0; i < 2; i++ {
				select {
				case out1 <- v:
					out1 = nil
				case out2 <- v:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
