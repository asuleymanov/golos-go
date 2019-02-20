package api

//account_by_key

//GetKeyReferences api request get_key_references
func (api *API) GetKeyReferences(pubkey string) ([]*string, error) {
	var resp []*string
	err := api.call("account_by_key", "get_key_references", []interface{}{pubkey}, &resp)
	return resp, err
}
