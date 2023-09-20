package floorAccess

import (
	"fmt"
	"strings"
)

func CheckAccess()  {
		fmt.Println("Введите ID:")
		var id string
		_, err := fmt.Scan(&id)
		if err != nil {
			fmt.Println("Ошибка ID:")
			return
		}
		floor := 0
		fmt.Println("выберите этож от 1 до 3:")
		_, err = fmt.Scan(&floor)
		if err != nil {
			fmt.Println("Ошибка этажа", err.Error())
		}
		fmt.Scan(&floor)
		switch floor {
		case 1,2,3 : 
			fmt.Println("ok")
		default: 
			fmt.Println("нет такого этожа!")
		}
		if strings.HasPrefix(id, "1") {
				if floor > 2 {
					fmt.Println("Нет доступа")
				}
		} else if strings.HasPrefix(id, "2") {
			if floor > 3 {
				fmt.Println("Нет доступа")
			}
		} else {
				fmt.Println("wellcome")
		}
}