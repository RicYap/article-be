package boot

import (
	"article/docs"
	"log"
	"net/http"

	"article/internal/config"
	"article/internal/entity/article"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	articleData "article/internal/data/article"
	articleServer "article/internal/delivery/http"
	articleHandler "article/internal/delivery/http/article"
	articleService "article/internal/service/article"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg := config.Get()

	// Open Databases
	// db, db2, err := openDatabases(cfg)
	db, err := gorm.Open(mysql.Open(cfg.Database.Master), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = db.AutoMigrate(&article.Post{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	// Diganti dengan domain yang anda buat
	ad := articleData.New(db)
	as := articleService.New(ad)
	ah := articleHandler.New(as)

	s := articleServer.Server{
		Article: ah,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
