package np

import "sort"

func (a Float64OneDArray) Max() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	max := a.Arr[0]
	for _, value := range a.Arr {
		if value > max {
			max = value
		}
	}

	return &max, nil
}

func (a Float64OneDArray) Min() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	min := a.Arr[0]
	for _, value := range a.Arr {
		if value < min {
			min = value
		}
	}

	return &min, nil
}

func (a Float64OneDArray) Ptp() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	max, _ := a.Max()
	min, _ := a.Min()
	ptp := *max - *min

	return &ptp, nil
}

func (a Float64OneDArray) Sum() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	sum := float64(0)
	for _, value := range a.Arr {
		sum += value
	}

	return &sum, nil
}

func (a Float64OneDArray) Mean() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	sum, _ := a.Sum()
	mean := float64(*sum) / float64(len(a.Arr))

	return &mean, nil
}

func (a Float64OneDArray) Median() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	sort.SliceStable(a.Arr, func(i, j int) bool { return a.Arr[i] < a.Arr[j] })

	median := float64(0)
	if len(a.Arr)%2 != 0 {
		median = float64(a.Arr[len(a.Arr)/2])
		return &median, nil
	}

	median = float64(a.Arr[len(a.Arr)/2]+a.Arr[len(a.Arr)/2-1]) / 2

	return &median, nil
}

func (a Float64OneDArray) Mode() (*float64, error) {
	if err := validateArray(a.Err, len(a.Arr)); err != nil {
		return nil, err
	}

	numberMap := make(map[float64]int)
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

	return &mode, nil
}
