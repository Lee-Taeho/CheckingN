# 2021-fall-cs160-chocolate-cake

<h1 align="center"> CheckingN </h1><br>

![CheckingN](https://user-images.githubusercontent.com/60446185/144513345-67246f9a-1ad9-45b5-b729-93af535b28be.png)

<p align="center"> Book tutoring appointments with CheckingN. </p>

## Table of Contents

- [Key Features](#key-features)
- [Dependencies](#dependencies)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Running](#running)
- [Contributors](#contributors)

## Key Features

1. Google Sign-in
2. Book an appointment with available tutor
3. Book online appointments through Zoom

## Dependencies

- Golang
- React-datepicker
- Font Awesome

## Getting Started

### Installation
In IDE: 
1. Install MongoDB and Golang
2. Clone the Github repository with the following link:
https://github.com/Lee-Taeho/2021-fall-cs160-chocolate-cake.git
3. Go to the client directory in the terminal with:
```bash
cd client
```
4. Install packages and dependencies with the following:
```bash
npm install
```
5. Installing font awsome
In the client directory,
```bash
npm i --save @fortawesome/fontawesome-svg-core
npm install --save @fortawesome/free-solid-svg-icons
npm install --save @fortawesome/react-fontawesome
npm install react-datepicker --save
```

6. Installing react-datepicker
```bash
npm install react-datepicker --save
```
Building the backend server
8. Go to the server directory in the terminal with:
```bash
cd server
```
9. Build packages for Golang using:
```bash
go build
```
10. Then run the following to connect to the database:
```bash
./run.sh
```

In Docker:
1. Install/build Docker by following:
https://docs.docker.com/get-started/
2. Go the the client directory and run the following:
```bash
cd client 
docker build -t client 
```
3. Create a new terminal, go to the server directory, and run the following:
```bash
cd server 
docker build -t server 
```


## Running

In order to build the application in your IDE, run the following:

In the client directory

```bash
npm start
```

In the server directory

```bash
./run.sh
```
To run using Docker, run the following:

In terminal, run
```bash
docker run -p 3000:3000 client
docker run -p 8080:8080 server
```
Open http://localhost:3000 to view it in the browser.

## How to use branching? 
With Github Desktop
1. Choose the branch you want to use
2. Click pull origin
3. Click the current branch and click new branch
4. Type the name of the branch and click create branch



## Contributors

|         Name         | GitHub                                                       |
| :------------------: | ------------------------------------------------------------ |
|      Nhien Lam       | [@NhienLam](https://github.com/NhienLam)                     |
|      Tae Ho Lee      | [@Lee-Taeho](https://github.com/Lee-Taeho)                   |
| Ekaterina Kazantseva | [@kate-kazantseva](https://github.com/kate-kazantseva)       |
|  Christine Lantaca   | [@clantaca](https://github.com/clantaca)                     |
|     Khang Nguyen     | [@kharanga](https://github.com/kharanga)                     |
|   Ayush Maheshwari   | [@ayushmaheshwari768](https://github.com/ayushmaheshwari768) |
