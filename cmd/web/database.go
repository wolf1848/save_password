package main

//type PSQL struct {
//	pool *pgxpool.Pool
//}
//
//func PoolCreate() (db *PSQL) {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatalf("Error loading .env file: %v\n", err)
//	}
//
//	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
//	if err != nil {
//		log.Fatalf("Unable to connect to database: %v\n", err)
//	}
//
//	db = &PSQL{
//		pool: pool,
//	}
//
//	return
//}
