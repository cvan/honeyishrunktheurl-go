# honeyishrunktheurl-go

World's tiniest URL shortener. Written in Go.

Want to add a site? [Update `sites.json`.](https://github.com/cvan/honeyishrunktheurl-go/edit/master/sites.json)


## Dependencies

* Git
* Go 1.1+


## Installing

Copy over a new `sites.json`:

    cp sites.json{.dist,}

To build from source:

    go get -v github.com/cvan/honeyishrunktheurl-go
    go build github.com/cvan/honeyishrunktheurl-go
    $GOPATH/bin/server.o

Or from cloning the repo:

    git clone https://github.com/cvan/honeyishrunktheurl-go.git
    cd honeyishrunktheurl-go
    ln -s . $GOPATH/src/github.com/cvan/honeyishrunktheurl-go
    go get .
    go build -o server.o
    ./server.o


## Development

I use [autoenv](https://github.com/kennethreitz/autoenv) for managing environment variables. To use the sample `.env` file:

    cp .env{.dist,}


## Usage

Want to add a site? [Update `sites.json`.](https://github.com/cvan/honeyishrunktheurl-go/edit/master/sites.json)

To build the server:

    make  # go fmt, get, build, etc.
    ./server.o  # Run!

To start the server (running on port 5000 by default):

    ./server.o

Then fire up your browser to one of the short URL patterns specified in the `sites.json` file (e.g., [http://localhost:5000/cvan](http://localhost:5000/cvan)).

To run the server on a different port:

    PORT=8000 ./server.o


## Credits

* @potch for the idea
