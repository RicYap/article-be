package boot

import (
	"log"
	"net/http"
	"skeleton/docs"

	"skeleton/internal/config"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	skeletonData "skeleton/internal/data/skeleton"
	skeletonServer "skeleton/internal/delivery/http"
	skeletonHandler "skeleton/internal/delivery/http/skeleton"
	skeletonService "skeleton/internal/service/skeleton"
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
	db, err := openDatabases(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	// Diganti dengan domain yang anda buat
	sd := skeletonData.New(db)
	ss := skeletonService.New(sd)
	sh := skeletonHandler.New(ss)

	//watch config changes
	config.PrepareWatchPath()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := config.Init()
		if err != nil {
			log.Printf("[VIPER] Error get config file, %v", err)
		}
		cfg = config.Get()

		//open new db connection pool
		//dbNew, db2New, err := openDatabases(cfg)
		dbNew, err := openDatabases(cfg)
		if err != nil {
			log.Printf("[VIPER] Error open db connection, %v", err)
		} else {
			//replace all previous db connection pool
			//*db2 = *db2New
			*db = *dbNew

			//re-init all Data Layer
			//sd2.InitStmt()
			sd.InitStmt()
		}
	})

	s := skeletonServer.Server{
		Skeleton: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

//open all databases here
func openDatabases(cfg *config.Config) (db *sqlx.DB /*db2 *sqlx.DB,*/, err error) {
	db, err = openConnectionPool("mysql", cfg.Database.Master)
	if err != nil {
		return db, err
	}

	// db2, err = openConnectionPool("mysql", cfg.Database.DB2)
	// if err != nil {
	// 	return db, db2, err
	// }

	return db, err

	//return db, db2, err
}

//create new connection pool and test the connection
func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(driver, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
