### Simple Mono Repo With Golang

#Detail

- Thir include 2 Service (Auth have 3 API (Register, Login, Me)  and Fetch have 2 API (Get Resource and Get Aggregat just for role admin))
- For get reource using external API for get resource data and using converter currency for convert currency from <b>IDR</b> to <b>USD</b> but for call converter currency just 1 hour once, because using <b>*Redis*</b> for memory cache

#How to use

**without docker**
- Open terminal and run :
```console
foo@bar:~$ cd Efishery-Test
foo@bar/Efishery-Test:~$ cd /Auth
foo@bar/Efishery-Test:/Auth~$ go run main.go
```
- And open terminal again for running app fetch
```console
foo@bar:~$ cd Efishery-Test
foo@bar/Efishery-Test:~$ cd /Fetch
foo@bar/Efishery-Test:/Fetch~$ go run main.go
```
- Open Postman and Import postman collection from directory doc
- For detail can see a video in directory doc

**with docker**
- Open terminal and run :
```console
foo@bar:~$ cd Efishery-Test
foo@bar/Efishery-Test:~$ docker-compose build
foo@bar/Efishery-Test:~$ docker-compose up
```
- Open Postman and Import postman collection from directory doc
- For detail can see a video in directory doc

<b>*Notes*</b> :
- in thir Role admin is 1 and user 2
