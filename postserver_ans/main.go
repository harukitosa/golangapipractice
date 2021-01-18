package main

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

/*
練習問題
1. curlコマンドを使用して"today coding"以外の文字をPOSTリクエストで送ってみよう
2. 今回TaskのフィールドはContentのみだった、Taskに新たなフィールドNameを追加して以下のコマンドできちんとデータを送信できているか確認しよう
   curl -X POST -H "Content-Type: application/json" -d '{"content":"today coding", "name":"tosa"}' localhost:8080/post
3. [応用]グローバル変数を使用して、サーバーが動作している間はデータが保存されるようにしてみよう。golangグローバル変数で検索->POSTリクエストで
   送られてきたTaskのデータをグローバル変数の配列に格納->GETリクエストを作成して現在グローバル変数に保存されている値をwebブラウザの画面に表示
   させてみよう。

   3の回答例
*/

// Task is ...
type Task struct {
	Content string `json:"content"`
}

// グローバル変数にTask構造体のスライス（配列）を作成する
// この変数にデータを貯めていく
// もちろんプログラムが終了(サーバーが停止)したらデータは消える
// が作動中はデータを保持することができる。
var list []Task

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		// 現在のlistの値を返す
		ctx.JSON(200, list)
	})

	router.POST("/post", func(ctx *gin.Context) {
		var req Task
		ctx.BindJSON(&req)
		// 受け取ったreqの値をリストにappend関数を用いて追加する
		list = append(list, req)
		log.Println(list)
		ctx.JSON(200, "save ok")
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
		上のコマンドを複数回打ち込んだ後webブラウザでhttp://localhost:8080/と打ち込んでみよう
		いままでにPOSTしたデータをみることができる。
	*/
}
