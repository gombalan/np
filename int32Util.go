package np

func (a Int32OneDArray) Len() (*int, error) {
	if a.Err != nil {
		return nil, a.Err
	}

	return intPointer(len(a.Arr)), nil
}
