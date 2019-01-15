package lab1

var starContext = false

func Match(in, template rune) bool {
	if template == '*' {
		starContext = true
		return true
	}
	if in == template {
		starContext = false
		return true
	}
	if starContext {
		return true
	}
	return false
}

func TemplateSubstr(in, template string) string {
	buffer := make ([]rune, len ([]rune (in)))
	i, j := 0, 0
	templateR := []rune(template)
	maxLen := len(templateR)
	templateEnd := func() bool {
		return j >= maxLen
	}
	push := func(ch rune) {
		if !starContext {
			j++
		} else {
			for !templateEnd() && templateR[j] == '*' {
				j++
			}
		}
		buffer[i++] = ch
	}
	for _, ch := range in {
		if Match(ch, templateR[j]) {
			push(ch)
			if templateEnd() {
				if starContext {
					return in
				}
				return string(buffer[:i])
			}
		} else {
			break
		}
	}
	if templateEnd() {
		return string(buffer[:i])
	}
	return ""
}

func TemplateMatches(in, template string) []string {
	var result []string
	if len(template) == 0 {
		return result
	}
	inR := []rune(in)
	templateR := []rune(template)
	firstChar := templateR[0]
	match := func(ch rune) bool {
		return Match(ch, firstChar)
	}
	substr := func(i int) {
		tmp := TemplateSubstr(string(inR[i:]), template)
		if tmp != "" {
			result = append(result, tmp)
		}
	}
	for i := 0; i <= len(inR)-len(templateR); i++ {
		if match(inR[i]) {
			substr(i)
		}
	}
	return result
}

func Equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i, elem := range x {
		if y[i] != elem {
			return false
		}
	}
	return true
}
