package api

import (
	"encoding/json"
	"time"
)

//private_message

//GetInbox api request get_inbox
func (api *API) GetInbox(to string, limit, offset uint16) ([]*Message, error) {
	var resp []*Message
	err := api.call("private_message", "get_inbox", []interface{}{to, time.Now(), limit, offset}, &resp)
	return resp, err
}

//GetOutbox api request get_outbox
func (api *API) GetOutbox(from string, limit, offset uint16) ([]*Message, error) {
	var resp []*Message
	err := api.call("private_message", "get_outbox", []interface{}{from, time.Now(), limit, offset}, &resp)
	return resp, err
}

//GetThread api request get_thread
func (api *API) GetThread(from, to string, limit, offset uint16) ([]*Message, error) {
	var resp []*Message
	err := api.call("private_message", "get_thread", []interface{}{from, to, limit, offset}, &resp)
	return resp, err
}

//GetSettings api request get_settings
func (api *API) GetSettings(owner string) (*MessageSettings, error) {
	var resp MessageSettings
	err := api.call("private_message", "get_settings", []interface{}{owner}, &resp)
	return &resp, err
}

//GetContactsSize api request get_contacts_size
func (api *API) GetContactsSize(owner string) (*json.RawMessage, error) {
	var resp json.RawMessage
	err := api.call("private_message", "get_contacts_size", []interface{}{owner}, &resp)
	return &resp, err
}

//GetContactInfo api request get_contact_info
func (api *API) GetContactInfo(from, contact string) (*MessageContact, error) {
	var resp MessageContact
	err := api.call("private_message", "get_contact_info", []interface{}{from, contact}, &resp)
	return &resp, err
}

//GetContacts api request get_contacts
func (api *API) GetContacts(from, contact string) ([]*MessageContact, error) {
	var resp []*MessageContact
	err := api.call("private_message", "get_contacts", []interface{}{from, contact}, &resp)
	return resp, err
}
