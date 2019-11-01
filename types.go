package godash

// IntPtr will convert a int to an *int
func IntPtr(i int) *int {
	p := int(i)
	return &p
}

// Int8Ptr will convert a int8 to an *int8
func Int8Ptr(i int8) *int8 {
	p := int8(i)
	return &p
}

// Int16Ptr will convert a int16 to an *int16
func Int16Ptr(i int16) *int16 {
	p := int16(i)
	return &p
}

// Int32Ptr will convert a int32 to an *int32
func Int32Ptr(i int32) *int32 {
	p := int32(i)
	return &p
}

// Int64Ptr will convert a int64 to an *int64
func Int64Ptr(i int64) *int64 {
	p := int64(i)
	return &p
}
