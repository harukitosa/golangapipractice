package main

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// Task is ...
// 通信で受け取る構造体を定義
type Task struct {
	// jsonに変換されるにはフィールドの変数の頭が大文字でなくてはならない
	// × content ○ Content
	Content string `json:"content"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})

	// router.POST関数のコメントを読む
	// POST is a shortcut for router.Handle("POST", path, handle).
	// = router.Handle("POST", path, handle)を省略してrouter.POST(path, handle)と書けるようだ
	//
	// GETのときとコードはほとんど同じだが今度はサーバー側(このプログラム)で値を受け取らなければならない
	router.POST("/post", func(ctx *gin.Context) {
		// Task構造体の変数を呼び出す
		// Task構造体のような形のリクエストを受け取るために宣言している。
		var req Task

		/*
				c.BindJSON関数ののソースコードのコメントを読む
				BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
				上のように記述されている。c.MustBindWith(obj, binding.JSON)関数というものとc.BindJSONは同一のものらしい
				ではc.MustBindWith(obj, binding.JSON)はどのように動作しているのだろう。

				以下がgin(Web Framework)のMustBindWith関数の記述とコメントだ

			    MustBindWith binds the passed struct pointer using the specified binding engine.
				It will abort the request with HTTP 400 if any error occurs.
				See the binding package.
				```
				func (c *Context) MustBindWith(obj interface{}, b binding.Binding) error {
					if err := c.ShouldBindWith(obj, b); err != nil {
						c.AbortWithError(http.StatusBadRequest, err).SetType(ErrorTypeBind) // nolint: errcheck
						return err
					}
					return nil
				}
				```

				英語は難しいが、第一引数で渡された構造体のポインタに第二引数のバインディングエンジンと呼ばれるものを使用してバインドしているらしい。
				ここではJSON形式でデータが渡されることが想定されているのでJSON形式のものをgolangの構造体(今回はTask)に変換する部分
				である。
				なので今回はc.BindJSON(&req)を呼ぶとc.MustBindWith(&req, binding.JSON)関数がよばれてリクエストボディにあるデータを変換してくれている。

		*/
		ctx.BindJSON(&req)
		// reqにどのような値が入っているかログを出力して確認してみよう
		log.Println(req)

		// 最後に渡されたデータをそのまま返してみる。
		// ctx.JSONは前回にもお馴染みだ
		// 前回は言い忘れたが第一引数の200はhttpステータスコードだ200, 404, 500はよく使用される
		ctx.JSON(200, req)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	endless.ListenAndServe(":"+port, router)

	/*
		コードが読めたら実装に動作させてみよう
		_____
		cd postserver
		go build
		./postserver
		_____

		ログに以下のような感じに出力される
		[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

		[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		 - using env:	export GIN_MODE=release
		 - using code:	gin.SetMode(gin.ReleaseMode)

		[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
		[GIN-debug] POST   /post                     --> main.main.func2 (3 handlers)
		2021/01/17 11:56:00 Defaulting to port 8080
		2021/01/17 11:56:00 Listening on port 8080
		2021/01/17 11:56:00 81699 :8080

		今回はwebブラウザは使用できない(urlを打ち込んで飛ぶのはGETリクエストだからだ)
		なのでcurlと呼ばれる便利なコマンドラインツールを使用する。
		サーバー(./postserverが動作しているターミナル)意外にもう一つターミナルを開いて

		以下のコマンドを打ち込んでみる。（コマンドの詳細は調査してみよう)
		curl -X POST -H "Content-Type: application/json" -d '{"content":"today coding"}' localhost:8080/post

		サーバーのログが表示されていて、きちんとcurlを打った方にもデータが帰ってきていればOKだ

		ginだとこんな感じ
		[GIN] 2021/01/17 - 11:58:38 | 200 |     175.075µs |             ::1 | POST     "/post"
		2021/01/17 11:58:57 {today coding}
	*/
}

/*
練習問題
1. curlコマンドを使用して"today coding"以外の文字をPOSTリクエストで送ってみよう
2. 今回TaskのフィールドはContentのみだった、Taskに新たなフィールドNameを追加して以下のコマンドできちんとデータを送信できているか確認しよう
   curl -X POST -H "Content-Type: application/json" -d '{"content":"today coding", "name":"tosa"}' localhost:8080/post
3. [応用]グローバル変数を使用して、サーバーが動作している間はデータが保存されるようにしてみよう。golangグローバル変数で検索->POSTリクエストで
   送られてきたTaskのデータをグローバル変数の配列に格納->GETリクエストを作成して現在グローバル変数に保存されている値をwebブラウザの画面に表示
   させてみよう。
*/
