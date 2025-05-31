package translit

import (
	"strings"
)

// Таблица транслитерации
var ruToEn = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d",
	'е': "e", 'ё': "yo", 'ж': "zh", 'з': "z", 'и': "i",
	'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n",
	'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t",
	'у': "u", 'ф': "f", 'х': "kh", 'ц': "ts", 'ч': "ch",
	'ш': "sh", 'щ': "shch", 'ы': "y", 'э': "e", 'ю': "yu", 'я': "ya",
	'ь': "", 'ъ': "",
}

// ToLatin — простая транслитерация русского текста в латиницу
func ToLatin(input string) string {
	var result strings.Builder
	for _, r := range strings.ToLower(input) {
		if val, ok := ruToEn[r]; ok {
			result.WriteString(val)
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
