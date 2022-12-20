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

	cubes := make(map[string]int, 0)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		cubes[input] = 1
	}

	log.Println("cubes", cubes)
	log.Println("surface area:", calculateSurfaceArea(cubes))

	// cubeGrid := initializeGrid(maxSize, cubes)
	// log.Println("cubeGrid", cubeGrid)

	// openSurfaces := make(map[string]int, 0)

	// openSurfaceCount := countSurfaces(cubeGrid, openSurfaces)
	// total := len(cubeLocations) * 6
	// log.Println("total surface count", total)
	// log.Println("open surface area", openSurfaceCount)

	readFile.Close()
}

func calculateSurfaceArea(cubes map[string]int) int {
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

// source: https://github.com/hugseverycat/aoc2022/blob/1e879a8f374b915b12ac968e963bcaac4bac021b/day18.py
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
