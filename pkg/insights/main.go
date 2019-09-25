package insights

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Insights struct {
	Token   string
	Account string
}

func New(account, token string) Insights {
	return Insights{
		Token:   token,
		Account: account,
	}
}

// Send manages sending the necessary data to New Relic Insights
func (i *Insights) Send(data []byte) (string, error) {

	req, err := http.NewRequest("POST", "https://insights-collector.newrelic.com/v1/accounts/"+i.Account+"/events", bytes.NewBuffer(data))
	req.Header.Set("X-Insert-Key", i.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	return string(body), nil
}
