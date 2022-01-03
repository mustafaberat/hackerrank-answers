

func cutTheSticks(arr []int32) []int32 {
	var cutLen []int32

	for true {
		minStick := findMinNumberIntoArray(arr)
		cutLen = append(cutLen, int32(len(arr)))
		arr = cutSticks(arr, minStick)
		if len(arr) == 0 {
			break
		}
	}

	return cutLen
}

func cutSticks(arr []int32, minStickLen int32) []int32 {
	var newArr []int32

	for i := range arr {
		if arr[i]-minStickLen > 0 {
			newArr = append(newArr, arr[i]-minStickLen)
		}
	}

	return newArr
}

func findMinNumberIntoArray(arr []int32) int32 {
	min := arr[0]

	for i, l := 1, len(arr); i < l; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}
