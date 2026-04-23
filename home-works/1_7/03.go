package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var text = "У лукоморья дуб зелёный;\nЗлатая цепь на дубе том:\nИ днём и ночью кот учёный\nВсё ходит по цепи кругом;" +
	"\nИдёт направо — песнь заводит,\nНалево — сказку говорит.\n\nТам чудеса: там леший бродит," +
	"\nРусалка на ветвях сидит;\nТам на неведомых дорожках\nСледы невиданных зверей;\nИзбушка там на курьих ножках" +
	"\nСтоит без окон, без дверей;\n\nТам лес и дол видений полны;\nТам о заре прихлынут волны" +
	"\nНа брег песчаный и пустой,\nИ тридцать витязей прекрасных\nЧредой из вод выходят ясных," +
	"\nИ с ними дядька их морской;"
var sliceText []string
var lowerWords []string
var resMap = make(map[string]int)

func main() {
	setWords()
	toLowerWords()
	countWords()
	printRes()
}

func setWords() {
	var re = regexp.MustCompile(`(?m)[а-яА-ЯёЁ\-]+`)
	sliceText = re.FindAllString(strings.TrimSpace(text), -1)
}

func toLowerWords() {
	lowerWords = lowerWords[:0]
	for _, word := range sliceText {
		char := []rune(word)
		for i := range char {
			char[i] = unicode.ToLower(char[i])
		}
		lowerWords = append(lowerWords, string(char))
	}
}

func countWords() {
	for _, count := range lowerWords {
		resMap[count]++
	}
}

func printRes() {
	str := ""
	for _, word := range getUnique() {
		str += fmt.Sprintf("%s - %d\n", word, resMap[word])
	}
	fmt.Printf(str)
}

func getUnique() []string {
	found := make(map[string]bool)
	var uniq []string
	sort.Strings(lowerWords)
	for _, str := range lowerWords {
		if !found[str] {
			found[str] = true
			uniq = append(uniq, str)
		}
	}
	return uniq
}
