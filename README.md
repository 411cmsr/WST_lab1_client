# WST_lab1_client
#### To create a request, use the following command line arguments:

-url string 

######     _SOAP server URL (default "http://localhost:8094/soap")_ 

-method string 

######     _Method to call (addperson|getperson|getallpersons|updateperson|deleteperson|searchperson)_ 

-id int 

######     _ID of the person (required for getperson, updateperson and deleteperson)_ 

-name string 

######     _Name of the person (required for addperson and updateperson)_ 

-age int 

######     _Age of the person (required for addperson and updateperson)_ 

-surname string 

######     _Surname of the person (required for addperson and updateperson)_

-query string 

###### _Query for searching person (required for searchperson)_


### **example request:** 

#### go run main.go -method searchperson -query Иван -url http://127.0.0.1:8094/soap 

### **example response:**

_ID: 642, Name: Владимир, Surname: Иванов, Age: 26_
_ID: 643, Name: Иван, Surname: Иванов, Age: 27_
