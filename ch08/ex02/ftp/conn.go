package ftp

import "net"

// FTP接続のためのコネクションの状態を保持する
type Conn struct {
	conn     net.Conn
	dataType dataType  // データタイプ
	dataPort *dataPort // データポート
	rootDir  string    // ルートディレクトリ
	workDir  string    // ワーキングディレクトリ
}

// FTP接続のためのコネクションを返す
func NewConn(conn net.Conn, rootDir string) *Conn {
	return &Conn{
		conn:    conn,
		rootDir: rootDir,
		workDir: "/",
	}
}