package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/infra/repositories"
	"github.com/wrapped-owls/goremy-di/remy"
)

func RegisterDbConn(injector remy.Injector) {
	remy.Register(injector, remy.Singleton(func(retriever remy.DependencyRetriever) *sql.DB {
		config := remy.Get[DatabaseConfig](retriever)
		dataSourceUrl := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			config.User, config.Password, config.Host, config.Port, config.Database,
		)
		db, err := sql.Open("mysql", dataSourceUrl)
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
