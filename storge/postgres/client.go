package postgres

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/BurntSushi/toml"

	_ "github.com/lib/pq"
)

type Driver struct {
	cli *sql.DB
	pos int
}

type Config struct {
	Addr   string `toml:"addr"`
	User   string `toml:"user"`
	DB     string `toml:"db"`
	Passwd string `toml:"passwd"`
}

func NewDriver() (*Driver, error) {
	c := Config{}

	_, err := toml.DecodeFile(os.Getenv("POSTGRES_CONFIG"), &c)
	if err != nil {
		return nil, fmt.Errorf("DecodeFile Error[%s]", err.Error())
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, c.Passwd, c.Addr, c.DB)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	d := &Driver{
		cli: db,
	}

	if err := d.getCurrentPos(); err != nil {
		return nil, err
	}

	d.pos += 1

	return d, nil
}

func (d *Driver) getCurrentPos() error {
	row, err := d.cli.Query(`select count(*) from content`)
	if err != nil {
		return err
	}

	defer row.Close()

	var i int
	for row.Next() {
		row.Scan(&i)
	}

	d.pos = i
	return nil
}
