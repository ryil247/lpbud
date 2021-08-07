package main

import (
	"fmt"
	"time"

	"github.com/ryil247/lpbud/backend/internal/handler"
	"github.com/ryil247/lpbud/backend/internal/subscription/static"
)

func main() {
	h := handler.Handler{} // Handler handler = new Handler()
	sub := static.New()    // Subscription sub = new Static()
	sub.RegisterSubjectAddedHandler(h.AddSubjectHandler)

	<-time.After(10 * time.Second)

	fmt.Printf("handler received %d subject.add-requests\n", len(h.Subjects))
}
