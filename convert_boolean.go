package cdt

// ToBooleanE function to convert other type data to boolean, with error
func (i *convert) ToBooleanE() (bool, error) {
	return toBoolean(i.GetOriginValRef())
}

// ToBooleanPtr function to convert other type data to boolean, with pointer
func (i *convert) ToBooleanPtr() *bool {
	v, _ := i.ToBooleanE()
	return &v
}

// ToBooleanPtrE function to convert other type data to boolean, with pointer and error
func (i *convert) ToBooleanPtrE() (*bool, error) {
	v, err := i.ToBooleanE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToBoolean function to convert other type data to boolean
func (i *convert) ToBoolean() bool {
	v, _ := i.ToBooleanE()
	return v
}
