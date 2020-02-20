package api

//account_by_key

//GetKeyReferences returns a list of accounts by a list of public keys
func (api *API) GetKeyReferences(pubkey ...string) ([]*string, error) {
	var resp []*string
	err := api.call("account_by_key", "get_key_references", []interface{}{pubkey}, &resp)
	return resp, err
}
