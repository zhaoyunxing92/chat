package domain

type Employee struct {

	// 用户名字
	Name string `json:"name"`

	// 用户unionId
	UnionId string `json:"unionId"`

	// 用户的userid
	UserId string `json:"userid"`

	//级别。
	//
	//1：主管理员
	//
	//2：子管理员
	//
	//100：老板
	//
	//0：其他（如普通员工）
	Level int `json:"level"`

	//是否是管理员。
	//
	//true：是
	//
	//false：不是
	Admin bool `json:"admin"`

	// 设备ID
	DeviceId string `json:"deviceId"`
}
