# Redditex

a simple fullstack app inspired on redit
![](./docs/redditex/png)

## What is this?

This can be called a 'personal' project by [@Haizza1](https://github.com/Haizza1) with the original intentions of learning how to create a fullstack application, learn how to setup and deploy the project to production using a cloud service provider. for this project i use digital ocean to host the backend and vercel to host the frontend, and create a service that use graphql to communicate frontend with the backend. if you want know more about this project chekout the wiki

## Technologies used for this project
**Backend**
* [Golang](https://golang.org/) 
* [Echo](https://echo.labstack.com/)
* [Gorm](https://gorm.io/)
* [Gqlgen](https://gqlgen.com/)
* [Redis](https://redis.io/)
* [Postgresql](https://www.postgresql.org/)

**Frontend**
* [TypeScript](https://www.typescriptlang.org/)
* [Next.js](https://nextjs.org/)
* [React](https://reactjs.org/)
* [Urql](https://formidable.com/open-source/urql/)
* [Chakra-ui](https://chakra-ui.com/)

## Feedback

Should you like to provide any feedback, please open up an Issue, I appreciate feedback and comments.

## Usage
To run locally you will need to have Postgresql, Redis and yarn install on you machine if you don't know how to install all this check this links:
* [Install Postgres](https://www.postgresqltutorial.com/install-postgresql/)
* [Install Redis](https://redisson.org/articles/how-to-install-redis.html)
* [Install Yarn](https://classic.yarnpkg.com/en/docs/install/#windows-stable)

Now that you have installed all this is time to set up the backend

### **Backend Set Up**
First check the .env.example file. The only thing that you have to change is the password that you set for your db:

```
DATABASE_URL=postgresql://postgres:<your password>@localhost:5432/redittex
```

if you have redis setup with password will need to added in the REDIS_PWD env variable.

```
REDIS_PWD=<your redis password>
```

know its time to run the app:
```
$ go build -o main .

$ ./main 
```

you should see something like this: 
```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.2.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:4000
```

### **Frontend Set Up**

first you have to create a file in the web folder:

```
$ touch .env.local
```

the only thing you have to put in this file is: 
```
NEXT_PUBLIC_API_URL=http://localhost:4000/graphql
```

know the only thing to do is run this commands:

```
$ yarn

$ yarn dev
```