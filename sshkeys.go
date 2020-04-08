package gobitlaunch

import (
	"encoding/json"
	"time"
)

// SSHKey represents an ssh key
type SSHKey struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Fingerprint string    `json:"fingerprint"`
	Content     string    `json:"content"`
	Created     time.Time `json:"created"`
}

// SSHKeyService manages ssh key API actions
type SSHKeyService struct {
	client *Client
}

// Create ssh key
func (ss *SSHKeyService) Create(k *SSHKey) (*SSHKey, error) {
	b, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}
	req, err := ss.client.NewRequest("POST", "/ssh-keys", b)
	if err != nil {
		return nil, err
	}

	s := SSHKey{}
	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

// List ssh key
func (ss *SSHKeyService) List() ([]SSHKey, error) {
	req, err := ss.client.NewRequest("GET", "/ssh-keys", nil)
	if err != nil {
		return nil, err
	}

	s := struct {
		Keys []SSHKey
	}{}
	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return s.Keys, nil
}

// Delete an SSH Key
func (ss *SSHKeyService) Delete(id string) error {
	req, err := ss.client.NewRequest("DELETE", "/ssh-keys/"+id, nil)
	if err != nil {
		return err
	}

	if err := ss.client.DoRequest(req, nil); err != nil {
		return err
	}

	return nil
}
