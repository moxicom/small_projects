package main

import (
	"errors"
	"time"
)

/*
Реализовать in-memory key-value кэш с задаваемым временем жизни у записей.
Кэш должен принимать значения любого типа, ключи - string.
Время жизни должно работать по-настоящему, то-есть удалять записи по истечении заданного времени.
Должна быть возможность как создавать записи, так и получать их по заданному ключу.
При попытке создать запись с уже существующим ключом должна возвращаться ошибка.
Предусмотреть использование кэша несколькими горутинами одновременно. Разделить кэш на интерфейс и реализацию.
*/

type myMap struct {
	Key     string
	Value   int
	TimeOut time.Time
}

type storage struct {
	data []myMap
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) Insert(key string, value int, timeout time.Time) error {
	input := myMap{
		Key:     key,
		Value:   value,
		TimeOut: timeout,
	}

	if _, err := s.Get(input.Key); err == nil {
		return errors.New("this ket is already exists")
	}
	s.data = append(s.data, input)
	return nil
}

func (s *storage) Get(key string) (myMap, error) {
	var mp myMap
	var noDataText = "no data"
	var idx int
	var found = false

	for i, el := range s.data {
		if el.Key == key {
			mp = el
			idx = i
			found = true
			break
		}
	}

	if !found {
		return myMap{}, errors.New(noDataText)
	} else if mp.TimeOut.Before(time.Now()) {
		s.deleteById(idx)
		return myMap{}, errors.New(noDataText)
	}
	return mp, nil
}

func (s *storage) Delete(key string) {
	for i, el := range s.data {
		if el.Key == key {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return
		}
	}
}

func (s *storage) deleteById(i int) {
	if i < 0 || len(s.data) >= i {
		return
	}
	s.data = append(s.data[:i], s.data[i+1:]...)
}

func main() {

}
