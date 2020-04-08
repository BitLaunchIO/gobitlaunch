package gobitlaunch

import (
	"encoding/json"
	"time"
)

// Server represents a server
type Server struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	HostID             int       `json:"host"`
	Ipv4               string    `json:"ipv4"`
	Region             string    `json:"region"`
	Size               string    `json:"size"`
	SizeDesc           string    `json:"sizeDescription"`
	Image              string    `json:"image"`
	ImageDesc          string    `json:"imageDescription"`
	Created            time.Time `json:"created"`
	Rate               int       `json:"rate"`
	BandwidthUsed      int       `json:"bandwidthUsed"`
	BandwidthAllowance int       `json:"bandwidthAllowance"`
	Status             string    `json:"status"`
	ErrorText          string    `json:"errorText"`
	BackupsEnabled     bool      `json:"backupsEnabled"`
	Version            string    `json:"version"`
	Abuse              bool      `json:"abuse"`
	DiskGB             int       `json:"diskGB"`
}

// CreateServerOptions defines options for creating a new server
type CreateServerOptions struct {
	Name        string   `json:"name"`
	HostID      int      `json:"hostID"`
	HostImageID string   `json:"HostImageID"`
	SizeID      string   `json:"sizeID"`
	RegionID    string   `json:"regionID"`
	SSHKeys     []string `json:"sshKeys"`
	Password    string   `json:"password"`
	InitScript  string   `json:"initscript"`
}

// ServerService manages server API actions
type ServerService struct {
	client *Client
}

// Create server
func (ss *ServerService) Create(opts *CreateServerOptions) (*Server, error) {
	b, err := json.Marshal(map[string]*CreateServerOptions{
		"server": opts,
	})
	if err != nil {
		return nil, err
	}
	req, err := ss.client.NewRequest("POST", "/servers", b)
	if err != nil {
		return nil, err
	}

	s := Server{}
	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

// Show server
func (ss *ServerService) Show(id string) (*Server, error) {
	req, err := ss.client.NewRequest("GET", "/servers/"+id, nil)
	if err != nil {
		return nil, err
	}

	s := struct {
		Server Server
	}{}
	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s.Server, nil
}

// List servers
func (ss *ServerService) List() ([]Server, error) {
	req, err := ss.client.NewRequest("GET", "/servers", nil)
	if err != nil {
		return nil, err
	}

	servers := []Server{}
	if err := ss.client.DoRequest(req, &servers); err != nil {
		return nil, err
	}

	return servers, nil
}

// Destroy a server
func (ss *ServerService) Destroy(id string) error {
	req, err := ss.client.NewRequest("DELETE", "/servers/"+id, nil)
	if err != nil {
		return err
	}

	if err := ss.client.DoRequest(req, nil); err != nil {
		return err
	}

	return nil
}
