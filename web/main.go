package main

import (
	"log"
	// "github.com/julienschmidt/go-http-routing-benchmark"
	// "github.com/julienschmidt/httprouter"
	// "github.com/tnkyk/clean_book_go/config"
	// logging "github.com/tnkyk/clean_book_go/handler"
	// "github.com/tnkyk/clean_book_go/handler/rest"
	// "github.com/tnkyk/clean_book_go/infra/persistence"
	// "github.com/tnkyk/clean_book_go/usecase"
)

func main() {
	//指定したLogFileにlogを出力するための設定
	// logging.LoggingSetting(config.Config.LogFile)

	// userPersistence := persistence.NewUserPersistence()
	// userUseCase := usecase.NewUserUsecase(userPersistence)
	// userHandler := rest.NewUserHandler(userUseCase)

	// httprouterを用いてハンドラーを登録する
	// Router := httprouter.New()
	// Router.GET("/api/users", userHandler.Index)
	// Router.POST("/api/signup", userHandler.SignUp)
	// Router.POST("/api/signin", userHandler.SignIn)
	//ポート9000で待ち受け
	// err := http.ListenAndServe(":9000", Router)
	// if err != nil {
	// log.Fatalf("Listen and serve failed. %+v", err)
	// }
	log.Println("Server running...")
}
