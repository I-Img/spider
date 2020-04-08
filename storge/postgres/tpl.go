package postgres

const(
	selectContent = `SELECT id , name, auth, content, width, height FROM content ORDER BY id desc`
)
