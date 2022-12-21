package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("part1.txt")
	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	cubes := make(map[string]bool, 0)

	maxC := 0
	for fileScanner.Scan() {
		input := fileScanner.Text()
		x, y, z := readCubeCoordinates(input)
		maxC = max(x, y, z, maxC)
		cubes[input] = true
	}

	// # Create a larger cube of "air cubes" which is every cube from
	// # -1 to maxC in the x, y, and z direction that isn't a lava cube
	// air_cubes = {}
	airCubes := make(map[string]bool, 0)
	minC := -1

	for x := minC; x <= maxC; x++ {
		for y := minC; y <= maxC; y++ {
			for z := minC; z <= maxC; z++ {
				cubeIndex := getCubeKeyFromCoords(x, y, z)
				if _, ok := cubes[cubeIndex]; !ok {
					airCubes[cubeIndex] = false
				}
			}
		}
	}

	log.Println("cubes", cubes)

	// # This stuff is for part 2, where we flood fill starting with a far corner
	startCoord := "-1, -1, -1"
	floodFill(startCoord, airCubes)
	innerCubes := make(map[string]bool, 0)
	for k, v := range airCubes {
		if !v {
			innerCubes[k] = false
		}
	}
	innerCubesSurfaceArea := calculateSurfaceArea(innerCubes)

	part1SurfaceArea := calculateSurfaceArea(cubes)
	part2SurfaceArea := part1SurfaceArea - innerCubesSurfaceArea

	log.Println("part 1 surface area:", part1SurfaceArea)
	log.Println("part 2 surface area:", part2SurfaceArea)

	// cubeGrid := initializeGrid(maxSize, cubes)
	// log.Println("cubeGrid", cubeGrid)

	// openSurfaces := make(map[string]int, 0)

	// openSurfaceCount := countSurfaces(cubeGrid, openSurfaces)
	// total := len(cubeLocations) * 6
	// log.Println("total surface count", total)
	// log.Println("open surface area", openSurfaceCount)

	readFile.Close()
}

func calculateSurfaceArea(cubes map[string]bool) int {
	area := 0
	for k := range cubes {
		x, y, z := readCubeCoordinates(k)

		emptySurfaces := 0
		possibleNeighbors := []string{
			getCubeKeyFromCoords(x+1, y, z),
			getCubeKeyFromCoords(x-1, y, z),
			getCubeKeyFromCoords(x, y+1, z),
			getCubeKeyFromCoords(x, y-1, z),
			getCubeKeyFromCoords(x, y, z+1),
			getCubeKeyFromCoords(x, y, z-1),
		}

		//  For each possible neighboring coord, if it doesn't have a neighbor
		//  in the cubes iterable, then increase the surface area by 1
		for _, p := range possibleNeighbors {
			if _, ok := cubes[p]; !ok {
				log.Println("emptry surface for", "(", x, y, z, ")", "at", p)
				emptySurfaces++
			}
		}
		area += emptySurfaces
	}
	return area
}

