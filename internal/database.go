package database

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres123"
	dbname   = "docker-test"
)

type Customer struct {
	Id int,
	FirstName string,
	LastName string
}

func ConnectDB() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s " +
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	Database, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	// Validate database connection
	err = Database.Ping()
	if err != nil {
		panic(err)
	}
}

func getCustomer(id int) (Customer, error) {
		query := "select first_name, last_name" +
		"       	from customer " +
		"     	   where id = $1"

	res, err := db.QueryRow(query, id)

	if err != nil {
		return nil, err
	}

	customer := Customer{Id: id}

	err = res.Scan(&customer.FirstName, &customer.LastName)

	if err != nil {
		return nil, err 
	}

	return customer, nil
}