package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/zaza-hikayat/sample-api/lib"
)

func main() {

	// load config
	conf := lib.LoadConfig()
	db := lib.GetConnectionDB(conf)

	r := chi.NewRouter()
	app := &lib.App{DB: db}
	ctrl := lib.NewController(app)

	// option command
	migrateOpt := flag.String("db", "", "option for migration and seed data")
	flag.Parse()
	if migrateOpt != nil && *migrateOpt == "migrate" {
		if err := lib.RunMigration(db); err != nil {
			log.Fatal("failed migrate db")
		}
		log.Println("success migrate database")
		lib.RunSeeder(db)
		// os.Exit(0)
		// return

	}

	// setup controller
	r.Get("/api/v1/sequence", ctrl.PrintData)
	r.Get("/api/v1/fizzbuzz", ctrl.PrintFizzbuzz)
	r.Get("/api/v1/contact-chip", ctrl.ContactChip)
	r.Get("/api/v1/mahasiswa", ctrl.DataMahasiswa)
	r.Get("/api/v1/dosen", ctrl.DataDosen)
	r.Get("/api/v1/matakuliah", ctrl.DataMataKuliah)
	r.Get("/api/v1/nilai", ctrl.DataNilai)
	r.Post("/api/v1/show-json", ctrl.InsertData)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("views/assets"))))
	srv := http.Server{
		Addr:    ":" + conf.AppPort,
		Handler: r,
	}
	// run http server
	HttpServe(&srv)
}

func HttpServe(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	go func() {
		log.Println("server start running at IP", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("server cannot run ", err)
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	s := <-quit
	log.Printf("got signal: %v, shutting down server ...\n", s)

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
