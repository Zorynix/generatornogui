package main

import (
	"fmt"
	// "image/color"
	"utils/utils"
	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/widget"
	_ "golang.org/x/sys/windows"
)

func handler(gender int) (string, error) {
	//создание каналов
	var phoneNum string
	var pass string
	var region string
	var numCard string
	var cvv string
	var surname string
	var name string
	var patronymic string

	//запись данных
	switch gender {
	case 7:
		surnameMale := utils.ParseTxt("assets/familii_m.txt")
		nameMale := utils.ParseTxt("assets/imena_m.txt")
		patronymicMale := utils.ParseTxt("assets/otchestva_m.txt")
		phoneNum = utils.PhoneNumber()
		pass, region = utils.Pass()
		numCard, cvv = utils.Card()
		surname = utils.GetWord(surnameMale)
		name = utils.GetWord(nameMale)
		patronymic = utils.GetWord(patronymicMale)
	case 8:
		surnameFemale := utils.ParseTxt("assets/familii_zh.txt")
		nameFemale := utils.ParseTxt("assets/imena_zh.txt")
		patronymicFemale := utils.ParseTxt("assets/otchestva_zh.txt")
		phoneNum = utils.PhoneNumber()
		pass, region = utils.Pass()
		numCard, cvv = utils.Card()
		surname = utils.GetWord(surnameFemale)
		name = utils.GetWord(nameFemale)
		patronymic = utils.GetWord(patronymicFemale)

	case 56:
		fmt.Println("Выход...")
		return "suka", nil
	default:
		fmt.Println("Некорректный выбор. Выберите 7,8 или 56\n")
		return "blyat", nil
	}

	resultStr := fmt.Sprintf("\nФИО: %s %s %s\nНомер телефона: %s\nСерия, номер паспорта: %s\nДанные карты(номер, CVV): %s | %s\nСубъект РФ: %s",
		surname, name, patronymic, phoneNum, pass, numCard, cvv, region)

	return resultStr, nil
}

func main() {
	for {
		fmt.Println("\nВыберите пол (7 - мужской, 8 - женский, 56 - выход):")
		var gender int
		_, err := fmt.Scanln(&gender)
		if err != nil {
			fmt.Println("Ошибка при вводе:", err)
			continue
		}

		result, err := handler(gender)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else if result == "suka" {
			break // Выход из цикла при выборе 56
		} else if result == "blyat" {
			continue // Продолжаем цикл при некорректном выборе
		}

		fmt.Println(result)
	}
}
