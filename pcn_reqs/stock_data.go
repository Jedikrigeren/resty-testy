package pcn_reqs

import (
	"os"

	"github.com/imroc/req/v3"
)

type result struct {
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	Results []struct {
		Articleno   string `json:"articleno"`
		Description string `json:"description"`
		Instock     int    `json:"instock"`
		Onorder     int    `json:"onorder"`
		Forsale     int    `json:"forsale"`
		Barcode     string `json:"barcode"`
		Barcode2    string `json:"barcode2"`
	}
}

func FetchStockData() (*result, error) {
	client := req.C()
	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"cid":        os.Getenv("PCN_CID"),
			"olsuser":    os.Getenv("PCN_OLSUSER"),
			"olspass":    os.Getenv("PCN_OLSPASS"),
			"filter":     "all",
			"maxresults": 1000,
		}).
		SetBasicAuth(os.Getenv("PCN_AUTH_UN"), os.Getenv("PCN_AUTH_PW")).
		SetResult(result{}).
		Post("https://aal14.pakkecenternord.dk/rest/v6_1/api.php?rquest=stocklist")
	if err != nil {
		// Handle err
		return &result{}, err
	}

	return resp.Result().(*result), nil
}
