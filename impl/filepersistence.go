package impl

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/irgndsondepp/cleaningplan/interfaces"
)

type FilePersistence struct {
	filename       string
	conv           interfaces.Converter
	defaultPersons []interfaces.Person
	defaultTasks   []interfaces.Task
}

func NewFilePersistence(file string, c interfaces.Converter, defPers []interfaces.Person, defTasks []interfaces.Task) *FilePersistence {
	return &FilePersistence{
		filename:       file,
		conv:           c,
		defaultPersons: defPers,
		defaultTasks:   defTasks,
	}
}

func (f *FilePersistence) Load(cp interfaces.Plan) {
	ex, err := exists(f.filename)
	if err != nil {
		fmt.Println("Error checking if file exists: %v", err)
	}
	if !ex {
		cp.Init(f.defaultPersons, f.defaultTasks)
		f.Save(cp)
	} else {
		bytes, err := ioutil.ReadFile(f.filename)
		if err != nil {
			cp.Init(f.defaultPersons, f.defaultTasks)
		}
		err = f.conv.ReadFrom(bytes, cp)
		if err != nil {
			fmt.Printf("Error decoding file: %v\n", err)
			cp.Init(f.defaultPersons, f.defaultTasks)
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func (f *FilePersistence) Save(cp interfaces.Plan) {
	bytes, err := f.conv.ConvertTo(cp)
	if err != nil {
		fmt.Printf("Error trying to Encode Plan: %v\n", err)
	}
	err = ioutil.WriteFile(f.filename, bytes, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v", err)
	}
}
