package main

import (
	"fmt"
	"sort"
)

type TimeInfo struct {
	time uint64
	idx  uint64
}

func main() {
	var testsAmount int
	fmt.Scan(&testsAmount)

	for i := 0; i < testsAmount; i++ {
		var usrAmount uint64
		var usrIncrement uint64 = 0
		fmt.Scan(&usrAmount)

		usrsInfoV := make([]TimeInfo, usrAmount, usrAmount)
		usrsChain := make([]uint64, usrAmount, usrAmount)

		for ; usrIncrement < usrAmount; usrIncrement++ {
			var usrTime uint64
			fmt.Scan(&usrTime)
			usrsInfoV[usrIncrement] = TimeInfo{usrTime, usrIncrement}
		}

		sort.Slice(usrsInfoV, func(i, j int) bool {
			return usrsInfoV[i].time < usrsInfoV[j].time
		})

		// fmt.Println("\nsorted")
		// fmt.Println(usrsInfoV)

		var place uint64 = 1

		usrsChain[usrsInfoV[0].idx] = 1
		prev := usrsInfoV[0]

		usrIncrement = 1
		for ; usrIncrement < usrAmount; usrIncrement++ {
			current := usrsInfoV[usrIncrement]
			// fmt.Println(current.time, prev.time)
			if current.time-prev.time > 1 {
				place = usrIncrement + 1
			}
			usrsChain[current.idx] = place
			prev = current
		}

		usrIncrement = 0
		for ; usrIncrement < usrAmount; usrIncrement++ {
			fmt.Printf("%v ", usrsChain[usrIncrement])
		}
		// fmt.Println()
	}
}

