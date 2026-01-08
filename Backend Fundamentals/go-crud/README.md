The applications server expects that you have postgres running on `PORT:5432`. Enter postgres database connection URL in `.env` file as following

```
DB_URL= //your_database_url
```

After making database conection use make command mentioned in `Makefile` to migrate up the database to include the required table `users`.

Postman collection link:
[collection](https://warped-meadow-913182.postman.co/workspace/New-Team-Workspace~850b93a7-4078-4f7e-bcb5-331e137d6e73/folder/32759292-ffab62fa-9825-4940-8972-730a5a448c16?action=share&creator=32759292&ctx=documentation)

Since this was just a basic assignment I have kept all application logic in `main.go` if this was a larger project then we can also use the expected MVC architecture.

It'd be great if you'd accept this assignment in this much shorter time, I've tried best of my knowledge!
