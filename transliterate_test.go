package str

import "testing"

func TestTransliterateBasic(t *testing.T) {
	t.Parallel()

	cases := map[rune]string{
		'à': "a", 'á': "a", 'â': "a", 'ã': "a", 'ä': "a", 'å': "a", 'æ': "a",
		'ç': "c",
		'è': "e", 'é': "e", 'ê': "e", 'ë': "e",
		'ì': "i", 'í': "i", 'î': "i", 'ï': "i",
		'ñ': "n",
		'ò': "o", 'ó': "o", 'ô': "o", 'õ': "o", 'ö': "o", 'ø': "o",
		'ß': "ss",
		'ù': "u", 'ú': "u", 'û': "u", 'ü': "u",
		'ý': "y", 'ÿ': "y",
	}

	for r, want := range cases {
		got, ok := transliterateBasic(r)
		if !ok || got != want {
			t.Fatalf("transliterateBasic(%q) = %q, %v; want %q, true", string(r), got, ok, want)
		}
	}
	if _, ok := transliterateBasic('z'); ok {
		t.Fatalf("unexpected transliteration for z")
	}
}
