package database

// Cannot perform SET, GET, or DELETE operation on DB
type OperationError struct {
	operation string
}

func (err *OperationError) Error() string {
	return "Could not perform: " + err.operation
}

// DownError => when res is not redis.Nil => db is down
type DownError struct{}

func (dbe *DownError) Error() string {
	return "DB is down"
}

// Cannot perform set on DB
type CreateDatabaseError struct{}

func (err *CreateDatabaseError) Error() string {
	return "Could not create database"
}

// When creating not implemented DB
type NotImplementedDatabaseError struct {
	database string
}

func (err *NotImplementedDatabaseError) Error() string {
	return err.database + " not implemented"
}
