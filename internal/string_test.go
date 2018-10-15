package internal

import "testing"

func TestReplaceRight(t *testing.T) {
	t.Run("1 times from ',' to and", func(t *testing.T) {
		text := "underline, underscore, stroke, slash, virgule"
		expected := "underline, underscore, stroke, slash and virgule"
		actual := ReplaceRight(text, ", ", " and ", 1)

		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("2 times from ',' to and", func(t *testing.T) {
		text := "underline, underscore, stroke, slash, virgule"
		expected := "underline, underscore, stroke and slash and virgule"
		actual := ReplaceRight(text, ", ", " and ", 2)

		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})

	t.Run("with no effect", func(t *testing.T) {
		text := "underline"
		expected := "underline"
		actual := ReplaceRight(text, ", ", " and ", 1)

		if expected != actual {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	})
}
