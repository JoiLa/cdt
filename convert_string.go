package cdt

// ToStringE function to convert other type data to string, with error
func (i *convert) ToStringE() (string, error) {
	return toString(i.GetOriginValRef())
}

// ToStringPtr function to convert other type data to string, with pointer
func (i *convert) ToStringPtr() *string {
	v, _ := i.ToStringE()
	return &v
}

// ToStringPtrE function to convert other type data to string, with pointer and error
func (i *convert) ToStringPtrE() (*string, error) {
	v, err := i.ToStringE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToString function to convert other type data to string
func (i *convert) ToString() string {
	v, _ := i.ToStringE()
	return v
}
