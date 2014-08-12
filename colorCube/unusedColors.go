package colorCube

func (cube *ColorCube) FindUnusedColorInCubeN(startX, startY, startZ, n int) (foundIt bool, foundX, foundY , foundZ int) {
	foundIt, foundX, foundY, foundZ = cube.findUnusedColorsInTop(startX, startY, startZ, n)
	if foundIt {
		return
	}
	foundIt, foundX, foundY, foundZ = cube.findUnusedColorsInBottom(startX, startY, startZ, n)
	if foundIt {
		return
	}
	return

}

func (cube *ColorCube) findUnusedColorsInTop(startX, startY, startZ, distFromCenter int) (foundIt bool, foundX, foundY , foundZ int) {
	minX := startX - distFromCenter
	minZ := startZ - distFromCenter
	maxX := startX + distFromCenter
	maxZ := startZ + distFromCenter
	yPos := startY + distFromCenter
	for x := minX ; x <= maxX; x++ {
		for z := minZ; z <= maxZ; z++ {
			if !cube.IsUsed(x, yPos, z) {
				return true, x, yPos, z
			}
		}
	}
	return false, 0, 0, 0
}

func (cube *ColorCube) findUnusedColorsInBottom(startX, startY, startZ, distFromCenter int) (foundIt bool, foundX, foundY , foundZ int) {
	minX := startX - distFromCenter
	minZ := startZ - distFromCenter
	maxX := startX + distFromCenter
	maxZ := startZ + distFromCenter
	yPos := startY - distFromCenter
	for x := minX ; x <= maxX; x++ {
		for z := minZ; z <= maxZ; z++ {
			if !cube.IsUsed(x, yPos, z) {
				return true, x, yPos, z
			}
		}
	}
	return false, 0, 0, 0
}

func (cube *ColorCube) findUnusedColorsInLeft(startX, startY, startZ, distFromCenter int) (foundIt bool, foundX, foundY , foundZ int) {
	minY := startY - distFromCenter
	minZ := startZ - distFromCenter
	maxY := startY + distFromCenter
	maxZ := startZ + distFromCenter
	xPos := startX - distFromCenter
	for y := minY ; y <= maxY; y++ {
		for z := minZ; z <= maxZ; z++ {
			if !cube.IsUsed(xPos, y, z) {
				return true, xPos, y, z
			}
		}
	}
	return false, 0, 0, 0
}

func (cube *ColorCube) findUnusedColorsInRight(startX, startY, startZ, distFromCenter int) (foundIt bool, foundX, foundY , foundZ int) {
	minY := startY - distFromCenter
	minZ := startZ - distFromCenter
	maxY := startY + distFromCenter
	maxZ := startZ + distFromCenter
	xPos := startX + distFromCenter
	for y := minY ; y <= maxY; y++ {
		for z := minZ; z <= maxZ; z++ {
			if !cube.IsUsed(xPos, y, z) {
				return true, xPos, y, z
			}
		}
	}
	return false, 0, 0, 0
}
