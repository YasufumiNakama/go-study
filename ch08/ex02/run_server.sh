go run main.go -port 8000

<<COMMENT_OUT
(base) nakamayasufumi@NAKAMAnoMacBook-Pro ex02 % ./run_server.sh
2022/01/19 00:52:00 >> 220 Service ready for new user.
2022/01/19 00:52:03 << USER [user]
2022/01/19 00:52:03 >> 230 User user logged in, proceed.
2022/01/19 00:52:06 << PORT [127,0,0,1,201,165]
2022/01/19 00:52:06 >> 200 Command okay.
2022/01/19 00:52:06 << LIST []
2022/01/19 00:52:06 >> 150 File status okay; about to open data connection.
2022/01/19 00:52:06 >> 226 Closing data connection. Requested file action successful.
2022/01/19 00:52:11 << PORT [127,0,0,1,201,167]
2022/01/19 00:52:11 >> 200 Command okay.
2022/01/19 00:52:11 << RETR [main.go]
2022/01/19 00:52:11 >> 150 File status okay; about to open data connection.
2022/01/19 00:52:11 >> 226 Closing data connection. Requested file action successful.
2022/01/19 00:52:15 << CWD [image]
2022/01/19 00:52:15 >> 200 Command okay.
2022/01/19 00:52:22 << PORT [127,0,0,1,201,169]
2022/01/19 00:52:22 >> 200 Command okay.
2022/01/19 00:52:22 << LIST []
2022/01/19 00:52:22 >> 150 File status okay; about to open data connection.
2022/01/19 00:52:22 >> 226 Closing data connection. Requested file action successful.
2022/01/19 00:52:24 << TYPE [I]
2022/01/19 00:52:24 >> 200 Command okay.
2022/01/19 00:52:29 << PORT [127,0,0,1,201,172]
2022/01/19 00:52:29 >> 200 Command okay.
2022/01/19 00:52:29 << RETR [sample.png]
2022/01/19 00:52:29 >> 150 File status okay; about to open data connection.
2022/01/19 00:52:29 >> 226 Closing data connection. Requested file action successful.
2022/01/19 00:52:33 << QUIT []
2022/01/19 00:52:33 >> 221 Service closing control connection.
COMMENT_OUT
