package config

type Mysql struct {
	Dbname string `mapstructure:"db-name" json"dbname" yaml:"db-name"`
}
