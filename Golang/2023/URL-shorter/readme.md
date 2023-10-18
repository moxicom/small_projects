	```
    rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name, email string
		if err := rows.Scan(&name, &email); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %s, Email: %s\n", name, email)
	}
    ```