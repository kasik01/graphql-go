package main

import (
	"graphql-hasura-demo/graph"
	"graphql-hasura-demo/internal/config"
	"graphql-hasura-demo/internal/database"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

// const defaultPort = "8080"

func main() {
	config.LoadEnv()
	defaultPort := config.GetAppConfig().APP_PORT
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	database.InitDb()
	// Load env variables

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	// Route cho download file
	// http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
	// 	filename := filepath.Base(r.URL.Path)
	// 	filePath := filepath.Join("/tmp", filename)

	// 	// Kiểm tra file tồn tại
	// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
	// 		http.Error(w, "File không tồn tại", http.StatusNotFound)
	// 		return
	// 	}

	// 	// Thiết lập header để download
	// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	// 	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// 	// Phục vụ file
	// 	http.ServeFile(w, r, filePath)

	// 	// Xóa file sau khi phục vụ (tùy chọn)
	// 	if err := os.Remove(filePath); err != nil {
	// 		log.Printf("Xóa file %s thất bại: %v", filePath, err)
	// 	}
	// })

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
