package urbndct

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Terms struct {
	List []struct {
		Definition string `json:"definition"`
		Permalink  string `json:"permalink"`
		ThumbsUp   int    `json:"thumbs_up"`
		Example    string `json:"example"`
		ThumbsDown int    `json:"thumbs_down"`
	} `json:"list"`
}

func GetTopResult(q string) (string, error) {
	terms, err := search(q)
	if err != nil {
		return "", err
	}

	return fmtSearchResults(q, terms), nil
}

func search(q string) (*Terms, error) {
	udUrl := fmt.Sprintf("https://api.urbandictionary.com/v0/define?term=%s", url.QueryEscape(q))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, udUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var definitions *Terms
	err = json.NewDecoder(res.Body).Decode(&definitions)
	if err != nil {
		return nil, err
	}

	return definitions, nil
}

func fmtSearchResults(q string, terms *Terms) string {
	res := fmt.Sprintf("\nNo results for \"%s\"\n", q)

	if len(terms.List) > 0 {
		result := terms.List[0]
		result.Definition = strings.ReplaceAll(result.Definition, "\n", "\n   ")
		result.Example = strings.ReplaceAll(result.Example, "\n", "\n   ")
		res = fmt.Sprintf("\n%s: %s\n\n   %s\n   👍 %d   👎 %d\n\n   See more results for \"%s\": %s\n", q, result.Definition, result.Example, result.ThumbsUp, result.ThumbsDown, q, result.Permalink)
	}

	return res
}
