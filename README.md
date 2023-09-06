# modak-test
This is a Tech Test for Modak.

### Quick rundown of the program
This is made with golang for a rate limit version a notification service.  
It has a in memory kind of database with all the scafolding and a baseline implementation of a notification client.

### How would I improve the code?
* I would implement an API using GIN for a better handling of events or even a queue system, ensuring the consistency of notifications and better handling of them.  
* Implement unit testing using gomock.  
* The notifications rules should be passed as a configuration in the deployment process, so we don't need to modify the code to add another rule.  
* We could add a metric for counting how many notifications are being rate limited.  


## How to run the program
```bash
go run main.go
```

To modify the number of users created and the rules for the notifications, you need to modify the main.go file.
