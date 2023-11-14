package config

type Config struct {
	Service `json:"service"`
	DB      `json:"db"`
	AWS     `json:"aws"`
}

type Service struct {
	Name        string      `envconfig:"SERVICE_NAME" json:"name"`
	Version     string      `envconfig:"SERVICE_VERSION" json:"version"`
	Environment Environment `envconfig:"ENVIRONMENT" json:"environment"`
	StaticDir   string      `envconfig:"STATIC_DIR" json:"static_dir"`
}

type DB struct {
	PrimaryHost string `envconfig:"DATABASE_PRIMARY_HOST" json:"primary_host" default:"localhost"`
	PrimaryPort int    `envconfig:"DATABASE_PRIMARY_PORT" json:"primary_port" default:"3306"`
	ReplicaHost string `envconfig:"DATABASE_REPLICA_HOST" json:"replica_host" default:"localhost"`
	ReplicaPort int    `envconfig:"DATABASE_REPLICA_PORT" json:"replica_port" default:"3306"`
	Database    string `envconfig:"DATABASE_NAME" json:"database" default:"line-go-sample"`
	User        string `envconfig:"DATABASE_USER" json:"user" default:"user"`
	Password    string `envconfig:"DATABASE_PASSWORD" json:"password" default:"P@ssw0rd"`
}

type AWS struct {
}
