package maindb

type TablesSt struct {
	GooseDbVersion string `json:"goose_db_version" db:"goose_db_version"`
	User           string `json:"user" db:"user"`
}

var Tables = TablesSt{
	GooseDbVersion: "goose_db_version",
	User:           "user",
}

// Named "T" for shortness
var T = Tables
