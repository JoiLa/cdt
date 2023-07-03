package cdt

// ToFloat64E function to convert other type data to float64, with error
func (i *convert) ToFloat64E() (float64, error) {
	return toFloat64(i.GetOriginValRef())
}

// ToFloat64Ptr function to convert other type data to float64, with pointer
func (i *convert) ToFloat64Ptr() *float64 {
	v, _ := i.ToFloat64E()
	return &v
}

// ToFloat64PtrE function to convert other type data to float64, with pointer and error
func (i *convert) ToFloat64PtrE() (*float64, error) {
	v, err := i.ToFloat64E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToFloat64 function to convert other type data to float64
func (i *convert) ToFloat64() float64 {
	v, _ := i.ToFloat64E()
	return v
}

// ToFloat32E function to convert other type data to float32
func (i *convert) ToFloat32E() (float32, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return float32(0), err
	}
	return float32(number), nil
}

// ToFloat32Ptr function to convert other type data to float32, with pointer
func (i *convert) ToFloat32Ptr() *float32 {
	v, _ := i.ToFloat32E()
	return &v
}

// ToFloat32PtrE function to convert other type data to float32, with pointer and error
func (i *convert) ToFloat32PtrE() (*float32, error) {
	v, err := i.ToFloat32E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToFloat32 function to convert other type data to float32
func (i *convert) ToFloat32() float32 {
	v, _ := i.ToFloat32E()
	return v
}
