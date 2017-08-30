package collection

import "errors"

type ArrayList struct {
	data []interface{}
}

func New() *ArrayList {
	return new(ArrayList)
}

func (list *ArrayList) Set(index int, v interface{}) error {
	if err := list.rangeCheck(index); err != nil {
		return err
	}
	list.data[index] = v
	return nil
}

func (list *ArrayList) Get(index int) (v interface{}, e error) {
	if err := list.rangeCheck(index); err != nil {
		return nil, err
	}
	return list.data[index], nil
}

func (list *ArrayList) Add(v interface{}) {
	list.data = append(list.data, v)
}

func (list *ArrayList) AddAll(l ArrayList) error {
	list.data = append(list.data, l.data...)
	return nil
}

func (list *ArrayList) Contains(v interface{}) bool {
	for _, value := range list.data {
		if value == v {
			return true
		}
	}
	return false
}

func (list *ArrayList) Clear() {
	list.data = make([]interface{}, 0)
}

func (list *ArrayList) GetSlice() []interface{} {
	return list.data
}

func (list *ArrayList) rangeCheck(index int) error {
	if len(list.data) < index {
		return errors.New("数组角标越界")
	}
	return nil
}
