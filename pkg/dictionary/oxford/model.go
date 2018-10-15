package dictionary

const MaxLength = 5

type OxfordResponse struct {
	Results []result `json:"results"`
}

func (r OxfordResponse) GetDefinitions() []string {
	return r.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions
}

func (r OxfordResponse) GetSynonyms() []string {
	cnt := 0
	var synonyms []string

outerLoop:
	for _, subSense := range r.Results[0].LexicalEntries[0].Entries[0].Senses[0].SubSenses {
		for _, synonym := range subSense.Synonyms {
			synonyms = append(synonyms, synonym.Text)
			cnt++
			if cnt == MaxLength {
				break outerLoop
			}
		}
	}
	return synonyms
}

type result struct {
	Id             string         `json:"id"`
	Language       string         `json:"language"`
	LexicalEntries []lexicalEntry `json:"lexicalEntries"`
	Type           string         `json:"type"`
	Word           string         `json:"word"`
}

type lexicalEntry struct {
	Entries         []entry `json:"entries"`
	Language        string  `json:"language"`
	LexicalCategory string  `json:"lexicalCategory"`
	Text            string  `json:"text"`
}

type entry struct {
	HomographNumber string  `json:"homographNumber"`
	Senses          []sense `json:"senses"`
}

type sense struct {
	Definitions []string   `json:"definitions"`
	SubSenses   []subSense `json:"subsenses"`
}

type subSense struct {
	Synonyms []synonym `json:"synonyms"`
}

type synonym struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Text     string `json:"text"`
}
