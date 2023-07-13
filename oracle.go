package oracle

type Response struct {
	itemsArray []Item `json:"items"`
	TotalResults int `json:"total_results"`
	Page int `json:"page"`
	Offset int `json:"offset"`
}

type Item struct{
	ItemName string `json:"item_name"`
	ItemNumeber string `json:"item_number"`
	Expiry string `json:"expiry"`
	Quantity string `json:"total_quantity"`
}


type OracleGateway interface{
	FetchItemsToReceive(ctx context.Context)
}

func (o *OracleGateway) FetchItemsToReceive(ctx context.Context) (*,) {
	response, err := http.Get("http://oracle-link")
    if err != nil {
        log,Error(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)
	return &responseObject,nil
}