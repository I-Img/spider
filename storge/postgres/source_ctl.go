package postgres

import (
	"fmt"

	"spider/common"
)

func (d *Driver) FetchSource(name string) (source common.Source, err error) {
	rows, err := d.cli.Query(`SELECT  id, name, last_time, pid FROM source_ctl WHERE name=$1`, name)
	if err != nil {
		return
	}

	if rows.Next() {
		if err = rows.Scan(&source.ID, &source.Name, &source.LastTime, &source.PID); err != nil {
			return
		}

		return
	}

	return source, fmt.Errorf("Can not find [%s] source ", name)
}

func (d *Driver) UpdateSource(s common.Source) error {
	if _, err := d.cli.Query(`UPDATE source_ctl SET last_time=$1, pid =$2 WHERE name = $3`, s.LastTime, s.PID, s.Name); err != nil {
		return err
	}

	return nil
}
