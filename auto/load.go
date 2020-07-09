package auto

import (
  "didit/api/database"
  "didit/api/database"
  "log"
)

func Load() {
  db, err := database.Connect()
  if err != nil {
    log.Fatalf("Cannot connect to database: %v", err)
  }
  defer db.Close()
  err = dbDebug().DropTableIfExists(&models.Post{}, &models.User{}).Error
  if err != nil {
    log.FatalF("Cannot drop table: %v", err)
  }
  err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
  if err != nil {
    log.Fatalf("Cannot drop table: %v", err)
  }
  err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
  if err != nil {
    log.Fatalf("Attaching foreing key error: %v", err)
  }
}
