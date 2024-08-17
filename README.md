## Corty.co 🩳

Corty.co is a simple and efficient URL shortener app written in Go. It allows users to shorten long URLs into concise, easy-to-share links.

## Features

- **Shorten URLs**: Convert long URLs into short, shareable links.
- **Redirects**: Redirects users to the original URL when the short link is accessed.

This is using the The [LeapKit](https://leapkit.dev/) template with [Postgres](https://www.postgresql.org/).

## Installation

#### **Clone the repository**

```bash
$ git clone https://github.com/wawandco/cortico.git
$ cd cortico
```

#### **Setup**
```sh
$ go mod download
$ go run ./cmd/setup
```

### Running the application
To run the application in development mode execute:

```sh
$ kit s
```

And open `http://localhost:3000` in your browser.

### ENV Variables
```sh
DATABASE_URL
BASE_URL
```
