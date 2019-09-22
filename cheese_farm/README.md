# The Cheese Service

This is a small service using gRPC. The message definition is on `cheese/cheese.proto` and we implement
a client and a server. To check this service locally, please follow the instructions below:


### Install dependencies

First, install library dependencies on your machine. Make sure to find the proper package for your
Operating System. Some of them are presented here:

```bash
$> zypper install protobuf-devel  # For Opensuse linux distribution
$> apt install protobuf-compiler  # For Ubuntu linux distribution
$> brew install protobuf          # For Mac OS X operating system
```


### Compilation

First we need to compile Golang binaries. We use a `Makefile` to manage local tasks.
To see all available target rules on the Makefile run:

```bash
$> make
```

We have two binaries to be compiled. The `server` and `client` binaries.
In order to compile both of them you should run:

```bash
$> make build
```

Both binary files will be placed at project `bin` directory.


### Run the server

The server will listen on localhost at port 4000:

```bash
$> make run_server
```

### Run clients

Open as many terminal you want and run an instance of the client code:

```bash
$> make run_client
```


### Other language client

There is a client written in python. You can try it with:

```bash
$> make run_pyclient
```

You'll be able to see different clients talking to the server using the
same contract: the protocol buffer definition. Try reading all the server and
clients code to get the gRPC way of building reliable and efficient services.

