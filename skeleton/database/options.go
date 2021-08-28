package database

var defaultOptions = options{
	dialect:     "mysql",
	autoMigrate: true,
	debug:       true,

	mysqlOptions:  defaultMysqlOptions,
	sqliteOptions: defaultSqliteOptions,
}
var defaultMysqlOptions = mysqlOptions{
	user:     "test",
	password: "test",
	host:     "localhost",
	port:     3306,
	name:     "test",
	charset:  "utf8mb4",
}
var defaultSqliteOptions = sqliteOptions{
	name: "test.db",
}

type mysqlOptions struct {
	user     string `yaml:"user"`
	password string `yaml:"password"`
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	name     string `yaml:"name"`
	charset  string `yaml:"charset"`
}

type sqliteOptions struct {
	name string `yaml:"name"`
}

type options struct {
	dialect     string `yaml:"dialect"`
	autoMigrate bool   `yaml:"autoMigrate"`
	debug       bool

	mysqlOptions  mysqlOptions  `yaml:"mysql" mapstructure:"mysql"`
	sqliteOptions sqliteOptions `yaml:"sqlite" mapstructure:"sqlite"`
}

// Option 定义参数项
type Option func(*options)
type MysqlOption func(*mysqlOptions)
type SqliteOption func(*sqliteOptions)

//WithDialect 设定 dialect
func WithDialect(dialect string) Option {
	return func(o *options) {
		o.dialect = dialect
	}
}

//WithAutoMigrate 设定 autoMigrate
func WithAutoMigrate(autoMigrate bool) Option {
	return func(o *options) {
		o.autoMigrate = autoMigrate
	}
}

//WithDebug 设定 debug
func WithDebug(debug bool) Option {
	return func(o *options) {
		o.debug = debug
	}
}

//WithMysqlOptions 设定 mysqlOptions
func WithMysqlOptions(opts ...MysqlOption) Option {
	mysqlOptions := defaultMysqlOptions
	for _, opt := range opts {
		opt(&mysqlOptions)
	}
	return func(o *options) {
		o.mysqlOptions = mysqlOptions
	}
}

//WithSqliteOptions 设定 sqliteOptions
func WithSqliteOptions(opts ...SqliteOption) Option {
	sqliteOptions := defaultSqliteOptions
	for _, opt := range opts {
		opt(&sqliteOptions)
	}
	return func(o *options) {
		o.sqliteOptions = sqliteOptions
	}
}

//WithMysqlUser 设定 user
func WithMysqlUser(user string) Option {
	return func(o *options) {
		o.mysqlOptions.user = user
	}
}
func WithMysqlPassword(password string) Option {
	return func(o *options) {
		o.mysqlOptions.password = password
	}
}
func WithMysqlHost(host string) Option {
	return func(o *options) {
		o.mysqlOptions.host = host
	}
}
func WithMysqlPort(port int) Option {
	return func(o *options) {
		o.mysqlOptions.port = port
	}
}
func WithMysqlName(name string) Option {
	return func(o *options) {
		o.mysqlOptions.name = name
	}
}
func WithMysqlCharset(charset string) Option {
	return func(o *options) {
		o.mysqlOptions.charset = charset
	}
}
func WithSqliteName(name string) Option {
	return func(o *options) {
		o.sqliteOptions.name = name
	}
}
