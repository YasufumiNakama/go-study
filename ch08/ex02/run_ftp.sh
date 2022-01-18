# brew install inetutils
ftp localhost 8000

<<COMMENT_OUT
(base) nakamayasufumi@NAKAMAnoMacBook-Pro ex02 % ./run_ftp.sh
ftp: connect to address ::1: Connection refused
ftp: Trying 127.0.0.1 ...
Connected to localhost.
220 Service ready for new user.
Name (localhost:nakamayasufumi): user
230 User user logged in, proceed.
ftp> ls
200 Command okay.
150 File status okay; about to open data connection.
.DS_Store
README.md
ftp
go.mod
image
main.go
run_ftp.sh
run_server.sh

226 Closing data connection. Requested file action successful.
ftp> get main.go 
200 Command okay.
150 File status okay; about to open data connection.
WARNING! 50 bare linefeeds received in ASCII mode
File may not have transferred correctly.
226 Closing data connection. Requested file action successful.
678 bytes received in 0.000553 seconds (1.17 Mbytes/s)
ftp> cd image 
200 Command okay.
ftp> ls
200 Command okay.
150 File status okay; about to open data connection.
sample.png

226 Closing data connection. Requested file action successful.
ftp> binary
200 Command okay.
ftp> get sample.png
200 Command okay.
150 File status okay; about to open data connection.
226 Closing data connection. Requested file action successful.
5967 bytes received in 0.000563 seconds (10.1 Mbytes/s)
ftp> bye
221 Service closing control connection.
COMMENT_OUT
