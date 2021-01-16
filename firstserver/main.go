package main

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	// ソースコードのコメントは以下のようになっている
	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()

	/*
		router.GET関数のソースコードに書いてあるコメント
		func (*gin.RouterGroup).GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
		(gin.RouterGroup).GET on pkg.go.dev
		GET is a shortcut for router.Handle("GET", path, handle)

		サーバーからの相対パスを第一引数
		リクエストが来た際に動作させる関数(ハンドラー)を第二引数に指定して実装することによりアプリケーションを作成していきます

		router.GET("/", func(ctx *gin.Context) {
			// ここに処理を書く
		})

		ではrouter.GETの第二引数に記述してある func(ctx *gin.Context)について読み解いていく

		下はソースコードのコメント
		Context is the most important part of gin. It allows us to pass variables between middleware,
		manage the flow, validate the JSON of a request and render a JSON response for example.

		いろいろなことができそう。
	*/
	router.GET("/", func(ctx *gin.Context) {
		/*
			ctx.JSON関数についての説明


			[前提知識]
			JSONとは「JavaScript Object Notation」の略で、「JavaScriptのオブジェクトの書き方を元にしたデータ定義方法」のこと。

			[関数についてのコメント]
			JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".

			構造体をJSONとしてレスポンスボディに記入する
			難しく考えず、JSON形式で表示すると読み替えていい。

			以下の例では"message": "hello, world"からなるgolangのmapをJSONに置き換えてレスポンスボディにシリアライズしている。
		*/
		ctx.JSON(200, gin.H{
			"message": "hello, world",
		})
	})

	// OSの環境変数にPORTが存在しているならばそのportを使用する
	port := os.Getenv("PORT")
	if port == "" {
		// 値が入っていなかったら:8080を表示
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	/*
		ListenAndServeについての説明

			ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Handler is typically nil, in which case the DefaultServeMux is used.

			ListenAndServe は TCP ネットワークアドレスのアドレ スをリッスンしてから Serve をハンドラで呼び出し、着信接続のリクエストを処理します。ハンドラは通常 nil で、その場合は DefaultServeMux が使用されます。

			これでポートで通信がくるのを待ち構えている。
	*/
	endless.ListenAndServe(":"+port, router)

	/*
		コードが読めたら実装に動作させてみよう
		_____
		cd firstserver
		go build
		./firstserver
		_____

		ログに以下のような感じに出力される
			[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

			[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
			 - using env:	export GIN_MODE=release
			 - using code:	gin.SetMode(gin.ReleaseMode)

			[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
			2021/01/16 17:37:54 Defaulting to port 8080
			2021/01/16 17:37:54 Listening on port 8080
			2021/01/16 17:37:54 68030 :8080

		webブラウザ(Safari Chrome..)でhttp://localhost:8080/と打ち込んで見て	"message": "hello, world"が表示されたら正しく動作しています。
	*/
}

/*
練習問題
1. "message": "hello, world"以外の文字を表示させてみよう
2. http://localhost:8080/goodbye/とブラウザで打ち込んだら "message": "goodbye"と表示させるようにrouter.GET関数を追加して
　 プログラムを改造してみよう。
*/
