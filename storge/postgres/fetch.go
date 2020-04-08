package postgres

import "spider/common"

func (d *Driver) FetchPictures(pos int)(content []common.Content, err error){
	row, err := d.cli.Query(selectContent)
	if err != nil{
		return
	}

	for row.Next(){
		var c common.Content
		if err = row.Scan(&c.ID, &c.Name, &c.Auth, &c.Content, &c.Width,&c.Height); err != nil{
			return
		}

		content = append(content, c)
	}

	return
}
