package telebot

import "encoding/json"

func (b *bot) AnswerWebAppQuery(webAppQueryID string, result QueryResult) (*SentWebAppMessage, error) {
	params := struct {
		ID    string      `json:"web_app_query_id"`
		Resul QueryResult `json:"result"`
	}{
		ID:    webAppQueryID,
		Resul: result,
	}

	var resp struct {
		Result *SentWebAppMessage `json:"result"`
	}

	res, err := b.sendMethodRequest(methodAnswerWebAppQuery, params)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, err
}
