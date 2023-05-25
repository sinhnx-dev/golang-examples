// package models

// import (
// 	"database/sql"
// 	"errors"
// 	"time"
// )

// type User struct {
// 	ID           uint
// 	Name         string
// 	Email        *string
// 	Age          uint8
// 	Birthday     *time.Time
// 	MemberNumber sql.NullString
// 	ActivatedAt  sql.NullTime
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// }

// // gorm.Model definition
// type Model struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// type User struct {
// 	Name string `gorm:"<-:create"`          // allow read and create
// 	Name string `gorm:"<-:update"`          // allow read and update
// 	Name string `gorm:"<-"`                 // allow read and write (create and update)
// 	Name string `gorm:"<-:false"`           // allow read, disable write permission
// 	Name string `gorm:"->"`                 // readonly (disable write permission unless it configured)
// 	Name string `gorm:"->;<-:create"`       // allow read and create
// 	Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
// 	Name string `gorm:"-"`                  // ignore this field when write and read with struct
// 	Name string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
// 	Name string `gorm:"-:migration"`        // ignore this field when migrate with struct
// }

// type User struct {
// 	CreatedAt time.Time // Set to current time if it is zero on creating
// 	UpdatedAt int       // Set to current unix seconds on updating or if it is zero on creating
// 	Updated   int64     `gorm:"autoUpdateTime:nano"`  // Use unix nano seconds as updating time
// 	Updated   int64     `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
// 	Created   int64     `gorm:"autoCreateTime"`       // Use unix seconds as creating time
// }

// type User struct {
// 	gorm.Model
// 	Name string
// }

// // equals
// type User struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Name      string
// }

// type Author struct {
// 	Name  string
// 	Email string
// }

// type Blog struct {
// 	ID      int
// 	Author  Author `gorm:"embedded"`
// 	Upvotes int32
// }

// // equals
// type Blog struct {
// 	ID      int64
// 	Name    string
// 	Email   string
// 	Upvotes int32
// }

// type Blog struct {
// 	ID      int
// 	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
// 	Upvotes int32
// }

// // equals
// type Blog struct {
// 	ID          int64
// 	AuthorName  string
// 	AuthorEmail string
// 	Upvotes     int32
// }

// func CreateRecord() {
// 	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

// 	// result := db.Create(&user) // pass pointer of data to Create

// 	// user.ID             // returns inserted data's primary key
// 	// result.Error        // returns error
// 	// result.RowsAffected // returns inserted records count

// 	// users := []*User{
// 	// 	User{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
// 	// 	User{Name: "Jackson", Age: 19, Birthday: time.Now()},
// 	// }

// 	// result := db.Create(users) // pass a slice to insert multiple row

// 	// result.Error        // returns error
// 	// result.RowsAffected // returns inserted records count

// 	db.Select("Name", "Age", "CreatedAt").Create(&user)
// 	// INSERT INTO `users` (`name`,`age`,`created_at`)
// 	// VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

// 	db.Omit("Name", "Age", "CreatedAt").Create(&user)
// 	// INSERT INTO `users` (`birthday`,`updated_at`)
// 	// VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

// 	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
// 	db.Create(&users)

// 	for _, user := range users {
// 		user.ID // 1,2,3
// 	}

// 	// 	var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}

// 	// // batch size 100
// 	// db.CreateInBatches(users, 100)

// 	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
// 	// 	CreateBatchSize: 1000,
// 	//   })

// 	//   db := db.Session(&gorm.Session{CreateBatchSize: 1000})

// 	//   users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}

// 	//   db.Create(&users)
// 	//   // INSERT INTO users xxx (5 batches)
// 	//   // INSERT INTO pets xxx (15 batches)

// 	// Get the first record ordered by primary key
// 	db.First(&user)
// 	// SELECT * FROM users ORDER BY id LIMIT 1;

// 	// Get one record, no specified order
// 	db.Take(&user)
// 	// SELECT * FROM users LIMIT 1;

// 	// Get last record, ordered by primary key desc
// 	db.Last(&user)
// 	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 	result := db.First(&user)
// 	result.RowsAffected // returns count of records found
// 	result.Error        // returns error or nil

// 	// check error ErrRecordNotFound
// 	errors.Is(result.Error, gorm.ErrRecordNotFound)

// 	db.First(&user, 10)
// 	// SELECT * FROM users WHERE id = 10;

// 	db.First(&user, "10")
// 	// SELECT * FROM users WHERE id = 10;

// 	db.Find(&users, []int{1, 2, 3})
// 	// SELECT * FROM users WHERE id IN (1,2,3);

