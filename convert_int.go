package cdt

// ToIntE function to convert other type data to int, with error
func (i *convert) ToIntE() (int, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return 0, err
	}
	return int(number), nil
}

// ToIntPtr function to convert other type data to int, with pointer
func (i *convert) ToIntPtr() *int {
	v, _ := i.ToIntE()
	return &v
}

// ToIntPtrE function to convert other type data to int, with pointer and error
func (i *convert) ToIntPtrE() (*int, error) {
	v, err := i.ToIntE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToInt function to convert other type data to int
func (i *convert) ToInt() int {
	v, _ := i.ToIntE()
	return v
}

// ToInt8E function converts data of any type to int8, with error
func (i *convert) ToInt8E() (int8, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return 0, err
	}
	return int8(number), err
}

// ToInt8Ptr function to convert other type data to int8, with pointer
func (i *convert) ToInt8Ptr() *int8 {
	v, _ := i.ToInt8E()
	return &v
}

// ToInt8PtrE function to convert other type data to int8, with pointer and error
func (i *convert) ToInt8PtrE() (*int8, error) {
	v, err := i.ToInt8E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToInt8 function to convert other type data to int8
func (i *convert) ToInt8() int8 {
	v, _ := i.ToInt8E()
	return v
}

// ToInt16E function converts data of any type to int16, with error
func (i *convert) ToInt16E() (int16, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return 0, err
	}
	return int16(number), err
}

// ToInt16Ptr function to convert other type data to int16, with pointer
func (i *convert) ToInt16Ptr() *int16 {
	v, _ := i.ToInt16E()
	return &v
}

// ToInt16PtrE function to convert other type data to int16, with pointer and error
func (i *convert) ToInt16PtrE() (*int16, error) {
	v, err := i.ToInt16E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToInt16 function to convert other type data to int16
func (i *convert) ToInt16() int16 {
	v, _ := i.ToInt16E()
	return v
}

// ToInt32E function to convert other type data to int32, with error
func (i *convert) ToInt32E() (int32, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return 0, err
	}
	return int32(number), nil
}

// ToInt32Ptr function to convert other type data to int32, with pointer
func (i *convert) ToInt32Ptr() *int32 {
	v, _ := i.ToInt32E()
	return &v
}

// ToInt32PtrE function to convert other type data to int32, with pointer and error
func (i *convert) ToInt32PtrE() (*int32, error) {
	v, err := i.ToInt32E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToInt32 function to convert other type data to int32
func (i *convert) ToInt32() int32 {
	v, _ := i.ToInt32E()
	return v
}

// ToInt64E function to convert other type data to int64, with error
func (i *convert) ToInt64E() (int64, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return 0, err
	}
	return int64(number), nil
}

// ToInt64Ptr function to convert other type data to int64, with pointer
func (i *convert) ToInt64Ptr() *int64 {
	v, _ := i.ToInt64E()
	return &v
}

// ToInt64PtrE function to convert other type data to int64, with pointer and error
func (i *convert) ToInt64PtrE() (*int64, error) {
	v, err := i.ToInt64E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToInt64 function to convert other type data to int64
func (i *convert) ToInt64() int64 {
	v, _ := i.ToInt64E()
	return v
}
