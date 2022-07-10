package seeds



func (s Seed) CustomerSeed() {

	for i := 0; i < 100; i++ {
		//prepare the statement
		param1 := "param1"
		param2 := "param2"
		param3 := "param3"
		param4 := "param4"
		param5 := "param5"

		stmt, _ := s.db.Prepare(`INSERT INTO db (param1, param2, param3, param4, param5) VALUES (?,?,?,?,?)`)
		// execute query
		_, err := stmt.Exec(param1,param2, param3, param4, param5)
		if err != nil {
			panic(err)
		}
	}
}