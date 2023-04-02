package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"

	"gitlab.com/raihanlh/messenger-api/config"
	"gitlab.com/raihanlh/messenger-api/internal/constant"
	"gitlab.com/raihanlh/messenger-api/internal/model"
	"gitlab.com/raihanlh/messenger-api/pkg/postgres"
	"gorm.io/gorm"
)

type Seed struct {
	db *gorm.DB
}

func (s Seed) UserSeed() {
	fpath, err := os.Open(GetSourcePath() + "/json/users.json")
	if err != nil {
		log.Printf("cannot open file, err: %s", err)
	}
	defer fpath.Close()

	data, err := ioutil.ReadAll(fpath)
	if err != nil {
		log.Printf("cannot read file, err: %s", err)
	}

	var users []*model.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Printf("cannot unmarshal json, err: %s", err)
	}

	for _, user := range users {
		result := s.db.Debug().Table(constant.UserTable).Create(&user)
		if err := result.Error; err != nil {
			log.Printf("cannot seed apps table: %v", err)
		}
	}
}

func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// seed will run a selected method of Seed struct
// Example:
//      seed(s, PlatformSeed)
func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	method := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !method.IsValid() {
		log.Fatalf("No method named %s", seedMethodName)
	}

	// Execute method
	log.Println("Seeding", seedMethodName, "...")
	method.Call(nil)
	log.Println("Seed", seedMethodName, "finished")
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get method in current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}
	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

// Seed package is used to pre-fill certain data that is mostly static / rarely changes
// Usage: go run cmd/seeder/seeder.go UserSeed TestSeed
func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	conf := config.New()

	db := postgres.New(conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBTimezone)

	Execute(db, args...)
	os.Exit(0)
}
