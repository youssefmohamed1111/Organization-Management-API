package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type User struct{
	ID int		 `json:"id"`
	Name 	string	`json:"name"`
	Email 	string `json:"email"`
}	
func main(){
	// connect to database
		// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		db := options.Client().ApplyURI("mongodb+srv://Youssef:<password>@cluster0.wotxocw.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
		// Create a new client and connect to the server
		client, err := mongo.Connect(context.TODO(), db)
		if err != nil {
		  panic(err)
		}
		defer func() {
		  if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		  }
		}()
		// Send a ping to confirm a successful connection
		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		  panic(err)
		}
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	//create router
	router := mux.NewRouter()
	router.HandleFunc("/users",getUser(db)).Methods("GET")
	router.HandleFunc("/users/{id}",getUser(db)).Methods("GET")
	router.HandleFunc("/users/{id}",getUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}",getUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}",getUser(db)).Methods("DELETE")

	//Start Server
	log.fatal(http.ListenAndServer(":8000", jsonContentTypeMiddleware(router)))
	}		
	func jsonContentTypeMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
	
	// get all users
	func getUsers(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			rows, err := db.Query("SELECT * FROM users")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
	
			users := []User{}
			for rows.Next() {
				var u User
				if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
					log.Fatal(err)
				}
				users = append(users, u)
			}
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
	
			json.NewEncoder(w).Encode(users)
		}
	}
	
	// get user by id
	func getUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]
	
			var u User
			err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
	
			json.NewEncoder(w).Encode(u)
		}
	}
	
	// create user
	func createUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var u User
			json.NewDecoder(r.Body).Decode(&u)
	
			err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
			if err != nil {
				log.Fatal(err)
			}
	
			json.NewEncoder(w).Encode(u)
		}
	}
	
	// update user
	func updateUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var u User
			json.NewDecoder(r.Body).Decode(&u)
	
			vars := mux.Vars(r)
			id := vars["id"]
	
			_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
			if err != nil {
				log.Fatal(err)
			}
	
			json.NewEncoder(w).Encode(u)
		}
	}
	
	// delete user
	func deleteUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]
	
			var u User
			err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
				if err != nil {
					//todo : fix error handling
					w.WriteHeader(http.StatusNotFound)
					return
				}
		
				json.NewEncoder(w).Encode("User deleted")
			}
		}
	}