// 	db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// 	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

// 	var user = User{ID: 10}
// 	db.First(&user)
// 	// SELECT * FROM users WHERE id = 10;

// 	var result User
// 	db.Model(User{ID: 10}).First(&result)
// 	// SELECT * FROM users WHERE id = 10;

// 	// Get all records
// 	result := db.Find(&users)
// 	// SELECT * FROM users;

// 	result.RowsAffected // returns found records count, equals `len(users)`
// 	result.Error        // returns error

// 	// Get first matched record
// 	db.Where("name = ?", "jinzhu").First(&user)
// 	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// 	// Get all matched records
// 	db.Where("name <> ?", "jinzhu").Find(&users)
// 	// SELECT * FROM users WHERE name <> 'jinzhu';

// 	// IN
// 	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// 	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// 	// LIKE
// 	db.Where("name LIKE ?", "%jin%").Find(&users)
// 	// SELECT * FROM users WHERE name LIKE '%jin%';

// 	// AND
// 	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// 	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// 	// Time
// 	db.Where("updated_at > ?", lastWeek).Find(&users)
// 	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// 	// BETWEEN
// 	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// 	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

// 	// Struct
// 	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// 	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

// 	// Map
// 	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// 	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// 	// Slice of primary keys
// 	db.Where([]int64{20, 21, 22}).Find(&users)
// 	// SELECT * FROM users WHERE id IN (20, 21, 22);

// 	db.Select("name", "age").Find(&users)
// 	// SELECT name, age FROM users;

// 	db.Select([]string{"name", "age"}).Find(&users)
// 	// SELECT name, age FROM users;

// 	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// 	// SELECT COALESCE(age,'42') FROM users;

// 	db.Order("age desc, name").Find(&users)
// 	// SELECT * FROM users ORDER BY age desc, name;

// 	// Multiple orders
// 	db.Order("age desc").Order("name").Find(&users)
// 	// SELECT * FROM users ORDER BY age desc, name;

// 	db.Clauses(clause.OrderBy{
// 		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
// 	}).Find(&User{})
// 	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)

// 	db.Limit(3).Find(&users)
// 	// SELECT * FROM users LIMIT 3;

// 	// Cancel limit condition with -1
// 	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// 	// SELECT * FROM users LIMIT 10; (users1)
// 	// SELECT * FROM users; (users2)

// 	db.Offset(3).Find(&users)
// 	// SELECT * FROM users OFFSET 3;

// 	db.Limit(10).Offset(5).Find(&users)
// 	// SELECT * FROM users OFFSET 5 LIMIT 10;

// 	// Cancel offset condition with -1
// 	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// 	// SELECT * FROM users OFFSET 10; (users1)
// 	// SELECT * FROM users; (users2)

// 	// type result struct {
// 	// 	Date  time.Time
// 	// 	Total int
// 	//   }

// 	//   db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
// 	//   // SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1

// 	//   db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
// 	//   // SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

// 	//   rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
// 	//   defer rows.Close()
// 	//   for rows.Next() {
// 	// 	...
// 	//   }

// 	//   rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
// 	//   defer rows.Close()
// 	//   for rows.Next() {
// 	// 	...
// 	//   }

// 	//   type Result struct {
// 	// 	Date  time.Time
// 	// 	Total int64
// 	//   }
// 	//   db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

// 	db.First(&user)

// 	user.Name = "jinzhu 2"
// 	user.Age = 100
// 	db.Save(&user)
// 	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01',
// 	// updated_at = '2013-11-17 21:34:10' WHERE id=111;

// 	db.Save(&User{Name: "jinzhu", Age: 100})
// 	// INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES
// 	// ("jinzhu",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")

// 	db.Save(&User{ID: 1, Name: "jinzhu", Age: 100})
// 	// UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",
// 	// `update_at`="0000-00-00 00:00:00" WHERE `id` = 1

// 	// Email's ID is `10`
// 	db.Delete(&email)
// 	// DELETE from emails where id = 10;

// 	// Delete with additional conditions
// 	db.Where("name = ?", "jinzhu").Delete(&email)
// 	// DELETE from emails where id = 10 AND name = "jinzhu";

// 	db.Delete(&User{}, 10)
// 	// DELETE FROM users WHERE id = 10;

// 	db.Delete(&User{}, "10")
// 	// DELETE FROM users WHERE id = 10;

// 	db.Delete(&users, []int{1, 2, 3})
// 	// DELETE FROM users WHERE id IN (1,2,3);

// }
