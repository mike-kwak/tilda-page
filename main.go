package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var document = `
<!DOCTYPE html>
<html lang="ko">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>tilda.hellomike.page</title>
<style>
body {
    margin: 0;
    background-color: #F6F0E1;
    height: 100vh;
}
main {
    margin: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
main > div {
    font-family: 'Roboto Mono', 'Bitstream Vera Sans Mono', Menlo, Consolas, monospace;
    font-size: 12em;
    color: #030303;
}
</style>
</head>
<body>
<main><div>~</div></main>
</body>
</html>
`

var rootCmd = &cobra.Command{
	Use:   "tilda-page [address (default ':8181')]",
	Short: "tilda-page serve web page.",
	Run: func(cmd *cobra.Command, args []string) {
		addr := ":8181"
		if len(args) > 0 {
			addr = args[0]
		}
		http.HandleFunc("/", index)
		log.Fatal(http.ListenAndServe(addr, nil))
	},
}

func main() {
	err := Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Execute() error {
	return rootCmd.Execute()
}

func index(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()

	if r.Method == http.MethodGet && (r.URL.Path == "/" || r.URL.Path == "") {
		if strings.HasPrefix(r.Header["Accept"][0], "text/html") {
			w.Header().Add("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(document))
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("\"~\""))
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
