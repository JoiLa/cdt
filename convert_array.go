package cdt

// ToArrayE function to convert other type data to array, with error
func (i *convert) ToArrayE() ([]any, error) {
	return toArray(i.GetOriginValRef())
}

// ToArrayPtr function to convert other type data to array, with pointer
func (i *convert) ToArrayPtr() *[]any {
	v, _ := i.ToArrayE()
	return &v
}

// ToArrayPtrE function to convert other type data to array, with pointer and error
func (i *convert) ToArrayPtrE() (*[]any, error) {
	v, err := i.ToArrayE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToArray function to convert other type data to array
func (i *convert) ToArray() []any {
	v, _ := i.ToArrayE()
	return v
}
