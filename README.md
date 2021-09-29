Results of working through the examples at: 

https://www.howtographql.com/graphql-go/0-introduction/

To regenerate the tables.

docker run -v /Users/timcutsforth/src/learn/go/gosamples/graphql_demo/internal/pkg/db/migrations/mysql/:/migrations --network host migrate/migrate -path=/migrations/ -database mysql://root:dbpass@/hackernews up
