func threeSum(nums []int) [][]int {
	ret := [][]int{}

	idxcont := 0
	newnums := []int{}
	//flag := false
	nummap := map[int]map[int]bool{}
	for i, v := range nums {
		if _, ok := nummap[v]; !ok {
			nummap[v] = map[int]bool{}
		}
		if len(nummap[v]) < 3 {
			nummap[v][i-idxcont] = true
			newnums = append(newnums, v)
		} else {
			idxcont++
		}
	}
	nums = newnums
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
	sum := 0
	hash := map[[3]int]bool{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				sum = nums[i] + nums[j]
				n := 0
				if v, ok := nummap[-sum]; ok {
					if v[i] {
						n++
					}
					if v[j] {
						n++
					}
					if len(v) > n {
						tmp := [3]int{-sum, nums[i], nums[j]}
						if tmp[0] > tmp[1] {
							tmp[0], tmp[1] = tmp[1], tmp[0]
						}
						if tmp[1] > tmp[2] {
							tmp[1], tmp[2] = tmp[2], tmp[1]
							if tmp[0] > tmp[1] {
								tmp[0], tmp[1] = tmp[1], tmp[0]
							}
						}
						if _, ok := hash[tmp]; !ok {
							ret = append(ret, tmp[:])
							hash[tmp] = true
						}
					}
				}
			}
		}
	}
	return ret
}
