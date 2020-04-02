package postgres

import (
	"github.com/lib/pq"
	"spider/common"
	"spider/storge"
)

func (d *Driver) SaveContents(content []common.Content) error {
	txn, err := d.cli.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn(storge.CONTENT, "name", "auth", "crawl_time", "content", "c_type", "width", "height", "remarks", "id"))
	if err != nil {
		return err
	}

	for _, c := range content {
		_, err = stmt.Exec(c.Name, c.Auth, c.CrawlTime, c.Content, c.CType, c.Width, c.Height, c.Remarks, d.pos)
		if err != nil {
			return err
		}

		d.pos++
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}
	return nil
}
