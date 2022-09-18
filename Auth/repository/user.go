package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/salimkun/Efishery-Test/Auth/model"
)

func CreateUser(user *model.User) error {
	getAllfile, err := ReadFile()
	if err != nil {
		return err
	}
	getAllfile = append(getAllfile, user)
	file, _ := json.MarshalIndent(getAllfile, "", " ")

	err = ioutil.WriteFile("user.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile() ([]*model.User, error) {
	users := []*model.User{}
	jsonFile, err := os.Open("user.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	return users, nil
}
