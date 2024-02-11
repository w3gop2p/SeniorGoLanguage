package main

// some issues with pointers

type Database struct {
	user string
}

type Server1 struct {
	db *Database
}

func (s *Server1) GetUserFromDb() string {
	// golang is going to "dereference" the db pointer
	// its going to lookup the memory address of the pointer, but the address is null, is nothing there
	// so we can check in order to get more information about what is going on
	if s.db == nil {
		panic("database == nil hence, is not initialized")
	}
	return s.db.user
}

func main() {
	q := &Server1{} // in this case it will be an error
	q.GetUserFromDb()

	// like this it will work
	/*s := &Server1{
		db: &Database{},
	}
	s.GetUserFromDb()*/
}
