package infra

import (
	"database/sql"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/utils"
	_ "github.com/lib/pq"
	"github.com/wrapped-owls/goremy-di/remy"
)

func init() {
	remy.Register(utils.Injector, remy.Singleton(func(retriever remy.DependencyRetriever) *sql.DB {
		connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		return db
	}))
}
