package infra

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/core"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/domain/interfaces"
	"github.com/jictyvoo/multi_client_rest_api/modules/abz_1_core/internal/infra/repositories"
	"github.com/wrapped-owls/goremy-di/remy"
)

func init() {
	remy.Register(core.Injector, remy.Singleton(func(retriever remy.DependencyRetriever) *sql.DB {
		db, err := sql.Open("mysql", "user:password@/dbname")
		if err != nil {
			panic(err)
		}
		return db
	}))
}

func init() {
	// Factory for contacts repository
	remy.Register(core.Injector, remy.Factory(func(retriever remy.DependencyRetriever) interfaces.ContactsRepository {
		db := remy.Get[*sql.DB](retriever)
		return repositories.NewContactsDbRepository(db)
	}))
}
