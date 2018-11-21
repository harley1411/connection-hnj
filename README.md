"# connection-hnj
    Host your IP
    Set SetPgsql only first call
    connection.SetPgsql(connection.Pgsql{
		Host:"13.211.107.123123",
		Port:5432,
		User:"your_user",
		Password:"your_pw",
		Dbname:"db_name",
	})
	db :=connection.GetConnectionPsql()

"
