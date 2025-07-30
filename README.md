
# Airline Voucher Seat Assignment App

A voucher seat assignment web application for an airline campaign. For each
flight, the airline wants to randomly assign 3 unique seat numbers, with specific seat maps
depending on aircraft type.
## Candidate

- [Nicolas Raja Oloan Manurung](https://github.com/Nicomnr02)


## Tech Stack

**Client:** React

**Server:** Go

**Database**: SQLite


## Run Locally 

Clone the project

```bash
  git clone https://github.com/Nicomnr02/bookcabin-flight-voucher-assignment.git
```

Go to the **BE** project directory

```bash
  cd backend
```

Install dependencies

```bash
  go mod tidy
```

Init SQLite GUI (Optional)

```bash
  make run-sqlite
```

Start the server

```bash
  make run-server
```


Go to the **FE** project directory

```bash
  cd ../frontend
```

Install dependencies

```bash
  npm install
```

Start the Web Application

```bash
  npm run dev
```

Go the Web Application

-  [http://127.0.0.0:5173](http://127.0.0.0:5173)

See the SQLite GUI

-  [http://127.0.0.0:8080](http://127.0.0.0:8080)

