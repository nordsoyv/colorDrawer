package strategy

type ColorStrategy interface {
	GenerateImage(chan bool, chan ImageUpdate)
}

type ImageUpdate struct {
	X, Y    int
	R, G, B byte
}
