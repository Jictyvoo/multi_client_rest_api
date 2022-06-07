package infra

import (
	"database/sql"
	"fmt"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/xyc_2_core/internal/infra/repositories"
	_ "github.com/lib/pq"
	"github.com/wrapped-owls/goremy-di/remy"
)

func RegisterDbConn(injector remy.Injector) {
	remy.Register(injector, remy.Singleton(func(retriever remy.DependencyRetriever) *sql.DB {
		config := remy.Get[DatabaseConfig](retriever)
		connStr := fmt.Sprintf(
			"user=%s password='%s' host=%s port=%d dbname=%s sslmode=disable",
			config.User, config.Password, config.Host, config.Port, config.Database,
		)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		return db
	}))
}

func RegisterRepositories(injector remy.Injector) {
	// Factory for contacts repository
	remy.Register(injector, remy.Factory(func(retriever remy.DependencyRetriever) interfaces.ContactsRepository {
		db := remy.Get[*sql.DB](retriever)
		return repositories.NewContactsDbRepository(db)
	}))
}
