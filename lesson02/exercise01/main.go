package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Робота із рядками. Виконання популярних пошуків і конкатенація
(всі функції пакета strings за посиланням
https://pkg.go.dev/strings
до "type Builder")
*/

func main() {
	fmt.Println("Операції з строками\n")

	fmt.Println("Функція Compare:")
	fmt.Printf("Compare(\"a\", \"b\") res: %d\n", strings.Compare("a", "b"))
	fmt.Printf("Compare(\"a\", \"a\") res: %d\n", strings.Compare("a", "a"))
	fmt.Printf("Compare(\"b\", \"a\") res: %d\n\n", strings.Compare("b", "a"))

	fmt.Println("Функція Contains:")
	fmt.Printf("Contains(\"seafood\", \"foo\") res: %t\n", strings.Contains("seafood", "foo"))
	fmt.Printf("Contains(\"seafood\", \"bar\") res: %t\n", strings.Contains("seafood", "bar"))
	fmt.Printf("Contains(\"seafood\", \"\") res: %t\n", strings.Contains("seafood", ""))
	fmt.Printf("Contains(\"\", \"\") res: %t\n\n", strings.Contains("", ""))

	fmt.Println("Функція ContainsAny:")
	fmt.Printf("ContainsAny(\"team\", \"i\") res: %t\n", strings.ContainsAny("team", "i"))
	fmt.Printf("ContainsAny(\"fail\", \"ui\") res: %t\n", strings.ContainsAny("fail", "ui"))
	fmt.Printf("ContainsAny(\"ure\", \"ui\") res: %t\n", strings.ContainsAny("ure", "ui"))
	fmt.Printf("ContainsAny(\"failure\", \"ui\") res: %t\n", strings.ContainsAny("failure", "ui"))
	fmt.Printf("ContainsAny(\"foo\", \"\") res: %t\n", strings.ContainsAny("foo", ""))
	fmt.Printf("ContainsAny(\"\", \"\") res: %t\n\n", strings.ContainsAny("", ""))

	fmt.Println("Функція ContainsRune:") //lowercase letter "a", for example, is 97
	fmt.Printf("ContainsRune(\"aardvark\", 97) res: %t\n", strings.ContainsRune("aardvark", 97))
	fmt.Printf("ContainsRune(\"timeout\", 97) res: %t\n\n", strings.ContainsRune("timeout", 97))

	fmt.Println("Функція Count:")
	fmt.Printf("Count(\"cheese\", \"e\") res: %d\n", strings.Count("cheese", "e"))
	fmt.Printf("Count(\"five\", \"\") res: %d\n\n", strings.Count("five", "")) // before & after each rune

	fmt.Println("Функція Cut:")
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) res: %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")

	fmt.Println("")
	fmt.Println("Функція EqualFold:")

	fmt.Printf("EqualFold(\"Go\", \"go\") res: %t\n", strings.EqualFold("Go", "go"))
	fmt.Printf("EqualFold(\"AB\", \"ab\") res: %t\n", strings.EqualFold("AB", "ab")) // true because comparison uses simple case-folding
	fmt.Printf("EqualFold(\"ß\", \"ss\") res: %t\n\n", strings.EqualFold("ß", "ss")) // false because comparison does not use full case-folding

	fmt.Println("Функція Fields:")
	fmt.Printf("Fields are: %q\n\n", strings.Fields("  foo bar  baz   "))

	fmt.Println("Функція FieldsFunc:")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n\n", strings.FieldsFunc("  foo1;bar2,baz3...", f))

	fmt.Println("Функція HasPrefix:")
	fmt.Printf("HasPrefix(\"Gopher\", \"Go\") res: %t\n", strings.HasPrefix("Gopher", "Go"))
	fmt.Printf("HasPrefix(\"Gopher\", \"C\") res: %t\n", strings.HasPrefix("Gopher", "C"))
	fmt.Printf("HasPrefix(\"Gopher\", \"\") res: %t\n\n", strings.HasPrefix("Gopher", ""))

	fmt.Println("Функція HasSuffix:")
	fmt.Printf("HasSuffix(\"Amigo\", \"go\") res: %t\n", strings.HasSuffix("Amigo", "go"))
	fmt.Printf("HasSuffix(\"Amigo\", \"O\") res: %t\n", strings.HasSuffix("Amigo", "O"))
	fmt.Printf("HasSuffix(\"Amigo\", \"Ami\") res: %t\n", strings.HasSuffix("Amigo", "Ami"))
	fmt.Printf("HasSuffix(\"Amigo\", \"\") res: %t\n\n", strings.HasSuffix("Amigo", ""))

	fmt.Println("Соррі яле я небачу сенса в тому щоб переписувати довідник по функціям")
	/*
		func Index ¶
		func Index(s, substr string) int
		Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			fmt.Println(strings.Index("chicken", "ken"))
			fmt.Println(strings.Index("chicken", "dmr"))
		}

		Output:

		4
		-1
		Share
		Format
		Run
		func IndexAny ¶
		func IndexAny(s, chars string) int
		IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			fmt.Println(strings.IndexAny("chicken", "aeiouy"))
			fmt.Println(strings.IndexAny("crwth", "aeiouy"))
		}

		Output:

		2
		-1
		Share
		Format
		Run
		func IndexByte ¶
		added in go1.2
		func IndexByte(s string, c byte) int
		IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			fmt.Println(strings.IndexByte("golang", 'g'))
			fmt.Println(strings.IndexByte("gophers", 'h'))
			fmt.Println(strings.IndexByte("golang", 'x'))
		}

		Output:

		0
		3
		-1
		Share
		Format
		Run
		func IndexFunc ¶
		func IndexFunc(s string, f func(rune) bool) int
		IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
			"unicode"
		)

		func main() {
			f := func(c rune) bool {
				return unicode.Is(unicode.Han, c)
			}
			fmt.Println(strings.IndexFunc("Hello, 世界", f))
			fmt.Println(strings.IndexFunc("Hello, world", f))
		}

		Output:

		7
		-1
		Share
		Format
		Run
		func IndexRune ¶
		func IndexRune(s string, r rune) int
		IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s. If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			fmt.Println(strings.IndexRune("chicken", 'k'))
			fmt.Println(strings.IndexRune("chicken", 'd'))
		}

		Output:

		4
		-1
		Share
		Format
		Run
		func Join ¶
		func Join(elems []string, sep string) string
		Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			s := []string{"foo", "bar", "baz"}
			fmt.Println(strings.Join(s, ", "))
		}

		Output:

		foo, bar, baz
		Share
		Format
		Run
		func LastIndex ¶
		func LastIndex(s, substr string) int
		LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.

		Example ¶
		func LastIndexAny ¶
		func LastIndexAny(s, chars string) int
		LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.

		Example ¶
		func LastIndexByte ¶
		added in go1.5
		func LastIndexByte(s string, c byte) int
		LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

		Example ¶
		func LastIndexFunc ¶
		func LastIndexFunc(s string, f func(rune) bool) int
		LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.

		Example ¶
		func Map ¶
		func Map(mapping func(rune) rune, s string) string
		Map returns a copy of the string s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement.

		Example ¶
		package main

		import (
			"fmt"
			"strings"
		)

		func main() {
			rot13 := func(r rune) rune {
				switch {
				case r >= 'A' && r <= 'Z':
					return 'A' + (r-'A'+13)%26
				case r >= 'a' && r <= 'z':
					return 'a' + (r-'a'+13)%26
				}
				return r
			}
			fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
		}

		Output:

		'Gjnf oevyyvt naq gur fyvgul tbcure...
	*/

}
