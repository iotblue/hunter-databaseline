package irrigationdomain

import (
	"databaselineservice/domain/common"
	"databaselineservice/domain/crudfunctions"
	"databaselineservice/domain/infrastructuredomain"
	"databaselineservice/sdk/cervello"
	"errors"
	"fmt"
	"strconv"
)

type HunterStation struct {
	GlobalId          string                 `json:"globalId"`
	Name              string                 `json:"name"`
	IntegrationId     string                 `json:"integrationId"`
	Model             string                 `json:"model"`
	Brand             string                 `json:"brand"`
	Type              string                 `json:"type"`
	LayerName         string                 `json:"layerName"`
	LayerId           float64                `json:"layerId"`
	LayerType         string                 `json:"layerType"`
	ControllerName    string                 `json:"controllerName"`
	ControllerId      string                 `json:"controllerId"`
	ControllerLayerId float64                `json:"controllerLayerId"`
	AreaId            string                 `json:"areaId"`
	AreaNameEnglish   string                 `json:"areaNameEnglish"`
	AreaNameArabic    string                 `json:"areaName"`
	AreaLayerId       float64                `json:"areaLayerId"`
	CityId            string                 `json:"cityId"`
	CityNameEnglish   string                 `json:"cityNameEnglish"`
	CityNameArabic    string                 `json:"cityName"`
	CityLayerId       float64                `json:"cityLayerId"`
	X                 float64                `json:"x"`
	Y                 float64                `json:"y"`
	CreatedAt         string                 `json:"createdAt,omitempty"`
	UpdatedAt         string                 `json:"updatedAt,omitempty"`
	AdditionalInfo    map[string]interface{} `json:"additionalInfo"`
}

func (thisObj *HunterStation) GetMac() string {
	//return thisObj.MacAddress
	return common.EmptyField
}

func (thisObj *HunterStation) GetGlobalId() string {
	return thisObj.GlobalId
}

func (thisObj *HunterStation) GetCommunicationProtocolConfiguration() (string, interface{}) {
	return common.EmptyField, nil
}

func (thisObj *HunterStation) GetName() string {
	return thisObj.Name
}

func (thisObj *HunterStation) GetModel() string {
	return thisObj.Model
}

func (thisObj *HunterStation) ValidateModel() error {
	// un implemented until i hava a list of models
	return nil
}

func (thisObj *HunterStation) GetReferenceName() string {
	return thisObj.IntegrationId
}

func (thisObj *HunterStation) GetClientId() string {
	return ""
}

func (thisObj *HunterStation) GetIP() string {
	return common.EmptyField
}

func (thisObj *HunterStation) GetParentAssetId() string {
	return thisObj.AreaId
}

func (thisObj *HunterStation) GetParentGatewayId() string {
	return thisObj.ControllerId
}

func (thisObj *HunterStation) GetDeviceType() string {
	return cervello.DeviceTypePeriphral
}

func (thisObj *HunterStation) GetTags() []string {
	deviceStateTag := common.MockDevice
	if common.IsPhysicalDevice {
		deviceStateTag = common.GisDevice
	}
	return []string{deviceStateTag,
		"Hunter",
		"station",
		"irrigation",
		thisObj.ControllerName,
		fmt.Sprintf("%s_alarms", thisObj.ControllerName),
		thisObj.AreaNameEnglish,
		thisObj.AreaNameArabic}
}

func (thisObj *HunterStation) GetSearchTag() string {
	return "station"
}

func (thisObj *HunterStation) SetParentAssetInfo(parentAsset cervello.Asset) error {
	part, err := infrastructuredomain.MigrateCervelloAssetToArea(parentAsset)
	if err != nil {
		return errors.New("error fetching parent area: " + err.Error())
	}

	thisObj.CityId = part.CityId
	thisObj.CityNameEnglish = part.CityNameEnglish
	thisObj.CityNameArabic = part.CityNameArabic
	thisObj.AreaId = part.GlobalId
	thisObj.AreaNameEnglish = part.NameEnglish
	thisObj.AreaNameArabic = part.NameArabic
	thisObj.CityLayerId = part.CityLayerId
	thisObj.AreaLayerId = part.LayerId
	return nil
}

func (thisObj *HunterStation) SetParentGatewayInfo(parentDevice cervello.Device) error {
	panel, err := MigrateHunterControllerFromCervelloDevice(parentDevice)
	if err != nil {
		return errors.New("error fetching parent controller: " + err.Error())
	}
	thisObj.ControllerName = panel.Name
	thisObj.ControllerLayerId = panel.LayerId
	return nil
}

func (thisObj *HunterStation) GetLayerName() string {
	return thisObj.LayerName
}

func (thisObj *HunterStation) GetLayerId() float64 {
	return thisObj.LayerId
}

func (thisObj *HunterStation) GetLayerType() string {
	return thisObj.LayerType
}

func (thisObj *HunterStation) Validate() error {
	return crudfunctions.ValidateDeviceEntity(thisObj)
}

func (thisObj *HunterStation) MigrateFromCsvLine(csvLine []string, keysMap map[string]int) error {
	var err error

	thisObj.GlobalId = csvLine[keysMap["globalid"]]
	thisObj.IntegrationId = csvLine[keysMap["integrationuuid"]]
	thisObj.Name = csvLine[keysMap["name"]]
	thisObj.CityId = csvLine[keysMap["city_uuid"]]
	thisObj.AreaId = csvLine[keysMap["area_uuid"]]

	thisObj.X, err = strconv.ParseFloat(csvLine[keysMap["x"]], 64)
	if err != nil {
		return err
	}
	thisObj.Y, err = strconv.ParseFloat(csvLine[keysMap["y"]], 64)
	if err != nil {
		return err
	}
	thisObj.LayerId, err = strconv.ParseFloat(csvLine[keysMap["layerid"]], 64)
	if err != nil {
		return err
	}
	thisObj.LayerName = csvLine[keysMap["layername"]]
	thisObj.ControllerId = csvLine[keysMap["controller_uuid"]]
	thisObj.Brand = csvLine[keysMap["brand"]]
	thisObj.Model = csvLine[keysMap["modelno"]]
	thisObj.Type = "station"
	thisObj.LayerType = "point"
	thisObj.AdditionalInfo = common.SetupAdditionalInfo(keysMap, thisObj.GetEssentialKeys(), csvLine)

	return err
}

func (thisObj *HunterStation) GetEssentialKeys() []string {
	return []string{
		"id",
		"globalid",
		"integrationuuid",
		"name",
		"city_uuid",
		"area_uuid",
		"x",
		"y",
		"layerid",
		"layername",
		"controller_uuid",
		"brand",
		"modelno",
	}
}

func (thisObj *HunterStation) GetNonDuplicatingKeys() []string {
	return []string{
		"id",
		"globalid",
		"integrationuuid",
		"name",
	}
}

func (thisObj *HunterStation) GetParentAssetKey() string {
	return "area_uuid"
}

func (thisObj *HunterStation) GetParentGatewayKey() string {
	return "controller_uuid"
}
