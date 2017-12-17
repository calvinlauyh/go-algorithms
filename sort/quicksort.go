package sort

import (
	"errors"
)

const (
	QUICKSORT_PIVOT_FIRST           = 1 << iota // pick first element as privot
	QUICKSORT_PIVOT_MIDDLE                      // pick middle element as pivot
	QUICKSORT_PIVOT_LAST                        // pick last element as pivot
	QUICKSORT_PIVOT_MEDIAN_OF_THREE             // pick median of first, middel and last element as pivot
	// TODO
	QUICKSORT_PIVOT_RANDOM // pick random element as pivot
	QUICKSORT_ORDER_ASC    // sort the list in ascending order
	QUICKSORT_ORDER_DESC   // sort the list in descending order

	QUICKSORT_PIVOT_DEFAULT = QUICKSORT_PIVOT_MEDIAN_OF_THREE
	QUICKSORT_ORDER_DEFAULT = QUICKSORT_ORDER_ASC
	QUICKSORT_DEFAULT       = QUICKSORT_PIVOT_DEFAULT | QUICKSORT_ORDER_DEFAULT
)

// interfaceToInt converts an interface containing integer to integer
func interfaceToInt(v interface{}) (int, error) {
	switch v.(type) {
	case int:
		vInt := v.(int)
		return vInt, nil
	case int8:
		vInt8 := v.(int8)
		return int(vInt8), nil
	case int16:
		vInt16 := v.(int16)
		return int(vInt16), nil
	case int32:
		vInt32 := v.(int32)
		return int(vInt32), nil
	case int64:
		vInt64 := v.(int64)
		return int(vInt64), nil
	case uint:
		vUint := v.(uint)
		return int(vUint), nil
	case uint8:
		vUint8 := v.(uint8)
		return int(vUint8), nil
	case uint16:
		vUint16 := v.(uint16)
		return int(vUint16), nil
	case uint32:
		vUint32 := v.(uint32)
		return int(vUint32), nil
	case uint64:
		vUint64 := v.(uint64)
		return int(vUint64), nil
	default:
		return 0, errors.New("Value cannot be casted to int")
	}
}

// findMediaOfThree returns the median of the three provided values
func findMedianOfThree(v1, v2, v3 int) int {
	// This method relies on the fact that if v1 is the median, then (v2-v1)
	// and (v3-v1) will be 1 positive, 1 negative, and their multiplication
	// is negative. Otherwise, their muliplcation is always positive
	if (v2-v1)*(v3-v1) < 0 {
		return v1
	} else if (v1-v2)*(v3-v2) < 0 {
		return v2
	} else {
		return v3
	}
}

// pickPivot picks the pivot from the list using the specified method and
// returns its index. Returns -1 if the list value is not supported by the
// method
func pickPivot(list []interface{}, method int) int {
	if method&QUICKSORT_PIVOT_FIRST != 0 {
		return 0
	} else if method&QUICKSORT_PIVOT_MIDDLE != 0 {
		return (len(list) - 1) / 2
	} else if method&QUICKSORT_PIVOT_LAST != 0 {
		return len(list) - 1
	} else if method&QUICKSORT_PIVOT_MEDIAN_OF_THREE != 0 {
		vFirst, err := interfaceToInt(list[0])
		if err != nil {
			return -1
		}
		mid := (len(list) - 1) / 2
		vMid, err := interfaceToInt(list[mid])
		if err != nil {
			return -1
		}
		last := len(list) - 1
		vLast, err := interfaceToInt(list[last])
		if err != nil {
			return -1
		}
		m := findMedianOfThree(vFirst, vMid, vLast)
		if m == vFirst {
			return 0
		} else if m == vMid {
			return mid
		} else {
			return last
		}
	}
	return -1
}

func swap(list []interface{}, a, b int) {
	tmp := list[a]
	list[a] = list[b]
	list[b] = tmp
}

func quicksort(
	list []interface{},
	comparator func(interface{}, interface{}) (int, error),
	options int,
) error {
	if len(list) > 1 {
		p, err := partition(list, comparator, options)
		if err != nil {
			return err
		}
		err = quicksort(list[:p], comparator, options)
		if err != nil {
			return err
		}
		err = quicksort(list[p+1:], comparator, options)
		if err != nil {
			return err
		}
	}
	return nil
}

func partition(
	list []interface{},
	comparator func(interface{}, interface{}) (int, error),
	options int,
) (int, error) {
	pIdx := pickPivot(list, options)
	pivot := list[pIdx]
	if pivot == -1 {
		return -1, errors.New("Value in list is not supported by the pivot picking method")
	}

	swap(list, 0, pIdx)
	var result int
	var err error
	i := 0
	for j, l := 1, len(list); j < l; j++ {
		result, err = comparator(list[j], pivot)
		if err != nil {
			return -1, err
		}
		if result < 0 { // list[j] < pivot
			i++
			swap(list, j, i)
		}
	}
	// Adjust the pivot position in list
	i++
	result, err = comparator(list[i], pivot)
	if err != nil {
		return -1, err
	}
	if result < 0 { // list[i] < pivot
		swap(list, i, pIdx)
	}

	return i, nil
}

// Use Quicksort to sort the provided list with the specified options and
// return the resulting list, return error if there is a problem.
// comparator guides the function how to do comparison on the list value. It
// accepts two interface{} value, and should returns positive value if the 1st
// argument > 2nd argument, negative value if 1st argument < 2nd argument and 0
// if they are equal.
// By default, the function sorts directly on the underlying array of the
// slice, you will need to provide a slice with a cloned underlying array if
// this is a problem.
func QuicksortWithOptions(
	list []interface{},
	comparator func(interface{}, interface{}) (int, error),
	options int,
) ([]interface{}, error) {
	// TODO: options to sort on a cloned slice instead of the original slice
	err := quicksort(list, comparator, options)
	if err != nil {
		return nil, err
	}

	return list, nil
}
