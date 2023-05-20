package config
import ("gorm.io/gorm")

var DB*gotm.DB

fucn Connect(){
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost/user_management_api?ssslmode=disable"),&gorm.Config)


	if err != nil { panic(err)
	}
	sqlDB, err := db.DB() if err != nil { panic(err)
	}
	start:=  time.Now()
	for sqlDB.Ping() != nil {
		if start.After (start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect DB after 10 seconds")
			break
	}
	
	fmt.Println("Connected to DB: ", sqlDB.Ping() == nil)

	db.AutoMigrate(&models.User{})

}