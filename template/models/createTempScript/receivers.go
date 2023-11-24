package main

var receiver = `package {dataModel}

func ({lowerModelName} *{upperModelName}) TableName() string {
	return "{tableName}"
}
`
