package consul

//将bool值转换为指针
func BoolToPtr(v bool) *bool {
	return &v
}

// 将bool指针转换为bool,如果指针为nil，则直接返回false
func PtrToBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func makeBoolPtr(pri *bool, sec bool) *bool {
	if pri != nil {
		return pri
	}
	return &sec
}

//将int值转换为指针
func IntToPtr(v int) *int {
	return &v
}

// 将int指针转换为int,如果指针为nil，则直接返回0
func PtrToInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func makeIntPtr(pri *int, sec int) *int {
	if pri != nil {
		return pri
	}
	return &sec
}
