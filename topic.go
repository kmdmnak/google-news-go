package googlenews

type NewsTopic uint8

const (
	TOPIC_WORLD NewsTopic = iota
	TOPIC_NATION
	TOPIC_BUSINESS
	TOPIC_TECHNOLOGY
	TOPIC_ENTERTAINMENT
	TOPIC_SPORTS
	TOPIC_SCIENCE
)

func (topic NewsTopic) string() string {
	switch topic {
	case TOPIC_WORLD:
		return "WORLD"
	case TOPIC_NATION:
		return "NATION"
	case TOPIC_BUSINESS:
		return "BUSINESS"
	case TOPIC_TECHNOLOGY:
		return "TECHNOLOGY"
	case TOPIC_ENTERTAINMENT:
		return "ENTERTAINMENT"
	case TOPIC_SPORTS:
		return "SPORTS"
	case TOPIC_SCIENCE:
		return "SCIENCE"
	}
	return ""
}
