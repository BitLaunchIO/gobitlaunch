package gobitlaunch

import (
	"encoding/json"
	"time"
)

// Ports represents a port slice object
type Ports struct {
	PortNumber int    `json:"portNumber"`
	Protocol   string `json:"protocol"`
}

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
	Protection         struct {
		Enabled bool `json:"enabled"`
		Proxy   struct {
			IP     string  `json:"ip"`
			Region string  `json:"region"`
			Ports  []Ports `json:"ports"`
			Target string  `json:"target"`
		} `json:"proxy"`
	} `json:"protection"`
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

// RebuildOptions defines options for rebuilding a server
type RebuildOptions struct {
	ID          string `json:"hostImageID"`
	Description string `json:"imageDescription"`
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

// Stop a server
func (ss *ServerService) Stop(id string) error {
	req, err := ss.client.NewRequest("POST", "/servers/"+id+"/stop", nil)
	if err != nil {
		return err
	}

	if err := ss.client.DoRequest(req, nil); err != nil {
		return err
	}

	return nil
}

// Rebuild server
func (ss *ServerService) Rebuild(id string, opts *RebuildOptions) error {
	b, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	req, err := ss.client.NewRequest("POST", "/servers/"+id+"/rebuild", b)
	if err != nil {
		return err
	}

	if err := ss.client.DoRequest(req, nil); err != nil {
		return err
	}

	return nil
}

// Resize server
func (ss *ServerService) Resize(id string, size string) error {
	b, err := json.Marshal(map[string]string{
		"size": size,
	})
	if err != nil {
		return err
	}
	req, err := ss.client.NewRequest("POST", "/servers/"+id+"/resize", b)
	if err != nil {
		return err
	}

	if err := ss.client.DoRequest(req, nil); err != nil {
		return err
	}

	return nil
}

// Restart server
func (ss *ServerService) Restart(id string) error {
	req, err := ss.client.NewRequest("POST", "/servers/"+id+"/restart", nil)
	if err != nil {
		return err
	}

	if err := ss.client.DoRequest(req, nil); err != nil {
		return err
	}

	return nil
}

// Protection is to enable/disable DDoS protection on a server
func (ss *ServerService) Protection(id string, enabled bool) (*Server, error) {
	opts := struct {
		Enabled bool   `json:"enable"`
		Region  string `json:"region"`
	}{
		Enabled: enabled,
		Region: func() string {
			if enabled {
				return "bvm-lux"
			}
			return ""
		}(),
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req, err := ss.client.NewRequest("POST", "/servers/"+id+"/protection", b)
	if err != nil {
		return nil, err
	}

	s := Server{}

	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

// SetPorts sets the enabled ports on a server for DDoS protection (Needs Protection enabled)
func (ss *ServerService) SetPorts(id string, ports *[]Ports) (*Server, error) {
	b, err := json.Marshal(ports)
	if err != nil {
		return nil, err
	}

	req, err := ss.client.NewRequest("POST", "/servers/"+id+"/protection/ports", b)
	if err != nil {
		return nil, err
	}

	s := Server{}

	if err := ss.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
