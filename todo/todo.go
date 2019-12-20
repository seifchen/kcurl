package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Url  string
	Name string
	Env  string
}

func SaveItem(filename string, itmes []Item) error {
	b, err := json.Marshal(itmes)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return items, err
	}
	return items, nil
}

func (i *Item) SetEnv(env string) {
	switch env {
	case "online":
		i.Env = "online"
	default:
		i.Env = "dev"

	}
}

func (i *Item) SetName(name string) {
	i.Name = name
}
