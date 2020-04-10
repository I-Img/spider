package postgres

import "time"

func (d *Driver) FetchPx500LastContent() (time.Time, error) {
	var t time.Time

	rows, err := d.cli.Query(`SELECT create_time FROM px500pic ORDER  BY id DESC LIMIT 1;`)
	if err != nil {
		return t, err
	}

	for rows.Next() {
		if err = rows.Scan(&t); err != nil {
			return t, err
		}
	}

	return time.Now().Add(time.Hour * -1), nil
}
