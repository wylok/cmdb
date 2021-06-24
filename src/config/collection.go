package config

type DiskInfo struct {
	MountPoint string `json:"mount_point"`
	FsType     string `json:"fs_type"`
	Size       uint64 `json:"size"`
}

type NetInfo struct {
	HardwareAddr string   `json:"hardwareaddr"`
	Addrs        []string `json:"addrs"`
}

type HardWareConf struct {
	UUID         string              `json:"uuid"`
	SerialNumber string              `json:"serial_number"`
	Manufacturer string              `json:"manufacturer"`
	ProductName  string              `json:"product_name"`
	CpuCore      int                 `json:"cpu_core"`
	CpuInfo      string              `json:"cpu_info"`
	MemSize      uint64              `json:"mem_size"`
	Disk         map[string]DiskInfo `json:"disk"`
	Net          map[string]NetInfo  `json:"net"`
}

type SystemConf struct {
	HostID          string `json:"host_id"`
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	Uptime          uint64 `json:"uptime"`
}

type CollectionData struct {
	HardWare HardWareConf `json:"hardware"`
	System   SystemConf   `json:"system"`
}
