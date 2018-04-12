package util

// CheckError 检查一个错误,确定错误panic程序,否则什么也不做
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
