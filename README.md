# WST_lab1_client
#To create a request, use the following command line arguments: 
-url string SOAP server URL (default "http://localhost:8094/soap") 
-method string Method to call (addperson|getperson|getallpersons|updateperson|deleteperson|searchperson) 
-id int ID of the person (required for getperson, updateperson and deleteperson) 
-name string Name of the person (required for addperson and updateperson) 
-age int Age of the person (required for addperson and updateperson) 
-surname string Surname of the person (required for addperson and updateperson)

-query string Query for searching person (required for searchperson)

example request: go run main.go -method searchperson -query Иван -url http://127.0.0.1:8094/soap example response:

ID: 642, Name: Владимир, Surname: Иванов, Age: 26
ID: 643, Name: Иван, Surname: Иванов, Age: 27
