package database

import (
	"log"

	"github.com/gocql/gocql"
)

type DatabaseCollection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var connection DatabaseCollection

func InitDatabase() {
	connection.cluster = gocql.NewCluster("127.0.0.1")
	connection.cluster.Consistency = gocql.Quorum // Consistency
	connection.cluster.Keyspace = "first_app"
	connection.session, _ = connection.cluster.CreateSession()
}

func ExecuteQuery(query string, values ...interface{}) error {
	if err := connection.session.Query(query).Bind(values...).Exec(); err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	return nil
}

func GetQuery(query string, values ...interface{}) (map[string]interface{}, error) {
	dest := map[string]interface{}{}
	if err := connection.session.Query(query).Bind(values...).MapScan(dest); err != nil {
		log.Printf("Error scan query: %v", err)
		return nil, err
	}
	return dest, nil

}
