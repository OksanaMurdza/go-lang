package lab1

import "testing"

func TestMatch(t *testing.T) {
	tables := []struct {
		x        rune
		y        rune
		expected bool
	}{
		{'a', 'a', true},
		{'a', 'b', false},
		{'b', 'a', false},
		{'*', 'a', false},
		{'a', '*', true},
		{'*', '*', true},
		{'ы', 'ы', true},
	}

	for _, table := range tables {
		result := Match(table.x, table.y)
		if result != table.expected {
			t.Errorf("Match failed for (%c, %c). Expected %t, got %t",
				table.x, table.y, table.expected, result)
		}
	}
}

func TestTemplateSubstr(t *testing.T) {
	tables := []struct {
		x        string
		y        string
		expected string
	}{
		{"", "", ""},
		{"", "a", ""},
		{"b", "ba", ""},
		{"", "*", ""},
		{"ы", "ы", "ы"},
		{"Їсаfr", "Їсаf", "Їсаf"},
		{"Їсаfr", "Ї*f", "Їсаf"},
		{"aaa", "a", "a"},
		{"aaa", "b", ""},
		{"abca", "abc", "abc"},
		{"aaa", "*", "aaa"},
		{"abc", "*", "abc"},
		{"abcada", "*a", "abca"},
		{"abcada", "*a*", "abcada"},
		{"abcada:;**b", "*a*b", "abcada:;**b"},
		{"abcada:;**bcdf", "*a***bcd", "abcada:;**bcd"},
	}

	for _, table := range tables {
		result := TemplateSubstr(table.x, table.y)
		if result != table.expected {
			t.Errorf("Match failed for (%s, %s). Expected %s, got %s",
				table.x, table.y, table.expected, result)
		}
	}
}

func TestTemplateMatches(t *testing.T) {
	tables := []struct {
		x        string
		y        string
		expected []string
	}{
		{"", "", []string{}},
		{"", "a", []string{}},
		{"b", "ba", []string{}},
		{"", "*", []string{}},
		{"aaa", "a", []string{"a", "a", "a"}},
		{"aaa", "b", []string{}},
		{"efabca", "abc", []string{"abc"}},
		{"aaa", "*", []string{"aaa", "aa", "a"}},
		{"abc", "*", []string{"abc", "bc", "c"}},
		{"abcada", "*a", []string{"abca", "bca", "ca", "ada", "da"}},
		{"abcada", "*a*", []string{"abcada", "bcada", "cada"}},
		{"abcada:;**beff", "cad*;*bef", []string{"cada:;**bef"}},
		{"cafaeceerfcccafcf", "c*f", []string{"caf", "ceerf", "cccaf", "ccaf", "caf"}},
	}

	for _, table := range tables {
		result := TemplateMatches(table.x, table.y)
		if !Equal(result, table.expected) {
			t.Errorf("Match failed for (%s, %s). Expected %s, got %s",
				table.x, table.y, table.expected, result)
		}
	}
}

// to grant 100% coverage
func TestEqual(t *testing.T) {
	tables := []struct {
		x        []string
		y        []string
		expected bool
	}{
		{[]string{}, []string{}, true},
		{[]string{}, []string{""}, false},
		{[]string{""}, []string{""}, true},
		{[]string{""}, []string{"a"}, false},
		{[]string{"a"}, []string{"a", "a"}, false},
		{[]string{"b"}, []string{"a"}, false},
		{[]string{"", "", ""}, []string{"", "", ""}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "e"}, false},
	}

	for _, table := range tables {
		result := Equal(table.x, table.y)
		if result != table.expected {
			t.Errorf("Match failed for (%s, %s). Expected %t, got %t",
				table.x, table.y, table.expected, result)
		}
	}
}
