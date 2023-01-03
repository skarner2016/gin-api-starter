package mysql

type Instance string

type DBConf struct {
	Host            string
	Port            int64
	User            string
	Pass            string
	Database        string
	Charset         string
	Collation       string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

const InstanceDefault Instance = "default"
