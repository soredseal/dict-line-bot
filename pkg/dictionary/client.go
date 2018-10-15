package dictionary

type Client interface {
	GetDefinitions(word string) ([]string, error)
	GetSynonyms(word string) ([]string, error)
}
