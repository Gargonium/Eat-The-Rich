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
	mapSlice[1] = []int{BlankSpace, 2, 6, 7}
	mapSlice[2] = []int{BlankSpace, 1, 3, 7}
	mapSlice[3] = []int{BlankSpace, 2, 4, 8}
	mapSlice[4] = []int{ShopSpace, 3, 5, 9, 10}
	mapSlice[5] = []int{GoodSpace, 4, 6, 11}
	mapSlice[6] = []int{ShopSpace, 1, 12, 17, 16}
	mapSlice[7] = []int{BadSpace, 1, 2, 8, 12}
	mapSlice[8] = []int{GoodSpace, 7, 3, 9}
	mapSlice[9] = []int{BlankSpace, 8, 4, 15, 13, 14}
	mapSlice[10] = []int{BlankSpace, 4, 11, 15}
	mapSlice[11] = []int{BlankSpace, 10, 5, 20, 15}
	mapSlice[12] = []int{BlankSpace, 6, 7, 17}
	mapSlice[13] = []int{BadSpace, 9, 14, 18, 22}
	mapSlice[14] = []int{BlankSpace, 13, 9, 15}
	mapSlice[15] = []int{BadSpace, 9, 10, 11, 20, 19}
	mapSlice[16] = []int{ShopSpace, 6, 21}
	mapSlice[17] = []int{GambleSpace, 6, 12, 21}
	mapSlice[18] = []int{BlankSpace, 13, 22, 24}
	mapSlice[19] = []int{BlankSpace, 18, 15, 20, 25}
	mapSlice[20] = []int{GambleSpace, 11, 15, 19, 25, 26}
	mapSlice[21] = []int{BlankSpace, 16, 17, 22}
	mapSlice[22] = []int{BlankSpace, 21, 13, 18}
	mapSlice[23] = []int{GoodSpace, 16, 24}
	mapSlice[24] = []int{BlankSpace, 23, 18, 25}
	mapSlice[25] = []int{ShopSpace, 24, 19, 20, 26}
	mapSlice[26] = []int{BlankSpace, 25, 20, 3}
}
