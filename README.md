## goweb
A simple web application demostrating depency injection in golang.

### Note
Dependency Injection is a design pattern used to achieve loose coupling between components and Loosely coupled code refers to a design where the components or modules in a system have minimal dependencies on each other.
In a loosely coupled system, the components interact through well-defined interfaces or abstractions rather than directly relying on concrete implementations. This promotes flexibility, reusability, and maintainability, as individual components can be modified or replaced without affecting the entire system.


### Required
Go
sqlite3

### Installation
```
git clone https://github.com/ftsog/goweb
```

#### Run
change into the program directory
```
cd goweb
```

run the program using the below command once in the project directory
```
go run .
```
### Routes
1. 127.0.0.1:8081/create (POST)
2. 127.0.0.1:8081/read?postID=3 (GET)
3. 127.0.0.1:8081/delete?postID=3 (DELETE)

#### Message
When you make a request to delete the blog using the url at number 3 above, the blog post will got deleted and you won't be able to access it anymore meaning you will need to create another blog post using the url at number 1 above, once you create a new blog post the new blog post ID auto increament in the database table meaning to perfom subsequent read and delete request, you will need to change the postID valude to 4 or increament the last deleted blogpost ID by 1.

### Note
If you understand this program source code then you understand dependency injection and how it is done in golang
