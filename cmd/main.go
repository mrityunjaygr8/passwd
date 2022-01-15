package main

import "github.com/mrtyunjaygr8/passwd/cmd/cmd"

// "context"
// "fmt"
// "log"

// "github.com/mrtyunjaygr8/passwd/ent"

// _ "github.com/lib/pq"

func main() {
	// fmt.Println("yp go")

	// fmt.Println("another thing")
	// fmt.Println("ek aur")

	// client, err := ent.Open("postgres", "host=localhost port=5432 user=mgr8 dbname=passwd password=dr0w.Ssap sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Close()

	// ctx := context.Background()
	// if err := client.Schema.Create(ctx); err != nil {
	// 	log.Fatal("failed creating schema resources: %w", err)
	// }

	// newUser, err := client.User.Create().SetEmail("msyt1969@gmail.com").SetPassword("dr0w.Ssap").Save(ctx)
	// if err != nil {
	// 	log.Println("error in creating user: %w", err)
	// }
	// fmt.Println(newUser)

	// rootCmd.Execute()

	cmd.Execute()

}
