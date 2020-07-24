package providers

import "gopkg.in/webdeskltd/dadata.v2"

type Dadata struct {
	Token string
}

func (p *Dadata) Suggest(query string) ([]string, error) {
	api := dadata.NewDaData(p.Token, "")
	params := dadata.SuggestRequestParams{Query: query}
	response, err := api.SuggestAddresses(params)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, res := range response {
		result = append(result, res.Value)
	}
	return result, nil
}
