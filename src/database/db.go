package database

import (
	"cmdb-go/kits"
	"cmdb-go/modules"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Init() (*gorm.DB, error) {
	// 从apollo获取数据库配置信息
	v := modules.ApolloConfig("sql_go_config")
	conf := kits.StringToMap(v.(string))
	var err error
	//open a db connection
	DB, err = gorm.Open("mysql", conf["neolink_cmdb"].(string))
	if err == nil {
		DB.DB().SetMaxOpenConns(50)
		DB.DB().SetMaxIdleConns(50)
	}
	if err != nil {
		println(err.Error())
	} else {
		println("数据库连接成功!")
		return DB, err
	}
	return nil, err
}
func init() {
	_, _ = Init()
}

type UpdateAgent struct {
	AgentId string `json:"agent_id" binding:"required"`
	Data    string `json:"data" binding:"required"`
}

type ApschedulerJobs struct {
	Id          uint64 `gorm:"primary_key" json:"id"`
	NextRunTime uint64 `gorm:"column:next_run_time" json:"next_run_time"`
	JobState    byte   `gorm:"column:job_state" json:"job_state"`
}

func (ApschedulerJobs) TableName() string {
	return "apscheduler_jobs"
}

type Partition struct {
	Id       uint64 `gorm:"primary_key" json:"id"`
	TenantId string `gorm:"column:tenant_id;type:varchar(100)" json:"tenant_id"`
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`
	Pid      string `gorm:"column:pid;type:varchar(100)" json:"pid"`
	Source   string `gorm:"column:source;type:varchar(100)" json:"source"`
	Rule     string `gorm:"column:rule;type:varchar(100)" json:"rule"`
}

func (Partition) TableName() string {
	return "partition"
}

type AssetIdc struct {
	Id       uint64 `gorm:"primary_key" json:"id"`
	IdcId    string `gorm:"column:idc_id;type:varchar(100)" json:"idc_id"`
	Idc      string `gorm:"column:idc;type:varchar(50)" json:"idc"`
	IdcCn    string `gorm:"column:idc_cn;type:varchar(50)" json:"idc_cn"`
	Cabinet  string `gorm:"column:cabinet;type:varchar(50)" json:"cabinet"`
	Region   string `gorm:"column:region;type:varchar(50)" json:"region"`
	RegionCn string `gorm:"column:region_cn;type:varchar(50)" json:"region_cn"`
}

func (AssetIdc) TableName() string {
	return "asset_idc"
}

type AssetServer struct {
	Id            uint64    `gorm:"primary_key" json:"id"`
	Uuid          string    `gorm:"column:uuid;type:varchar(100)" json:"uuid"`
	HostId        string    `gorm:"column:host_id;type:varchar(100)" json:"host_id"`
	IdcId         string    `gorm:"column:uuid;type:varchar(100)" json:"idc_id"`
	Sn            string    `gorm:"column:sn;type:varchar(200)" json:"sn"`
	ProductName   string    `gorm:"column:product_name;type:varchar(200)" json:"product_name"`
	Manufacturer  string    `gorm:"column:Manufacturer;type:varchar(200)" json:"Manufacturer"`
	CpuInfo       string    `gorm:"column:cpu_info;type:varchar(100)" json:"cpu_info"`
	ChargeType    string    `gorm:"column:ChargeType;type:varchar(100)" json:"ChargeType"`
	Tags          string    `gorm:"column:Tags;type:varchar(100)" json:"Tags"`
	Status        string    `gorm:"column:status;type:varchar(100)" json:"status"`
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time;type:datetime" json:"heartbeat_time"`
	Comment       string    `gorm:"column:comment;type:varchar(100)" json:"comment"`
	Md5Verify     string    `gorm:"column:md5_verify;type:varchar(100)" json:"md5_verify"`
}

func (AssetServer) TableName() string {
	return "asset_server"
}

type AssetExtend struct {
	Id         uint64    `gorm:"primary_key" json:"id"`
	HostId     string    `gorm:"column:Host_id;type:varchar(100)" json:"Host_id"`
	OrderId    string    `gorm:"column:order_id;type:varchar(100)" json:"order_id"`
	OrderType  string    `gorm:"column:order_type;type:varchar(100)" json:"order_type"`
	PurchTime  time.Time `gorm:"column:purch_time;type:datetime" json:"purch_time"`
	ExpirdTime time.Time `gorm:"column:expird_time;type:datetime" json:"expird_time"`
	Md5Verify  string    `gorm:"column:md5_verify;type:varchar(100)" json:"md5_verify"`
}

func (AssetExtend) TableName() string {
	return "asset_extend"
}

type AssetSystem struct {
	Id        uint64 `gorm:"primary_key" json:"id"`
	HostId    string `gorm:"column:Host_id;type:varchar(100)" json:"Host_id"`
	HostName  string `gorm:"column:host_name;type:varchar(200)" json:"host_name"`
	AssetType string `gorm:"column:asset_type;type:enum('physical', 'virtual',
'container', 'aliyun_ecs')" json:"asset_type"`
	AssetTypeCn string `gorm:"column:asset_type_cn;type:enum('物理机', '虚拟机', '容器',
'阿里云主机')" json:"asset_type_cn"`
	Cpu             uint8     `gorm:"column:cpu;type:int(11)" json:"cpu"`
	Memory          uint32    `gorm:"column:memory;type:bigint(50)" json:"memory"`
	Disk            uint64    `gorm:"column:disk;type:bigint(50)" json:"disk"`
	Os              string    `gorm:"column:os;type:varchar(100)" json:"os"`
	Platform        string    `gorm:"column:platform;type:varchar(100)" json:"platform"`
	PlatformVersion string    `gorm:"column:platform_version;type:varchar(100)" json:"platform_version"`
	KernelVersion   string    `gorm:"column:kernel_version;type:varchar(100)" json:"kernel_version"`
	Uptime          uint64    `gorm:"column:uptime;type:int(11)" json:"uptime"`
	UpdateTime      time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	Status          string    `gorm:"column:status;type:enum('Used', 'Available')" json:"status"`
	Md5Verify       string    `gorm:"column:md5_verify;type:varchar(100)" json:"md5_verify"`
	Comment         string    `gorm:"column:comment;type:varchar(100)" json:"comment"`
}

func (AssetSystem) TableName() string {
	return "asset_system"
}

type AssetDisk struct {
	Id         uint64 `gorm:"primary_key" json:"id"`
	HostId     string `gorm:"column:host_id;type:varchar(100)" json:"host_id"`
	Name       string `gorm:"column:name;type:varchar(200)" json:"name"`
	MountPoint string `gorm:"column:mountpoint;type:varchar(100)" json:"mountpoint"`
	Fstype     string `gorm:"column:fstype;type:varchar(100)" json:"fstype"`
	Size       uint64 `gorm:"column:size;type:bigint(50)" json:"size"`
	Md5Verify  string `gorm:"column:Md5_verify;type:varchar(100)" json:"Md5_verify"`
}

func (AssetDisk) TableName() string {
	return "asset_disk"
}

type AssetNet struct {
	Id        uint64 `gorm:"primary_key" json:"id"`
	HostId    string `gorm:"column:host_id;type:varchar(100)" json:"host_id"`
	Name      string `gorm:"column:name;type:varchar(200)" json:"name"`
	Addr      string `gorm:"column:addr;type:varchar(100)" json:"addr"`
	Ip        string `gorm:"column:ip;type:varchar(100)" json:"ip"`
	Netmask   string `gorm:"column:netmask;type:varchar(100)" json:"netmask"`
	PublicIp  string `gorm:"column:public_ip;type:varchar(100)" json:"public_ip"`
	Valid     string `gorm:"column:valid;type:enum('True', 'False')" json:"valid"`
	Md5Verify string `gorm:"column:Md5_verify;type:varchar(100)" json:"Md5_verify"`
}

func (AssetNet) TableName() string {
	return "asset_net"
}

type AgentPool struct {
	Id      uint64 `gorm:"primary_key" json:"id"`
	Uid     string `gorm:"column:uid;type:varchar(100)" json:"uid"`
	AgentId string `gorm:"column:agent_id;type:varchar(100)" json:"agent_id"`
	Status  string `gorm:"column:status;type:enum('active', 'close')" json:"status"`
	Comment string `gorm:"column:comment;type:varchar(100)" json:"comment"`
}

func (AgentPool) TableName() string {
	return "agent_pool"
}

type AliyunExtend struct {
	Id                      uint64 `gorm:"primary_key" json:"id"`
	HostId                  string `gorm:"column:host_id;type:varchar(100)" json:"host_id"`
	AutoReleaseTime         string `gorm:"column:AutoReleaseTime;type:varchar(100)" json:"AutoReleaseTime"`
	DeletionProtection      string `gorm:"column:DeletionProtection;type:varchar(100)" json:"DeletionProtection"`
	ImageId                 string `gorm:"column:ImageId;type:varchar(100)" json:"ImageId"`
	InstanceNetworkType     string `gorm:"column:InstanceNetworkType;type:varchar(100)" json:"InstanceNetworkType"`
	InternetMaxBandwidthIn  string `gorm:"column:InternetMaxBandwidthIn;type:varchar(100)" json:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut string `gorm:"column:InternetMaxBandwidthOut;type:varchar(100)" json:"InternetMaxBandwidthOut"`
	SaleCycle               string `gorm:"column:SaleCycle;type:varchar(100)" json:"SaleCycle"`
	StoppedMode             string `gorm:"column:StoppedMode;type:varchar(100)" json:"StoppedMode"`
	Md5Verify               string `gorm:"column:md5_verify;type:varchar(100)" json:"md5_verify"`
}

func (AliyunExtend) TableName() string {
	return "aliyun_extend"
}

type AssetVpc struct {
	Id            uint64    `gorm:"primary_key" json:"id"`
	VpcId         string    `gorm:"column:VpcId;type:varchar(200)" json:"VpcId"`
	IdcId         string    `gorm:"column:idc_id;type:varchar(200)" json:"idc_id"`
	Status        string    `gorm:"column:Status;type:varchar(100)" json:"Status"`
	VpcName       string    `gorm:"column:VpcName;type:varchar(500)" json:"VpcName"`
	CreationTime  time.Time `gorm:"column:CreationTime;type:datetime" json:"CreationTime"`
	IPv4          string    `gorm:"column:IPv4;type:varchar(100)" json:"IPv4"`
	VRouterId     string    `gorm:"column:VRouterId;type:varchar(200)" json:"VRouterId"`
	IsDefault     string    `gorm:"column:IsDefault;type:varchar(100)" json:"IsDefault"`
	VSwitchIds    string    `gorm:"column:VSwitchIds;type:varchar(500)" json:"VSwitchIds"`
	Tags          string    `gorm:"column:Tags;type:varchar(100)" json:"Tags"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time;type:datetime" json:"heartbeat_time"`
	Md5Verify     string    `gorm:"column:md5_verify;type:varchar(100)" json:"md5_verify"`
	Comment       string    `gorm:"column:comment;type:varchar(100)" json:"comment"`
}

func (AssetVpc) TableName() string {
	return "asset_vpc"
}

type AssetVSwitches struct {
	Id                      uint64    `gorm:"primary_key" json:"id"`
	VSwitchId               string    `gorm:"column:VSwitchId;type:varchar(200)" json:"VSwitchId"`
	VpcId                   string    `gorm:"column:VpcId;type:varchar(200)" json:"VpcId"`
	Status                  string    `gorm:"column:Status;type:varchar(100)" json:"Status"`
	VSwitchName             string    `gorm:"column:VSwitchName;type:varchar(500)" json:"VSwitchName"`
	CreationTime            time.Time `gorm:"column:CreationTime;type:datetime" json:"CreationTime"`
	IPv4                    string    `gorm:"column:IPv4;type:varchar(100)" json:"IPv4"`
	ZoneId                  string    `gorm:"column:ZoneId;type:varchar(100)" json:"ZoneId"`
	AvailableIpAddressCount uint64    `gorm:"column:AvailableIpAddressCount;
	type:varchar(100)" json:"AvailableIpAddressCount"`
	IsDefault     string    `gorm:"column:IsDefault;type:varchar(100)" json:"IsDefault"`
	Tags          string    `gorm:"column:Tags;type:varchar(100)" json:"Tags"`
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time;type:datetime" json:"heartbeat_time"`
	Md5Verify     string    `gorm:"column:md5_verify;type:varchar(100)" json:"md5_verify"`
	Comment       string    `gorm:"column:comment;type:varchar(100)" json:"comment"`
}

func (AssetVSwitches) TableName() string {
	return "asset_vswitches"
}

type AssetRules struct {
	Id         uint64    `gorm:"primary_key" json:"id"`
	RuleId     string    `gorm:"column:RuleId;type:varchar(200)" json:"RuleId"`
	Cron       string    `gorm:"column:cron;type:varchar(200)" json:"cron"`
	RuleName   string    `gorm:"column:RuleName;type:varchar(200)" json:"RuleName"`
	AccessId   string    `gorm:"column:access_id;type:varchar(200)" json:"access_id"`
	AssetType  string    `gorm:"column:AssetType;type:enum('server', 'vpc', 'vswitch')" json:"AssetType"`
	Rule       string    `gorm:"column:rule;type:varchar(200)" json:"rule"`
	Status     string    `gorm:"column:status;type:enum('active', 'close')" json:"status"`
	Gid        string    `gorm:"column:gid;type:varchar(200)" json:"gid"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	RunTime    time.Time `gorm:"column:run_time;type:datetime" json:"run_time"`
	CreateUser string    `gorm:"column:create_user;type:varchar(200)" json:"create_user"`
	UpdateUser string    `gorm:"column:update_user;type:varchar(200)" json:"update_user"`
	Md5Verify  string    `gorm:"column:md5_verify;type:varchar(200)" json:"md5_verify"`
}

func (AssetRules) TableName() string {
	return "asset_rules"
}
