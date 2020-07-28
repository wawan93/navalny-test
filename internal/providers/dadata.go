package providers

import "gopkg.in/webdeskltd/dadata.v2"

type Dadata struct {
	api *dadata.DaData
}

func NewDadataProvider(token, secret string) *Dadata {
	api := dadata.NewDaData(token, secret)
	return &Dadata{
		api: api,
	}
}

func (p *Dadata) Suggest(query string) ([]string, error) {
	params := dadata.SuggestRequestParams{Query: query}
	response, err := p.api.SuggestAddresses(params)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, res := range response {
		result = append(result, res.Value)
	}
	return result, nil
}
