package main

var (
	mapSlice [][]int
)

const (
	BlankSpace  int = 0
	BadSpace    int = -1
	GoodSpace   int = -2
	ShopSpace   int = -3
	GambleSpace int = -4
)

func resolveSpaceEnum(spaceNum int) string {
	switch spaceNum {
	case 0:
		return "Blank Space"
	case -1:
		return "Bad Space"
	case -2:
		return "Good Space"
	case -3:
		return "Shop Space"
	case -4:
		return "Gamble Space"
	}
	return ""
}

func initMap() {
	mapSlice = append(mapSlice, []int{})                                // 0
	mapSlice = append(mapSlice, []int{BlankSpace, 2, 6, 7})             // 1
	mapSlice = append(mapSlice, []int{BlankSpace, 1, 3, 7})             // 2
	mapSlice = append(mapSlice, []int{BlankSpace, 2, 4, 8})             // 3
	mapSlice = append(mapSlice, []int{ShopSpace, 3, 5, 9, 10})          // 4
	mapSlice = append(mapSlice, []int{GoodSpace, 4, 6, 11})             // 5
	mapSlice = append(mapSlice, []int{ShopSpace, 1, 12, 17, 16})        // 6
	mapSlice = append(mapSlice, []int{BadSpace, 1, 2, 8, 12})           // 7
	mapSlice = append(mapSlice, []int{GoodSpace, 7, 3, 9})              // 8
	mapSlice = append(mapSlice, []int{BlankSpace, 8, 4, 15, 13, 14})    // 9
	mapSlice = append(mapSlice, []int{BlankSpace, 4, 11, 15})           // 10
	mapSlice = append(mapSlice, []int{BlankSpace, 10, 5, 20, 15})       // 11
	mapSlice = append(mapSlice, []int{BlankSpace, 6, 7, 17})            // 12
	mapSlice = append(mapSlice, []int{BadSpace, 9, 14, 18, 22})         // 13
	mapSlice = append(mapSlice, []int{BlankSpace, 13, 9, 15})           // 14
	mapSlice = append(mapSlice, []int{BadSpace, 9, 10, 11, 20, 19})     // 15
	mapSlice = append(mapSlice, []int{ShopSpace, 6, 21})                // 16
	mapSlice = append(mapSlice, []int{GambleSpace, 6, 12, 21})          // 17
	mapSlice = append(mapSlice, []int{BlankSpace, 13, 22, 24})          // 18
	mapSlice = append(mapSlice, []int{BlankSpace, 18, 15, 20, 25})      // 19
	mapSlice = append(mapSlice, []int{GambleSpace, 11, 15, 19, 25, 26}) // 20
	mapSlice = append(mapSlice, []int{BlankSpace, 16, 17, 22})          // 21
	mapSlice = append(mapSlice, []int{BlankSpace, 21, 13, 18})          // 22
	mapSlice = append(mapSlice, []int{GoodSpace, 16, 24})               // 23
	mapSlice = append(mapSlice, []int{BlankSpace, 23, 18, 25})          // 24
	mapSlice = append(mapSlice, []int{ShopSpace, 24, 19, 20, 26})       // 25
	mapSlice = append(mapSlice, []int{BlankSpace, 25, 20, 3})           // 26
}
