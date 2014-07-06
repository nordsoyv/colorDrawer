package workSurface

import "image/color"

type pixel struct {
	Color color.RGBA
	Used  bool
}

type Surface struct{
	pixels [][]pixel
	Size   int
}

func New(size int) Surface {
	topLevel := make([][] pixel, size)
	for i := range topLevel {
		topLevel[i] = make([]pixel, size)
	}
	return Surface{topLevel, size}
}



