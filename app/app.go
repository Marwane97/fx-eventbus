package app

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/service"
	"net/http"
)

func handle(search *service.Search) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		isException := false
		for i := 0; i < 3; i++ {
			if i > 3 {
				isException = true
			}

			err := search.Search(isException)
			if err != nil {
				fmt.Println("An Error Occurred ...")
				return
			}
		}
	})

	fmt.Println("Server is listening on :3003")
	if err := http.ListenAndServe(":3003", nil); err != nil {
		panic(err)
	}
}
