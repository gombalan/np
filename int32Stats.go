package np

import "sort"

func (a Int32OneDArray) Max() (*int32, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	max := a.Arr[0]
	for _, value := range a.Arr {
		if value > max {
			max = value
		}
	}

	return int32Pointer(max), nil
}

func (a Int32OneDArray) Min() (*int32, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	min := a.Arr[0]
	for _, value := range a.Arr {
		if value < min {
			min = value
		}
	}

	return int32Pointer(min), nil
}

func (a Int32OneDArray) Ptp() (*int32, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	max, _ := a.Max()
	min, _ := a.Min()
	ptp := *max - *min

	return int32Pointer(ptp), nil
}

func (a Int32OneDArray) Sum() (*int32, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	sum := int32(0)
	for _, value := range a.Arr {
		sum += value
	}

	return int32Pointer(sum), nil
}

func (a Int32OneDArray) Mean() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	sum, _ := a.Sum()
	mean := float64(*sum) / float64(len(a.Arr))

	return float64Pointer(mean), nil
}

func (a Int32OneDArray) Median() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	sort.SliceStable(a.Arr, func(i, j int) bool { return a.Arr[i] < a.Arr[j] })

	median := float64(0)
	if len(a.Arr)%2 != 0 {
		median = float64(a.Arr[len(a.Arr)/2])
		return float64Pointer(median), nil
	}

	median = float64(a.Arr[len(a.Arr)/2]+a.Arr[len(a.Arr)/2-1]) / 2
	return float64Pointer(median), nil
}

func (a Int32OneDArray) Mode() (*int32, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	numberMap := make(map[int32]int)
	count := 0
	mode := a.Arr[0]
	for _, value := range a.Arr {
		if _, exist := numberMap[value]; !exist {
			numberMap[value] = 0
		}
		numberMap[value]++

		if numberMap[value] > count {
			count = numberMap[value]
			mode = value
		}
	}

	return int32Pointer(mode), nil
}
