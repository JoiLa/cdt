package cdt

// ToUintE function to convert interface to uint, with error
func (i *convert) ToUintE() (uint, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return uint(0), err
	}
	return uint(number), nil
}

// ToUintPtr function to convert interface to uint, with pointer
func (i *convert) ToUintPtr() *uint {
	v, _ := i.ToUintE()
	return &v
}

// ToUintPtrE function to convert interface to uint, with pointer and error
func (i *convert) ToUintPtrE() (*uint, error) {
	v, err := i.ToUintE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToUint function to convert interface to uint
func (i *convert) ToUint() uint {
	v, _ := i.ToUintE()
	return v
}

// ToUint8E function to convert interface data to uint8, with error
func (i *convert) ToUint8E() (uint8, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return uint8(0), err
	}
	return uint8(number), nil
}

// ToUint8Ptr function to convert interface to uint8, with pointer
func (i *convert) ToUint8Ptr() *uint8 {
	v, _ := i.ToUint8E()
	return &v
}

// ToUint8PtrE function to convert interface to uint8, with pointer and error
func (i *convert) ToUint8PtrE() (*uint8, error) {
	v, err := i.ToUint8E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToUint8 function to convert interface to uint8
func (i *convert) ToUint8() uint8 {
	v, _ := i.ToUint8E()
	return v
}

// ToUint16E function to convert interface data to uint16, with error
func (i *convert) ToUint16E() (uint16, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return uint16(0), err
	}
	return uint16(number), nil
}

// ToUint16Ptr function to convert interface to uint16, with pointer
func (i *convert) ToUint16Ptr() *uint16 {
	v, _ := i.ToUint16E()
	return &v
}

// ToUint16PtrE function to convert interface to uint16, with pointer and error
func (i *convert) ToUint16PtrE() (*uint16, error) {
	v, err := i.ToUint16E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToUint16 function to convert interface to uint16
func (i *convert) ToUint16() uint16 {
	v, _ := i.ToUint16E()
	return v
}

// ToUint32E function to convert interface data to uint32, with error
func (i *convert) ToUint32E() (uint32, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return uint32(0), err
	}
	return uint32(number), nil
}

// ToUint32Ptr function to convert interface to uint32, with pointer
func (i *convert) ToUint32Ptr() *uint32 {
	v, _ := i.ToUint32E()
	return &v
}

// ToUint32PtrE function to convert interface to uint32, with pointer and error
func (i *convert) ToUint32PtrE() (*uint32, error) {
	v, err := i.ToUint32E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToUint32 function to convert interface to uint32
func (i *convert) ToUint32() uint32 {
	v, _ := i.ToUint32E()
	return v
}

// ToUint64E function to convert interface data to uint64, with error
func (i *convert) ToUint64E() (uint64, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return uint64(0), err
	}
	return uint64(number), nil
}

// ToUint64Ptr function to convert interface to uint64, with pointer
func (i *convert) ToUint64Ptr() *uint64 {
	v, _ := i.ToUint64E()
	return &v
}

// ToUint64PtrE function to convert interface to uint64, with pointer and error
func (i *convert) ToUint64PtrE() (*uint64, error) {
	v, err := i.ToUint64E()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToUint64 function to convert interface to uint64
func (i *convert) ToUint64() uint64 {
	v, _ := i.ToUint64E()
	return v
}

// ToUintptrE function to convert interface data to uintptr, with error
func (i *convert) ToUintptrE() (uintptr, error) {
	number, err := i.ToFloat64E()
	if err != nil {
		return uintptr(0), err
	}
	return uintptr(number), nil
}

// ToUintptrPtr function to convert interface to uintptr, with pointer
func (i *convert) ToUintptrPtr() *uintptr {
	v, _ := i.ToUintptrE()
	return &v
}

// ToUintptrPtrE function to convert interface to uintptr, with pointer and error
func (i *convert) ToUintptrPtrE() (*uintptr, error) {
	v, err := i.ToUintptrE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToUintptr function to convert interface to uintptr
func (i *convert) ToUintptr() uintptr {
	v, _ := i.ToUintptrE()
	return v
}
