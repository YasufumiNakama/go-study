package ftp

type dataType int

const (
	ascii dataType = iota
	image
)

func (c *Conn) setDataType(args []string) {
	if len(args) == 0 {
		c.respond(status501)
	}

	switch args[0] {
	case "A":
		c.dataType = ascii
	case "I":
		c.dataType = image
	default:
		c.respond(status504)
		return
	}
	c.respond(status200)
}