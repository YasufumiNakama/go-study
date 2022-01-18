go run main.go -port 8000

<<COMMENT_OUT
(base) nakamayasufumi@NAKAMAnoMacBook-Pro ex02 % ./run_server.sh
2022/01/19 00:17:19 >> 220 Service ready for new user.
2022/01/19 00:17:21 << USER [user]
2022/01/19 00:17:21 >> 230 User user logged in, proceed.
2022/01/19 00:17:25 << EPRT [|2|::1|51147|]
2022/01/19 00:17:25 >> 502 Command not implemented.
2022/01/19 00:17:25 << LPRT [6,16,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,2,199,203]
2022/01/19 00:17:25 >> 502 Command not implemented.
2022/01/19 00:17:25 << LIST []
2022/01/19 00:17:25 >> 150 File status okay; about to open data connection.
2022/01/19 00:17:25 dial tcp: missing address
2022/01/19 00:17:25 >> 425 Can't open data connection.
COMMENT_OUT
