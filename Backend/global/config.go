package global

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
	Database string `mapstructure:"database"`
}

func (m *MySQLConfig) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + "charset=utf8&parseTime=true"
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"DB"`
}

type COSConfig struct {
	Url       string `mapstructure:"url"`
	SecretID  string `mapstructure:"SecretID"`
	SecretKey string `mapstructure:"SecretKey"`
}

type RSAConfig struct {
	PublicKey  string `mapstructure:"PublicKey"`
	PrivateKey string `mapstructure:"PrivateKey"`
}
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}
type DefaultAvatar struct {
	UserUrl  string `mapstructure:"userUrl"`
	GroupUrl string `mapstructure:"groupUrl"`
}
type JWT struct {
	SecretKey  string `mapstructure:"SecretKey"`
	ExpireTime int    `mapstructure:"ExpireTime"`
}
type Config struct {
	MySQL  MySQLConfig   `mapstructure:"Mysql"`
	Redis  RedisConfig   `mapstructure:"Redis"`
	COS    COSConfig     `mapstructure:"COS"`
	RSA    RSAConfig     `mapstructure:"RSA"`
	Zap    Zap           `mapstructure:"Zap"`
	Avatar DefaultAvatar `mapstructure:"DefaultAvatar"`
	JWT    JWT           `mapstructure:"JWT"`
}

//var Server Config

//func ReadIn() {
//	viper.SetConfigName("global")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath("./global")
//	if err := viper.ReadInConfig(); err != nil {
//		log.Printf("global-配置信息读取出错，%v\n", err.Error())
//	}
//	if err := viper.Unmarshal(&Server); err != nil {
//		log.Printf("global-配置信息解码出错，%v\n", err.Error())
//	}
//	log.Printf("[info]读取到的配置信息：%v\n", Server)
//
//}
//func init() {
//	ReadIn()
//}
