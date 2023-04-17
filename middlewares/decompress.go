package middlewares

import (
	"kook-bot-chatgpt/utils"
	"compress/zlib"
	"io/ioutil"
	"log"
	"net/http"
)

// 解压压缩请求中间件
func DecompressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("compress") != "0" {
			log.Println("Decompressing request body...")
			// 创建zlib解压器
			zlibReader, err := zlib.NewReader(r.Body)
			if err != nil {
				utils.ErrorLogger(w, "Error creating zlib reader")
				return
			}
			defer zlibReader.Close()

			// 替换请求体
			r.Body = ioutil.NopCloser(zlibReader)
			r.Header.Set("Content-Encoding", "deflate")
			r.ContentLength = -1
		}

		next.ServeHTTP(w, r)
	})
}