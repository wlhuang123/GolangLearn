package tool

import "reflect"

// IsInArr 判断value是否在arr中
func IsInArr(arr interface{}, value interface{}) bool {
	array, ok := CreateArray(arr)
	if !ok {
		return false
	}

	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}

// CreateArray interface{}转为 []interface{}
// 如果输入不是数组，返回not ok
func CreateArray(i interface{}) ([]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() != reflect.Slice {
		return nil, false
	}

	var array []interface{}
	for i := 0; i < value.Len(); i++ {
		array = append(array, value.Index(i).Interface())
	}
	return array, true
}

// IsTowSliceEqual 判断两个切片有相同的元素,不考虑顺序
func IsTowSliceEqual(arr1 interface{}, arr2 interface{}) bool {
	a1, _ := CreateArray(arr1)
	a2, _ := CreateArray(arr2)
	if len(a1) != len(a2) {
		return false
	}

	for i := 0; i < len(a1); i++ {
		if !IsInArr(a2, a1[i]) {
			return false
		}
	}

	for i := 0; i < len(a2); i++ {
		if !IsInArr(a1, a2[i]) {
			return false
		}
	}
	return true
}

// IsTowSliceTheSame 判断两个切片数组元素一模一样,包括位置
func IsTowSliceTheSame(arr1 interface{}, arr2 interface{}) bool {
	a1, _ := CreateArray(arr1)
	a2, _ := CreateArray(arr2)
	if len(a1) != len(a2) {
		return false
	}

	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

// OrderChange 降序转为升序
func OrderChange(nums []int) {
	for i := 0; i < len(nums)-i; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
