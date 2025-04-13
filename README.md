# akicli-go
 
```go
func main() {
	a, err := api.NewAPI(
		client.WithBaseURL("https://baseurl"),
		client.WithToken("Token tkn"),
	)
	if err != nil {
		log.Fatal(err)
	}
	tags, err := a.GetProjectApp(9837, 16815)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tags)
}
```