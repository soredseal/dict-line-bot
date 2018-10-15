package dictionary

type OxfordClient struct {}

func (o OxfordClient) GetDefinitions(word string) ([]string, error) {
	return []string{}, nil
}

func (o OxfordClient) GetSynonyms(word string) ([]string, error) {
	return []string{}, nil
}