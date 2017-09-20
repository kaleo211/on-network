package switch_operations

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
	"strings"
	"fmt"
	"os"
	"log"
)

type Switch interface {
	Update(string, string) error
}


type SwitchType struct {
	Switch []struct {
		Name string `yaml: "name"`
		Models []struct {
			Name string `yaml: "name"`
			Disruptive bool `yaml :"disruptive"`
		} `yaml :"models"`

	} `yaml: "switch"`
}

func ReadYaml(fileName string) (*SwitchType, error) {
	switches := SwitchType{}

	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(fileData), &switches)
	if err != nil {
		return nil, err
	}
	return &switches, nil
}


func GetUpdateType( switchType , switchModel string) (string , error) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	switches, err := ReadYaml(os.Getenv("SWITCH_MODELS_FILE_PATH"))

	if err != nil{
		return "", err
	}
	for _, stype :=  range switches.Switch {

		if stype.Name == switchType {

			for _ , smodels := range stype.Models {

				if strings.Contains(strings.ToLower(switchModel), smodels.Name)  {

					if smodels.Disruptive == true {
						return "Disruptive", nil
					} else {
						return "NonDisruptive", nil
					}
				}
			}
			return "", errors.New("couldn't find switch model")
		}
		return "", errors.New("couldn't find switch type")
	}
	return "", errors.New("couldn't find switch type")
}