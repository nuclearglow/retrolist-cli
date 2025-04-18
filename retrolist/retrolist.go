package retrolist

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RetroList struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Items     []*Item   `json:"items"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewRetroList(title, subtitle string) *RetroList {
	return &RetroList{
		UUID:      uuid.New().String(),
		Title:     title,
		Subtitle:  subtitle,
		Items:     make([]*Item, 0),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func (r *RetroList) AddItem(item *Item) {
	r.Items = append(r.Items, item)
}

func (r *RetroList) RemoveItem(uuid string) {
	removeIndex := -1
	for i, item := range r.Items {
		if item.UUID == uuid {
			removeIndex = i
			break
		}
	}
	if removeIndex > -1 {
		r.Items = slices.Delete(r.Items, removeIndex, removeIndex+1)
	}
}

func (r *RetroList) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create RetroList at %s", path)
	}

	defer file.Close()

	var writer strings.Builder

	encoder := json.NewEncoder(&writer)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(r)
	if err != nil {
		return err
	}

	fmt.Fprintln(file, writer.String())

	return nil
}

func (r RetroList) Load(path string) (*RetroList, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("RetroList does not exist at %s\n", path)
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("could not open RetroList at %s\n", path)
		return nil, err
	}

	defer file.Close()

	var list *RetroList

	decoder := json.NewDecoder(file)
	err = decoder.Decode(list)
	if err != nil {
		fmt.Printf("could not decode RetroList at %s\n", path)
		return nil, err
	}

	return list, nil
}
