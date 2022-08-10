# tech-acc-chatroom
A community web app for communicating with members of TA team. 

## Application Design
The app follows:
http://<servername>/<handlername>?<parameters>

## Logic
When the request reaches the server, a multiplexer will inspect the URL being requested and redirect
the request to the correct handler. ONce the request reaches the handler, the handler will retrieve
information from the request and process it accordingly. After this, the handler passes the data to the
template engine, which uses templates to generate html for the client to see.

## Data model
Data is persisted in a PostgresSQL database. The data model consists of four data structures:
- a user: representing the forum user's information
- a session: representing the current login session of a user
- a thread: representing a conversation among the other users
- a post: representing a message sent by a user in or outside a thread.

All members of TA team can log in and create posts in threads. Non-members can read threads but not 
create posts.
The web app accesses the database in these four ways.

## Cookies for access control
When a user logs in, he has to has access control to specific pages as a user. This is made possible by cookies.
