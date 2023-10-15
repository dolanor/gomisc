package pkg

func Yolo(a string) string {
	if len(a) == 256 {
		panic("on ne rigole pas")
	}
	if a == "" {
		a += "LOL"
	}
	a += "YOLO"
	return a
}
