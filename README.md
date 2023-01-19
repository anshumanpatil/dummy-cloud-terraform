# Dummy Cloud - Web Console, Rest Client, Terraform

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

Create,Read,Update & Delete resorces and view on web console.

- Postman Support for Rest API's.
- Real-Time Web Console for resources.
- Custom Terraform Plugin for resource operations.


## Installation

Dummy Cloud requires [Golang](https://go.dev/) to compile terraform plugin locally.
Dummy Cloud requires [Docker](https://www.docker.com/) to run on docker(Web console & Rest API) .

Install the dependencies and devDependencies and start the server.

```sh
cd dillinger
npm i
node app
```

For production environments...

```sh
npm install --production
NODE_ENV=production node app
```


## Tech

Dillinger uses a number of open source projects to work properly:

- [AngularJS] - HTML enhanced for web apps!
- [Ace Editor] - awesome web-based text editor
- [markdown-it] - Markdown parser done right. Fast and easy to extend.
- [Twitter Bootstrap] - great UI boilerplate for modern web apps
- [node.js] - evented I/O for the backend
- [Express] - fast node.js network app framework [@tjholowaychuk]
- [Gulp] - the streaming build system
- [Breakdance](https://breakdance.github.io/breakdance/) - HTML
to Markdown converter
- [jQuery] - duh

And of course Dillinger itself is open source with a [public repository][dill]
 on GitHub.

## License

MIT
