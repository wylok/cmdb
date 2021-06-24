package api

import (
	"cmdb-go/config"
	"cmdb-go/database"
	"cmdb-go/kits"
	"cmdb-go/modules"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type UpdateAgent struct {
	AgentId string `json:"agent_id" binding:"required"`
	Data    string `json:"data" binding:"required"`
}

func Errors(err error, message string, success bool) (string, bool) {
	success = false
	message = err.Error()
	kits.Log(err.Error(), "error")
	return message, success
}

// @Tags 主机Agent
// @Summary Agent上报数据
// @Produce  json
// @Param agent_id query string true "AgentId"
// @Param data query string true "Data"
// @Success 200 {string} json "{"success":True,"data":{},"msg":"ok"}"
// @Router /web/v1/cmdb/agent [post]
func AgentUpdate(c *gin.Context) {
	//上报硬件信息
	var AssetServer []database.AssetServer
	var AssetIdc []database.AssetIdc
	var JsonData UpdateAgent
	var ServerData config.CollectionData
	now := time.Now()
	now.Format("2006-01-02 15:04:05")
	AssetType := "physical"
	AssetTypeCn := config.AssetTypes[AssetType]
	rc := modules.RC
	err := c.BindJSON(&JsonData)
	message := "ok"
	success := true
	data := map[string]string{}
	if err != nil {
		message, success = Errors(err, message, false)
	} else {
		// 解密数据
		Crypt := kits.NewEncrypt([]byte(config.CryptKey), 16)
		v, err := Crypt.DecryptString(JsonData.Data)
		if err != nil {
			message, success = Errors(err, message, false)
		} else {
			err := json.Unmarshal(v, &ServerData)
			if err != nil {
				message, success = Errors(err, message, false)
			} else {
				HardWare := &ServerData.HardWare
				System := &ServerData.System
				// 初始化数据库连接
				db, err := database.Init()
				if err != nil {
					message, success = Errors(err, message, false)
				} else {
					for _, AgentId := range config.ValidAgents {
						if JsonData.AgentId == AgentId {
							// cmdb信息采集以api信息为准
							rv, _ := rc.HExists(config.ApiAssetKey, System.HostID).Result()
							if rv {
							} else {
								// 硬件信息写入或者修改
								if strings.Contains(HardWare.Manufacturer, "VMware") {
									HardWare.Manufacturer = "VMware"
									AssetType = "virtual"
									AssetTypeCn = config.AssetTypes[AssetType]
									Md5Server := kits.MD5(HardWare.UUID + System.HostID + HardWare.ProductName +
										HardWare.Manufacturer + HardWare.CpuInfo + "Running")
									db.Select("md5_verify,host_id").Where(
										&database.AssetServer{Uuid: HardWare.UUID}).First(&AssetServer)
									for _, as := range AssetServer {
										if as.Md5Verify != Md5Server {
											db.Model(&database.AssetServer{}).Where(
												&database.AssetServer{Uuid: HardWare.UUID}).Updates(
												database.AssetServer{HostId: System.HostID,
													Sn:           HardWare.SerialNumber,
													Manufacturer: HardWare.Manufacturer,
													ProductName:  HardWare.ProductName,
													CpuInfo:      HardWare.CpuInfo,
													Status:       "Running",
													Md5Verify:    Md5Server})
											println(AssetTypeCn)
											db.Model(&database.AssetSystem{}).Where(
												&database.AssetSystem{HostId: System.HostID}).Updates(
												database.AssetSystem{Uptime: System.Uptime})
										}
										// host_id变更
										if as.HostId != System.HostID {
											db.Model(&database.AssetSystem{}).Where(&database.AssetSystem{
												HostId: as.HostId}).Updates(database.AssetSystem{HostId: System.HostID})
											db.Model(&database.AssetDisk{}).Where(&database.AssetDisk{
												HostId: as.HostId}).Updates(database.AssetDisk{HostId: System.HostID})
											db.Model(&database.AssetNet{}).Where(&database.AssetNet{
												HostId: as.HostId}).Updates(database.AssetNet{HostId: System.HostID})
											db.Model(&database.Partition{}).Where(&database.Partition{
												Pid: as.HostId}).Updates(database.Partition{Pid: System.HostID})
											db.Model(&database.AssetExtend{}).Where(&database.AssetExtend{
												HostId: as.HostId}).Updates(database.AssetExtend{HostId: System.HostID})
										}
										if as == (database.AssetServer{}) {
											db.Select("idc_id").Where(&database.AssetIdc{
												Idc: "21vianet", Region: "cn-beijing"}).First(&AssetIdc)
											for _, ai := range AssetIdc {
												if ai.IdcId != "" {
													var as = database.AssetServer{Uuid: HardWare.UUID,
														HostId: System.HostID, IdcId: ai.IdcId,
														Sn:           HardWare.SerialNumber,
														ProductName:  HardWare.ProductName,
														Manufacturer: HardWare.Manufacturer,
														CpuInfo:      HardWare.CpuInfo,
														ChargeType:   "None", Tags: "None",
														Status: "Running", CreateTime: now,
														HeartbeatTime: now, Comment: "None",
														Md5Verify: Md5Server}
													db.Create(&as)
												}
											}
										}
										// 系统信息写入或者修改
									}
								}
							}
						}
					}
				}
				defer db.Close()
			}
		}
	}
	defer func() {
		if r := recover(); r != nil {
			success = false
		}
		c.JSON(http.StatusOK, gin.H{"success": success, "message": message, "data": data})
	}()
}
