# Objective
- Build Blog with Go & React

# Implementation Details
- DB Connection using gorm
- Dealing with environment variables using godotenv
- Hash functions for password using dcrypt
- Routes for API requests using Fiber
- Middleware for handling Client's request 
- Post (Create, Update, Delete)
- Upload image to local storage

# What's not in this project
- Delete account
- Logout account with session

# CMD
- Live Run => <code>nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go </code>

# Development Tool (Optional)
- [Golang](https://go.dev/)
- [XAMPP](https://www.apachefriends.org/) for development in local environment
- [Postman](https://www.postman.com/), or Thunder Client for Testing API

# GO Packages
- [GORM](https://gorm.io/docs/index.html) for SQL connction
- [Fiber](https://github.com/gofiber/fiber) for Web Framework
- [godotenv](https://github.com/joho/godotenv) for using environment variables
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) for using hash function => <code>go get golang.org/x/crypto/bcrypt</code>
- [JWT](https://github.com/dgrijalva/jwt-go/) for client side authentication

# Resources
- [Tutorial](https://www.youtube.com/playlist?list=PLJ2eCBnvv6JVQtnuKbtJSRu0OkuNicOeW)
- [Official GitHub](https://github.com/kingztech2019/go-blogbackend)

# Personal-Use-Only
- [This URL](https://github.com/jinheehanaaa/TUT-Blog-with-Go-and-ReactJS-TailwindCSS)
- [PhpMyAdmin](http://localhost/phpmyadmin/)
- [Note](NOTE/note-taking/)