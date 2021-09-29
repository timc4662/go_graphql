To regenerate the tables.

timcutsforth@Tims-Air graphql_demo % docker run -v /Users/timcutsforth/src/learn/go/gosamples/graphql_demo/internal/pkg/db/migrations/mysql/:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://root:dbpass@/hackernews up
