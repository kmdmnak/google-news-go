package googlenews

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type searchKind string

const (
	AND searchKind = "AND"
	OR  searchKind = "OR"
)

func filterURL(urlStr string) string {
	if len(urlStr) > 0 {
		return fmt.Sprintf("inurl:%s", urlStr)
	}
	return ""
}

const dateLayout = "2006-01-02"

func dateString(t *time.Time) string {
	return t.Format(dateLayout)
}

func filterAfter(t *time.Time) string {
	if t != nil {
		return fmt.Sprintf("after:%s", dateString(t))
	}
	return ""
}

func filterBefore(t *time.Time) string {
	if t != nil {
		return fmt.Sprintf("before:%s", dateString(t))
	}
	return ""
}

func filterWords(words []string, kind searchKind) string {
	var f bool = true
	var b strings.Builder
	for _, w := range words {
		if w == "" {
			continue
		}
		escaped := fmt.Sprintf("\"%s\"", url.QueryEscape(w))
		if f {
			b.WriteString("(")
			b.WriteString(escaped)
		} else {
			b.WriteString(fmt.Sprintf(" %s %s", kind, escaped))
		}
		f = false
	}
	if !f {
		b.WriteString(")")
	}
	return b.String()
}

type QueryParameter struct {
	After  *time.Time
	Before *time.Time
	// TODO implement and or or search
	Words []string
	// source of news
	Media string
}

func (qp *QueryParameter) validate() error {
	if qp.After == nil || qp.Before == nil {
		return nil
	}
	if qp.After.After(*qp.Before) {
		return errors.New("after must be lower than before")
	}
	return nil
}

// buildQueryString build url query for query search
func (qp *QueryParameter) buildQueryString() string {
	conditions := []string{}
	if len(qp.Words) > 0 {
		conditions = append(conditions, filterWords(qp.Words, OR))
	}
	if qp.Media != "" {
		conditions = append(conditions, filterURL(qp.Media))
	}
	if qp.After != nil {
		conditions = append(conditions, filterAfter(qp.After))
	}
	if qp.Before != nil {
		conditions = append(conditions, filterBefore(qp.Before))
	}
	return strings.Join(conditions, " + ")
}
