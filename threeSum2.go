package main

func threeSum(nums []int) [][]int {
	ret := [][]int{}

	// newnums := []int{}
	// flag := false
	// nummap := map[int]map[int]bool{}
	// for i, v := range nums {
	// 	if _, ok := nummap[v]; !ok {
	// 		nummap[v] = map[int]bool{}
	// 	}
	// 	if len(nummap[v]) < 3 {
	// 		nummap[v][i] = true
	// 		newnums = append(newnums, v)
	// 	} else {
	// 		flag = true
	// 	}
	// }
	// if flag {
	// 	nums = newnums
	// 	nummap = map[int]map[int]bool{}
	// 	for i, v := range nums {
	// 		if _, ok := nummap[v]; !ok {
	// 			nummap[v] = map[int]bool{}
	// 		}
	// 		nummap[v][i] = true
	// 	}
	// }
	// sum := 0
	hash := map[[3]int]bool{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i {
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						if nums[i]+nums[j]+nums[k] == 0 {
							tmp := [3]int{nums[i], nums[j], nums[k]}
							if tmp[0] > tmp[1] {
								tmp[0], tmp[1] = tmp[1], tmp[0]
							}
							if tmp[1] > tmp[2] {
								tmp[1], tmp[2] = tmp[2], tmp[1]
								if tmp[0] > tmp[1] {
									tmp[0], tmp[1] = tmp[1], tmp[0]
								}
							}
							hash[tmp] = true
						}
					}
				}
			}
		}
	}

	for k, _ := range hash {
		ret = append(ret, k[:])
	}

	return ret
}

func main() {

}
