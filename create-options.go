package gobitlaunch

import (
	"strconv"
)

// HostImageVersion represents an image version
type HostImageVersion struct {
	ID                  string `json:"id"`
	Description         string `json:"description"`
	PasswordUnsupported bool   `json:"passwordUnsupported"`
}

// HostOptions represents what options a certain host provides
type HostOptions struct {
	Rebuild    bool `json:"rebuild"`
	Resize     bool `json:"resize"`
	Backups    bool `json:"backups"`
	Userscript bool `json:"userScript"`
}

// HostImage represents an image
type HostImage struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	Type               string             `json:"type"`
	MinDiskSize        int                `json:"minDiskSize"`
	UnavailableRegions []string           `json:"unavailableRegions"`
	DefaultVersion     HostImageVersion   `json:"version"`
	Versions           []HostImageVersion `json:"versions"`
	ExtraCostPerMonth  int                `json:"extraCostPerMonth"`
	Windows            bool               `json:"windows"`
}

// HostSubRegion represents a sub region
type HostSubRegion struct {
	ID               string   `json:"id"`
	Description      string   `json:"description"`
	Slug             string   `json:"slug"`
	UnavailableSizes []string `json:"unavailableSizes"`
}

// HostRegion represents a region
type HostRegion struct {
	ID               int             `json:"id"`
	Name             string          `json:"name"`
	ISO              string          `json:"iso"`
	DefaultSubregion HostSubRegion   `json:"subregion"`
	Subregions       []HostSubRegion `json:"subregions"`
}

// HostDisks represents a disk
type HostDisks struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
	Size  string `json:"size"`
	Unit  string `json:"unit"`
}

// HostPlanType represents the type of plan
type HostPlanType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// HostSize represents a server size
type HostSize struct {
	ID           string      `json:"id"`
	Slug         string      `json:"slug"`
	BandwidthGB  int         `json:"bandwidthGB"`
	CPUCount     int         `json:"cpuCount"`
	DiskGB       int         `json:"diskGB"`
	Disks        []HostDisks `json:"disks"`
	MemoryMB     int         `json:"memoryMB"`
	CostPerHour  int         `json:"costPerHr"`
	CostPerMonth float64     `json:"costPerMonth"`

	// PlanType indicates the hardware plan of the server (Standard, CPU focused)
	PlanType string `json:"planType"`
}

// ServerCreateOptions represents server creation options
type ServerCreateOptions struct {
	HostID        int            `json:"hostID"`
	Images        []HostImage    `json:"image"`
	Regions       []HostRegion   `json:"region"`
	Sizes         []HostSize     `json:"size"`
	Available     bool           `json:"available"`
	BandwidthCost int            `json:"bandwidthCost"`
	PlanTypes     []HostPlanType `json:"planTypes"`
}

// CreateOptionsService manages create options API actions
type CreateOptionsService struct {
	client *Client
}

// Show the server create options
func (co *CreateOptionsService) Show(hostID int) (*ServerCreateOptions, error) {
	hostIDStr := strconv.Itoa(hostID)
	req, err := co.client.NewRequest("GET", "/hosts-create-options/"+hostIDStr, nil)
	if err != nil {
		return nil, err
	}

	s := ServerCreateOptions{}
	if err := co.client.DoRequest(req, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
