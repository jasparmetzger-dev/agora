package conf

import (
	"os"
)

//get enviroment variables
//WHY DOESNT CONST WORK??

// database
var DATABASE_URL string = os.Getenv("DATABASE_URL")

// auth
var SECRET_KEY string = os.Getenv("SECRET_STRING")

// ports
var BACKEND_PORT string = os.Getenv("BACKEND_PORT")
