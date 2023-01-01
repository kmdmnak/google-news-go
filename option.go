package googlenews

type Language uint8

const (
	// us english
	USEN Language = iota
	// uk english
	UKEN
	// chinese
	CHIN
	// germany
	GERM
	// spanish
	SPAN
	// arabian
	ARAB
	// japanese
	JPN
)

type langProperty struct {
	Hl   string
	Gl   string
	CeID string
}

func getLanguageProperty(lang Language) *langProperty {
	switch lang {
	case USEN:
		return &langProperty{Hl: "en-US", Gl: "US", CeID: "US:en"}
	case UKEN:
		return &langProperty{Hl: "en-GB", Gl: "GB", CeID: "GB:en"}
	case CHIN:
		return &langProperty{Hl: "zh-CN", Gl: "CN", CeID: "CN:zh-Hans"}
	case GERM:
		return &langProperty{Hl: "de", Gl: "DE", CeID: "DE:de"}
	case SPAN:
		return &langProperty{Hl: "es-419", Gl: "US", CeID: "US:es-419"}
	case ARAB:
		return &langProperty{Hl: "ar", Gl: "EG", CeID: "EG:ar"}
	}
	return &langProperty{Hl: "ja", Gl: "JP", CeID: "JP:ja"}
}
