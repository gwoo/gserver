GServer
=======

A simple static file server in Go. The first argument is the directory to serve. If an `index.html` exists it will be output, otherwise a directory listing will be visible.

	Usage of ./gserver:
	  -port=8888: Server port.

Example
-------
	gserver -p 8888 /var/www

Adding SSL
----------
Files are served through SSL, if a `cert.pem` and `key.pem` exist in the current working directory where `gserver` is started.

To generate, use openssl

	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem