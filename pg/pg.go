package pg

// Usage:
//connectionString := postgresql.BuildConnectionString(
//	env.PostgresHostname,
//	env.PostgresPort,
//	env.PostgresDatabase,
//	env.PostgresSslMode,
//	env.PostgresUsername,
//	env.PostgresPassword,
//)
func BuildConnectionString(host string, port string, databaseName string, sslMode string, user string, password string) string {
	if sslMode == "" {
		sslMode = "disable"
	}
	connectionString := "host=" + host + " port=" + port + " dbname=" + databaseName + " sslmode=" + sslMode
	if len(user) > 0 {
		connectionString += " user=" + user + " password=" + password
	}
	return connectionString
}
