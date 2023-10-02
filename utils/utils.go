package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func randomGenerateInt(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func randomGenerateIntWithExclusions(min, max int, exclusions ...int) int {
	rand.NewSource(time.Now().UnixNano())
	num := min + rand.Intn(max-min+1)
	for contains(exclusions, num) {
		num = min + rand.Intn(max-min+1)
	}
	return num
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func PhoneNumber() string {
	var fullNumber string
	const code string = "+7"
	partOne := randomGenerateInt(900, 998)
	partTwo := randomGenerateInt(100, 899)
	partThree := randomGenerateInt(10, 99)
	partFour := randomGenerateInt(10, 99)
	fullNumber = code + " " + strconv.Itoa(partOne) + " " + strconv.Itoa(partTwo) + " " + strconv.Itoa(partThree) + " " + strconv.Itoa(partFour)
	//fmt.Println(fullNumber)
	return fullNumber
}

func Pass() (string, string) {
	var digits string
	var region string
	excludedNumbers := []int{2, 6, 9, 13, 16, 31, 39, 48, 51, 55, 59, 62, 72}
	partOne := randomGenerateIntWithExclusions(1, 99, excludedNumbers...)
	partTwo := randomGenerateInt(01, 23)
	parthThree := randomGenerateInt(000101, 999999)

	partOneStr := strconv.Itoa(partOne)
	partTwoStr := strconv.Itoa(partTwo)
	partThreeStr := strconv.Itoa(parthThree)

	if partOne >= 1 && partOne <= 9 {
		digits += "0" + partOneStr
	} else {
		digits += partOneStr
	}
	if partTwo >= 1 && partTwo <= 9 {
		digits += " 0" + partTwoStr + " "
	} else {
		digits += " " + partTwoStr + " "
	}

	switch {
	case parthThree >= 101 && parthThree <= 999:
		digits += "000" + partThreeStr
	case parthThree >= 1000 && parthThree <= 9999:
		digits += "00" + partThreeStr
	case parthThree >= 10000 && parthThree <= 99999:
		digits += "0" + partThreeStr
	default:
		digits += "" + partThreeStr
	}

	db, err := sql.Open("sqlite3", "cities.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return "", ""
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	query := fmt.Sprintf("SELECT subject FROM my_table WHERE okato = %d", partOne)
	err = db.QueryRow(query).Scan(&region)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", ""
	}
	return digits, region
}

func Card() (string, string) {
	var num string
	var cvv string
	cvv = strconv.Itoa(randomGenerateInt(100, 999))
	firstPart := strconv.Itoa(randomGenerateInt(2, 6))
	secondPart := strconv.Itoa(randomGenerateInt(100, 999))
	thirdPart := strconv.Itoa(randomGenerateInt(1000, 9999))
	fourthPart := strconv.Itoa(randomGenerateInt(1000, 9999))
	fifthPart := strconv.Itoa(randomGenerateInt(1000, 9999))
	num = firstPart + secondPart + " " + thirdPart + " " + fourthPart + " " + fifthPart
	return num, cvv
}

func ParseTxt(path string) []string {

	//читаем содержимое
	content, err := os.ReadFile(path)
	if err != nil {
		return []string{"Ошибка чтения"}
	}

	//преорбразуем
	dock := strings.Split(string(content), "\n")
	//for _, line := range dock {
	//	fmt.Println(line)
	//}

	var result []string
	for _, line := range dock {
		if trimmedLine := strings.TrimSpace(line); trimmedLine != "" {
			result = append(result, trimmedLine)
		}
	}

	return result
}

func GetWord(fio []string) string {
	// Создание генератора случайных чисел с использованием текущего времени как источника
	rand.NewSource(time.Now().Unix())
	// Получение случайного индекса из массива
	randomIndex := rand.Intn(len(fio))
	// Извлечение случайного слова из массива
	return fio[randomIndex]
}

func CheckDB() {
	//const filePath = "C:\\Users\\Alex\\GolandProjects\\mprj1"
	if _, err := os.Stat("cities.db"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Database does not exist, creating...")
			DBParse()
		}
	} else {
		fmt.Println("Database exist")
	}
	return
}

func DBParse() {
	db, err := sql.Open("sqlite3", "./cities.db")
	if err != nil {
		fmt.Println("Ошибка открытия БД:", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Создание таблицы
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS my_table (
						subject TEXT,
						okato INTEGER
					)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Открытие CSV-файла
	file, err := os.Open("assets/cities.csv")
	if err != nil {
		fmt.Println("Ошибка открытия CSV-файла:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Чтение CSV-файла
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Ошибка чтения CSV:", err)
		return
	}

	// Вставка данных в таблицу
	for _, row := range data {
		subject := row[0]
		okato, err := strconv.Atoi(row[1])
		if err != nil {
			fmt.Println("Ошибка конвертации 3-го столбца в целочисленный тип:", err)
			return
		}

		_, err = db.Exec("INSERT INTO my_table (subject, okato) VALUES (?, ?)", subject, okato)
		if err != nil {
			fmt.Println("Ошибка вставки данных:", err)
			return
		}
	}

	fmt.Println("Данные были успешно вставлены в базу данных.")
}

func PrintDB() {
	db, err := sql.Open("sqlite3", "./cities.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Выполнение запроса на выборку данных
	rows, err := db.Query("SELECT subject, okato FROM my_table")
	if err != nil {
		fmt.Println("Error querying data:", err)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// Обработка результатов
	for rows.Next() {
		var subject string
		var okato int
		err := rows.Scan(&subject, &okato)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("Subject: %s, OKATO: %d\n", subject, okato)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}

func Handler(gender int) (string, error) {
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
		surnameMale := ParseTxt("assets/familii_m.txt")
		nameMale := ParseTxt("assets/imena_m.txt")
		patronymicMale := ParseTxt("assets/otchestva_m.txt")
		phoneNum = PhoneNumber()
		pass, region = Pass()
		numCard, cvv = Card()
		surname = GetWord(surnameMale)
		name = GetWord(nameMale)
		patronymic = GetWord(patronymicMale)
	case 8:
		surnameFemale := ParseTxt("assets/familii_zh.txt")
		nameFemale := ParseTxt("assets/imena_zh.txt")
		patronymicFemale := ParseTxt("assets/otchestva_zh.txt")
		phoneNum = PhoneNumber()
		pass, region = Pass()
		numCard, cvv = Card()
		surname = GetWord(surnameFemale)
		name = GetWord(nameFemale)
		patronymic = GetWord(patronymicFemale)
	}

	resultStr := fmt.Sprintf("\nФИО: %s %s %s\nНомер телефона: %s\nСерия, номер паспорта: %s\nДанные карты(номер, CVV): %s | %s\nСубъект РФ: %s",
		surname, name, patronymic, phoneNum, pass, numCard, cvv, region)

	return resultStr, nil
}

func Startapp() {
	for {
		fmt.Println("\nВыберите пол (7 - мужской, 8 - женский, 56 - выход):")
		var gender int
		_, err := fmt.Scanln(&gender)
		if err != nil {
			fmt.Println("Ошибка при вводе:", err)
			continue
		}

		result, err := Handler(gender)
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
