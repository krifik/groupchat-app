# groupchat-app
a groupchat demo-app build with golang and vuejs


## Prerequisite
- Golang v1.19
- Nodejs v18.10.0
- PostgresSQL v14
- Vitejs v2.6.2

## Installation

Clone this repo

```bash
git clone https://github.com/krifik/groupchat-app.git
```
Change dir to /groupchat-app

```bash
cd /groupchat-app
```
Change dir again to /frontend
and install depedency using yarn or npm

```bash
cd /frontend
npm install
```
last step, go to /backend dir
and install depedency, make sure you have golang installed on your system

```bash
cd ../backend
go mod tidy
```

## Usage
When everything is installed, you can use this app by 
go to /backend dir and run :
```bash
go run main.go
```
then, go to /frontend dir and run :
```bash
npm run build
npm run preview
```
or, just 
```bash
npm run dev
```




