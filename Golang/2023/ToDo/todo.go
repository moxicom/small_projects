package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Task struct {
	Text         string
	IsDone       bool
	CreatedTime  time.Time
	CompleteTime time.Time
}

type Tasks []Task

func (t *Tasks) OpenTasks(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil

}

func (t *Tasks) UpdateJson(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Tasks) Add(s string) {
	todo := Task{
		Text:         s,
		IsDone:       false,
		CreatedTime:  time.Now(),
		CompleteTime: time.Now(),
	}

	*t = append(*t, todo)
}

func (t *Tasks) Remove(idx int) error {
	lst := *t
	if idx <= 0 || idx > len(lst) {
		return errors.New("Incorrect index")
	}

	*t = append(lst[:idx-1], lst[idx:]...)

	return nil
}

func (t *Tasks) Complete(idx int) error {
	lst := *t
	if idx <= 0 || idx > len(lst) {
		return errors.New("Incorrect index")
	}
	lst[idx-1].IsDone = true
	lst[idx-1].CompleteTime = time.Now()
	return nil
}
