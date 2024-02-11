package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
	// some values
}

type PostgresNumberStore struct {
	// postgres values (db connection)
}

// MongoDb tyoe implementation of the interface
func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mongoDB storage")
	return nil
}

// Postgres type implementation of the interfface
func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil
}

func (s PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the postgres storage")
	return nil
}

type ApiServer struct {
	mongoDBNumberStore    NumberStorer
	postgresDBNumberStore NumberStorer
}

func main() {
	// Creating API server which is using our MongoDb and Postgres Type
	apiServer := ApiServer{
		mongoDBNumberStore:    MongoDBNumberStore{},
		postgresDBNumberStore: PostgresNumberStore{},
	}
	// For MongoDb
	// Logic
	err := apiServer.mongoDBNumberStore.Put(1)
	if err != nil {
		panic(err)
	}
	// same code in short version
	if err := apiServer.mongoDBNumberStore.Put(1); err != nil {
		panic(err)
	}
	numbersMongo, err := apiServer.mongoDBNumberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbersMongo)

	// For PostGres
	// same code in short version
	if err := apiServer.postgresDBNumberStore.Put(1); err != nil {
		panic(err)
	}
	numbersPostgres, err := apiServer.postgresDBNumberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbersPostgres)
}
