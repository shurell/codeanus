package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var BS string
var xx = 0

func main() {

	functionsArray := make(map[string]int)
	proceduresArray := make(map[string]int)

	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	BS = string(bs)
	//fmt.Println(BS)
	fmt.Println("==================================================================")
	fmt.Println("========================   N  E  X  T   ==========================")
	fmt.Println("==================================================================")

	//убераем все табы и двойные пробелы

	for strings.Contains(BS, "\t") {
		xx += 1
		BS = strings.Replace(BS, "\t", " ", -1)
		fmt.Println("ITERATIONS:", xx)
	}

	for strings.Contains(BS, "  ") {
		xx += 1
		BS = strings.Replace(BS, "  ", " ", -1)
		fmt.Println("ITERATIONS:", xx)
	}

	//fmt.Println(BS)

	fmt.Println("==================================================================")
	fmt.Println("========================   F I N I S H  ==========================")
	fmt.Println("==================================================================")

	//разбиваем полотнище на массив строк
	strArray := strings.Split(BS, "\n")

	//Ищем функцию для внесения в справочник

	for i, mystr := range strArray {
		if strings.Contains(mystr, "Функция ") && !strings.Contains(mystr, "Функция =") && !strings.HasPrefix(mystr, "//") {
			if isSingle(mystr, "Функция", i) {
				tmpStrings := strings.Split(mystr, "(")
				tmpName := strings.Replace(tmpStrings[0], "Функция ", "", 1)
				fmt.Println("Найдена функция:", tmpName, "в строке", i+1)
				functionsArray[tmpName] = i + 1
			}
		}
	}

	//Ищем процедуру для внесения в справочник

	for i, mystr := range strArray {
		if strings.Contains(mystr, "Процедура ") && !strings.HasPrefix(mystr, "//") {
			if isSingle(mystr, "Процедура", i) {
				tmpStrings := strings.Split(mystr, "(")
				tmpName := strings.Replace(tmpStrings[0], "Процедура ", "", 1)
				fmt.Println("Найдена процедура:", tmpName, "в строке", i+1)
				proceduresArray[tmpName] = i + 1
			}
		}
	}

	//Ищем любые вызовы

	/* 	for i, mystr := range strArray {
		if !strings.Contains(mystr, "Процедура ") && !strings.Contains(mystr, "Функция ") && strings.Contains(mystr, "(") && !strings.HasPrefix(mystr, "//") {

			//возможно встречается вызов
					fmt.Println("Возможно, строка", i+1, "содержит вызов:", mystr)
		}

	} */

	fmt.Println("==================================================================")
	fmt.Println("========================      TREE      ==========================")
	fmt.Println("==================================================================")
	treeBuild("1c/CommonModules")

}

func treeBuild(str string) {
	// строим дерево по выбранному пути

	filepath.Walk(str, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println("ПАПКА:", path)
		} else {
			fmt.Println(path)
		}
		return nil
	})

}

func isSingle(s1, s2 string, i int) bool {
	st2 := strings.Split(s1, s2)
	if st2[0] != "" {
		fmt.Println("ПодСтрока грязная в строке", i+1)
		return false
	}
	return true
}
