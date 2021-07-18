package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/addressBook/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	router    *httprouter.Router
	db        *sql.DB
	webServer *http.Server
	config    *config.ServerConfig
}

// New create a new server instance with specified config
func New(config *config.ServerConfig) *Server {

	router := httprouter.New()

	// automatic handle for OPTIONS requests
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Headers", strings.Join(config.CORS.AllowedHeaders, ", "))
			header.Set("Access-Control-Max-Age", strconv.FormatInt(config.CORS.MaxAge, 10))
		}
		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	webServer := &http.Server{
		Addr:         config.Listen,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		router:    router,
		webServer: webServer,
		config:    config,
	}
}

// Run starts the server
func (s *Server) Run() {

	// set DB
	dsn := s.config.DB.GetDSN()
	log.Println("trying to connect to database ...")

	db, err := sql.Open(s.config.DB.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("unable to ping database, err : %v", err))
	}

	_, err = db.Exec("drop table if exists addresses;")
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.
		Prepare(`create table addresses(id int UNSIGNED AUTO_INCREMENT PRIMARY KEY, 
								fname VARCHAR(30) NOT NULL,
								lname VARCHAR(30) NOT NULL,
								phone_number VARCHAR(30));`)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	log.Println("database table 'addresses' created ....")

	_, err = db.Exec("insert into addresses (id,fname,lname,phone_number) values(1,'test','test','123456789')")
	if err != nil {
		panic(err)
	}

	s.db = db

	// init components
	s.registerComponents()

	// start web server
	log.Printf("Address Book API server started, listening => %s\n", s.config.Listen)

	if err := s.webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[ERROR] unable to start server, err : %v", err)
	}

}

// Stop stops the server and close all connections (DB)
func (s *Server) Stop(ctx context.Context) {

	if err := s.webServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	if err := s.db.Close(); err != nil {
		log.Printf("[ERROR] unable to disconnect database, err : %v\n", err)
	}

	log.Println("[DEBUG] server stopped gracefully ...")
}
