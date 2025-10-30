package actioninfo

import (
	"fmt"
	"log"
)

//import "github.com/gostaticanalysis/nilerr"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("error parse in string - %q: %s\n", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error receiving information about training from string - %q: %s\n", data, err)
			continue
		}

		fmt.Println(info)
	}
}
