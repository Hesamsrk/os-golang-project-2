# Problem 3 (Server client exchange time using websockets)

This example shows a simple client and server implemented in golang and pure JS (as front-end client).

The server sends its current time.
The client sends a message every second
and prints all messages received.

To run the example, start the server:````

    $ go run server.go

Next, start the client:

    $ go run client.go

The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.
