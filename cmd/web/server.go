package main

//func main() {
//
//	db := PoolCreate()
//	defer db.pool.Close()
//
//	//connStr := "host=localhost port=5432 user=postgres password=15975300 dbname=save_pass sslmode=disable"
//	//db, err := sql.Open("postgres", connStr)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer db.Close()
//	//
//	//var (
//	//	id   int
//	//	name string
//	//)
//	//
//	//rows, err := db.Query("select id, name from users")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer rows.Close()
//	//
//	//for rows.Next() {
//	//	err := rows.Scan(&id, &name)
//	//	if err != nil {
//	//		log.Fatal(err)
//	//	}
//	//	log.Println(id, name)
//	//}
//
//	//err = db.Ping()
//	//if err != nil {
//	//	fmt.Println("databaseConnection")
//	//}
//
//	e := echo.New()
//
//	router(e)
//
//	if err := e.Start(":3125"); err != http.ErrServerClosed {
//		e.Logger.Fatal(err)
//	}
//}
