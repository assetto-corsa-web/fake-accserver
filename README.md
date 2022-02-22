# Fake ACC Dedicated Server

This project was built to support the accweb development. It's a really dumb log file reader and dumps the lines to stdout every 50 milliseconds.

That's it!

## Getting Started

### Checkout the project

`$ git clone git@github.com:assetto-corsa-web/fake-accserver.git`

### Add the log files

Add accweb logs file inside the `logs` directory and it will be used automatically by the fake accServer.

Keep in mind that every time you add a new file in `logs` dir, you need to rebuild the fake accServer

### Build for your desired architecture

There is a `make` command to support you in the builds. 

OSX users:

`$ make build-for-osx`

Windows users:

`$ make build-for-windows`

### Configure accweb

Open the `config.yml` inside the accweb directory and change the attribute `acc.server_path` to the compiled fake accServer, e.g.: `/Users/pedro/workspace/assetto-corsa-web/fake-accserver/build/osx`
