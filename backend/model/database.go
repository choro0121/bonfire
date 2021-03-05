package model

import (
    "os"
    "log"
    "fmt"

    "github.com/lib/pq"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var db *gorm.DB

func New() {
    log.Print("initialize postgres")

    var err error

    // connect postgres
    connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }

    // open
    db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // migrate
    db.Migrator().DropTable(
        &TempEmailUser{},
        &TempOAuthUser{},
        &User{},
        &Follow{},
        &OAuth{},
        &Password{},
        &Post{},
        &Bookmark{},
        &Good{},
        &Tag{},
        &Comment{},
    )
    db.AutoMigrate(&TempEmailUser{})
    db.AutoMigrate(&TempOAuthUser{})
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Follow{})
    db.AutoMigrate(&OAuth{})
    db.AutoMigrate(&Password{})
    db.AutoMigrate(&Post{})
    db.AutoMigrate(&Bookmark{})
    db.AutoMigrate(&Good{})
    db.AutoMigrate(&Tag{})
    db.AutoMigrate(&Comment{})
}

func Close() error {
    log.Print("close postgres")

    sql, err := db.DB()

    if err != nil {
        return err
    }
    
    return sql.Close()
}

func CreateMany2ManyWidhoutDuplicate(data interface{}) (interface{}, error) {
    var count int64
    db.Model(data).Where(data).Count(&count)

    if count > 0 {
        return nil, fmt.Errorf("already created %v", data)
    } else {
        // create
        return data, db.Create(data).Error
    }
}
