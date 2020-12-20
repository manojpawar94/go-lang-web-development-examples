
## Creating template Using html/template package

Quite often, we would like to create HTML forms to get the information from a client in a specified format, upload files or folders to the server, and generate generic HTML templates, rather than repeating the same static text. We will be able to implement all these functionalities efficiently in Go.

Templates allow us to define placeholders for dynamic content that can be replaced with the values at runtime by a template engine. They can then be transformed into an HTML file and sent to the client. Creating templates in Go is fairly easy using
Go's html/template package.

### Serving static files over HTTP

While designing web applications, it’s always a best practice to serve static resources, such as .js, .css, and images from the filesystem, or any content delivery network (CDN), such as Akamai or Amazon CloudFront, rather than serving it from the web server. This is because all these types of files are static and do not need to be processed; so why should we put extra load on the server? Moreover, it helps to boost application performance, as all the requests for the static files will be served from external sources and therefore reduce the load on the server.

Go's net/http package is sufficient enough for serving static resources from the filesystem through FileServer.

## Screenshots
1. Index Page
![Image of Index](https://github.com/manojpawar94/go-web-app-examples/blob/main/gorilla-mux-http-server-html-template/screenshot/index_page.png?raw=true)

2. Login Page
![Image of Login](https://github.com/manojpawar94/go-web-app-examples/blob/main/gorilla-mux-http-server-html-template/screenshot/login_page.png?raw=true)