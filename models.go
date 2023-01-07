package config

import "time"

type Redis struct {
	Server           string `json:"server"`
	Port             int    `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         int    `json:"database"`
	PlayBackDatabase int    `json:"play_back_database"`
}

type MongoServer struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

type Mongo struct {
	Servers          []MongoServer `json:"servers"`
	Username         string        `json:"username"`
	Password         string        `json:"password"`
	AuthSource       string        `json:"auth_source"`
	ReplicaSet       string        `json:"replica_set"`
	ReadPref         string        `json:"read_pref"`
	PlayBackDatabase string        `json:"play_back_database"`
	AlarmDatabase    string        `json:"alarm_database"`
	Database         string        `json:"database"`
	AlarmTable       string        `json:"alarm_table"`
}

type MySQL struct {
	Server       string        `json:"server"`
	Port         int           `json:"port"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	Database     string        `json:"database"`
	MaxIdleConns int           `json:"max_idle_conns"`
	MaxOpenConns int           `json:"max_open_conns"`
	MaxLifeTime  time.Duration `json:"max_life_time"`
	MaxIdleTime  time.Duration `json:"max_idle_time"`
}

type Queue struct {
	Notification []string `json:"notification"`
	Location     []string `json:"location"`
}

type Exchange struct {
	Notification string `json:"notification"`
	Location     string `json:"location"`
}

type RabbitMQ struct {
	Server            string   `json:"server"`
	Port              int      `json:"port"`
	Username          string   `json:"username"`
	Password          string   `json:"password"`
	NotificationQueue string   `json:"notification_queue"`
	AlarmQueue        string   `json:"alarm_queue"`
	LocationQueue     string   `json:"location_queue"`
	IsEnable          bool     `json:"is_enable"`
	Queue             Queue    `json:"queue"`
	Exchange          Exchange `json:"exchange"`
}

type Metric struct {
	X3       int `json:"x3"`
	Gt06     int `json:"gt06"`
	Parser   int `json:"parser"`
	Firebase int `json:"firebase"`
	Syncer   int `json:"syncer"`
	Location int `json:"location"`
	Alarm    int `json:"alarm"`
}

type ParserTCP struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
}

type GRPC struct {
	ParserTcp ParserTCP `json:"parser_tcp"`
}

type Port struct {
	X3      int `json:"x3"`
	Gt06    int `json:"gt06"`
	Xenotic int `json:"xenotic"`
	Cooban  int `json:"cooban"`
}

type Server struct {
	Ip   string `json:"ip"`
	Port Port   `json:"port"`
}
