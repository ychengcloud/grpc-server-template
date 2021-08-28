package database

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/mysql"
	"gorm.io/gorm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	testopts = []Option{
		WithDebug(true),
		WithMysqlName("account"),
		WithMysqlHost("127.0.0.1"),
		WithMysqlPort(13306),
	}
)

func RunContainerForTest(ctx context.Context, queriesFile string) (*gnomock.Container, error) {
	o := defaultOptions
	for _, opt := range testopts {
		opt(&o)
	}

	p := mysql.Preset(
		mysql.WithUser(o.mysqlOptions.user, o.mysqlOptions.password),
		mysql.WithDatabase(o.mysqlOptions.name),
		mysql.WithQueriesFile(queriesFile),
	)
	path, _ := os.Getwd()
	fmt.Println("Path:", path)
	return gnomock.Start(p)
	// return gnomock.Start(p, gnomock.WithDebugMode())

}

func ConnectForTest(ctx context.Context, addr string) (db *gorm.DB, err error) {
	address := strings.Split(addr, ":")
	port, _ := strconv.Atoi(address[1])
	var opts []Option
	opts = append(opts, testopts...)
	opts = append(
		opts,
		WithMysqlHost(address[0]),
		WithMysqlPort(port),
	)
	return NewGorm(ctx, opts...)
}
