package cdt

import (
	"time"
)

// ToTimeE function to convert other type data to time, with error
func (i *convert) ToTimeE() (time.Time, error) {
	return toTime(i.GetOriginValRef())
}

// ToTimePtr function to convert other type data to time, with pointer
func (i *convert) ToTimePtr() *time.Time {
	v, _ := i.ToTimeE()
	return &v
}

// ToTimePtrE function to convert other type data to time, with pointer and error
func (i *convert) ToTimePtrE() (*time.Time, error) {
	v, err := i.ToTimeE()
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ToTime function to convert other type data to time
func (i *convert) ToTime() time.Time {
	v, _ := i.ToTimeE()
	return v
}
