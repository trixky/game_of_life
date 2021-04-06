package main

func next(cells *[cellNumber][cellNumber]bool) {
	// ---------- calculating the current proximity of cells
	cells_point := [cellNumber][cellNumber]int8{}

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if (*cells)[y][x] {
				if y > 0 {
					cells_point[y-1][x]++ // top
					if x > 0 {
						cells_point[y-1][x-1]++ // top left
					}
					if x < 49 {
						cells_point[y-1][x+1]++ // top right
					}
				}
				if y < 49 {
					cells_point[y+1][x]++ // bottom
					if x > 0 {
						cells_point[y+1][x-1]++ // bottom left
					}
					if x < 49 {
						cells_point[y+1][x+1]++ // bottom right
					}
				}
				if x > 0 {
					cells_point[y][x-1]++ // left
				}
				if x < 49 {
					cells_point[y][x+1]++ // right
				}
			}
		}
	}

	// ---------- refreshing the next pattern
	// ---------- according to the proximity of the cells of the previous generation
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if (*cells)[y][x] {
				if cells_point[y][x] != 2 && cells_point[y][x] != 3 {
					// ---------- cell die
					(*cells)[y][x] = false
				}
				// ---------- cell stays alive
			} else if cells_point[y][x] == 3 {
				(*cells)[y][x] = true
				// ---------- cell born
			}
		}
	}
}
