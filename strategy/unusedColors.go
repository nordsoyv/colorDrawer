package strategy

import "github.com/nordsoyv/colorDrawer/colorCube"

func findUnusedColorInCubeN(startX, startY, startZ,n int, cube *colorCube.ColorCube)  (foundIt bool, foundX, foundY , foundZ int) {
	foundIt, foundX, foundY, foundZ = findUnusedColorsInTop(startX,startY,startZ,n,cube)
	if foundIt {
		return
	}
	foundIt, foundX, foundY, foundZ = findUnusedColorsInBottom(startX,startY,startZ,n,cube)
	if foundIt {
		return
	}
	return

}

func findUnusedColorsInTop(startX, startY, startZ, distFromCenter int, cube *colorCube.ColorCube) (foundIt bool, foundX, foundY , foundZ int) {
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

func findUnusedColorsInBottom(startX, startY, startZ, distFromCenter int, cube *colorCube.ColorCube) (foundIt bool, foundX, foundY , foundZ int) {
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

