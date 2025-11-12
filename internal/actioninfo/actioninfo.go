package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, row := range dataset {
		if err := dp.Parse(row); err != nil {
			log.Printf("parse error: %v (row=%q)", err, row)
			continue
		}
		out, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action info error: %v (row=%q)", err, row)
			continue
		}
		fmt.Println(out)
	}
}