func getCubeKeyFromCoords(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func floodFill(startCoords string, airCubes map[string]bool) {
	// # Goes thru all cubes in air_c connected to start_c and "fills" them
	// # Unfilled cubes are False, filled cubes are True

	queue := make([]string, 0)
	// # initialize queue with start coordinate
	queue = append(queue, startCoords)

	for len(queue) > 0 {
		// # Pull from the front of the queue
		coord := queue[0]
		queue = queue[1:]
		// # Set it to filled
		airCubes[coord] = true
		// # Find all its unfilled, non-lava neighbors and add to queue
		x, y, z := readCubeCoordinates(coord)
		directions := []string{
			getCubeKeyFromCoords(x+1, y, z),
			getCubeKeyFromCoords(x-1, y, z),
			getCubeKeyFromCoords(x, y+1, z),
			getCubeKeyFromCoords(x, y-1, z),
			getCubeKeyFromCoords(x, y, z+1),
			getCubeKeyFromCoords(x, y, z-1),
		}

		for _, d := range directions {
			if isEmpty, ok := airCubes[d]; ok {
				dInQueue := false
				for _, val := range queue {
					if val == d {
						dInQueue = true
					}
				}
				if !isEmpty && !dInQueue { // and d not in queue
					queue = append(queue, d)
				}
			}
		}
	}
}

// source: https://github.com/hugseverycat/aoc2022/blob/1e879a8f374b915b12ac968e963bcaac4bac021b/day18.py
// def flood_fill(start_c: tuple, air_c: dict):
//     # Goes thru all cubes in air_c connected to start_c and "fills" them
//     # Unfilled cubes are False, filled cubes are True

//     queue = deque()
//     # initialize queue with start coordinate
//     queue.append(start_c)

//     while queue:
//         # Pull from the front of the queue
//         coord = queue.popleft()
//         # Set it to filled
//         air_c[coord] = True
//         # Find all its unfilled, non-lava neighbors and add to queue
//         cx, cy, cz = coord
//         directions = [(cx - 1, cy, cz), (cx + 1, cy, cz), (cx, cy + 1, cz),
//                       (cx, cy - 1, cz), (cx, cy, cz + 1), (cx, cy, cz - 1)]

//         for d in directions:
//             if d in air_c and not air_c[d] and d not in queue:
//                 queue.append(d)

// def get_surface_area(cubes):
//     # Takes an iterable of cubes and calculates its surface area
//     s_area = 0
//     for this_cube in cubes:
//         x, y, z = this_cube

//         total_neighbors = 0
//         possible_neighbors = [(x+1, y, z), (x-1, y, z), (x, y+1, z),
//                               (x, y-1, z), (x, y, z+1), (x, y, z-1)]

//         # For each possible neighboring coord, if it doesn't have a neighbor
//         # in the cubes iterable, then increase the surface area by 1
//         for p in possible_neighbors:
//             if p not in cubes:
//                 total_neighbors += 1

//         s_area += total_neighbors
//     return s_area

func readCubeCoordinates(input string) (int, int, int) {
	coords := strings.Split(input, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return x, y, z
}

func max(a, b, c, d int) int {
	max := a
	if b >= a {
		max = b
	}
	if c >= max {
		max = c
	}
	if d >= max {
		max = d
	}
	return max
}

func initializeGrid(size int, cubeLocations map[string]int) [][][]int {
	grid := make([][][]int, size)
	for i := range grid {
		grid[i] = make([][]int, size)
		for j := range grid[i] {
			grid[i][j] = make([]int, size)
			for k := range grid[i][j] {
				key := fmt.Sprintf("%d,%d,%d", i, j, k)
				if cubeLocations[key] == 1 {
					grid[i][j][k] = 1
				}
			}
		}
	}
	return grid
}

// countSurfaces traverses the 3d grid running bfs on each item.
func countSurfaces(grid [][][]int, openSurfaces map[string]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for k := 0; k < len(grid[0][0]); k++ {
				if grid[i][j][k] == 1 {
					dfs(grid, i, j, k, openSurfaces)
				}
			}
		}
	}
	log.Println("openSurfaces", openSurfaces)
	count := 0
	for _, v := range openSurfaces {
		count += v
	}

	return count
}

func dfs(grid [][][]int, x, y, z int, openSurfaces map[string]int) {
	// check bounds
	if !boundsCheck(x, y, z, grid) {
		log.Println("dfs skipping", "(", x, y, z, ")")
		return
	}

	// mark as visited
	grid[x][y][z] = -1

	log.Println("dfs checking sides for (", x, y, z, ")")
	// count empty sides
	// 		call dfs on non-empty sides

	// check 6 sides (if 2 coords equal they are connected on that surface)
	// directions := [6]int{{}}
	dx := [6]int{1, -1, 0, 0, 0, 0}
	dy := [6]int{0, 0, 1, -1, 0, 0}
	dz := [6]int{0, 0, 0, 0, 1, -1}
	for i := 0; i < 6; i++ {
		// 1,1,1
		// 1,2,1
		s := fmt.Sprintf("%d,%d,%d", x, y, z)
		if !isWithinGrid(x+dx[i], y+dy[i], z+dz[i], grid) {
			// this is a wall
			log.Println("(", x, y, z, ")", "counting wall towards", "(", x+dx[i], y+dy[i], z+dz[i], ")")
			openSurfaces[s]++
			continue
		}

		if grid[x+dx[i]][y+dy[i]][z+dz[i]] == 0 {
			log.Println("(", x, y, z, ")", "is open towards", "(", x+dx[i], y+dy[i], z+dz[i], ")")
			openSurfaces[s]++
		} else if grid[x+dx[i]][y+dy[i]][z+dz[i]] == 1 {
			dfs(grid, x+dx[i], y+dy[i], z+dz[i], openSurfaces)
		}
	}
}

func boundsCheck(x, y, z int, grid [][][]int) bool {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 || z >= len(grid[0][0]) || z < 0 || grid[x][y][z] != 1 {
		log.Println("bounds check failed", "(", x, y, z, ")")
		return false
	}
	return true
}

func isWithinGrid(x, y, z int, grid [][][]int) bool {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 || z >= len(grid[0][0]) || z < 0 {
		log.Println("is open check failed", "(", x, y, z, ")")
		return false
	}

	return true
}
