package cdt

// ToMapE function to convert other type data to map, with error
func (i *convert) ToMapE() (map[string]any, error) {
	return toMap(i.GetOriginValRef())
}

// ToMapPtr function to convert other type data to map, with pointer
func (i *convert) ToMapPtr() *map[string]any {
	v, _ := i.ToMapE()
	return &v
}

// ToMapPtrE function to convert other type data to map, with pointer and error
func (i *convert) ToMapPtrE() (*map[string]any, error) {
	v, err := i.ToMapE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToMap function to convert other type data to map
func (i *convert) ToMap() map[string]any {
	v, _ := i.ToMapE()
	return v
}
