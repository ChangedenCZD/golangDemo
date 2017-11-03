package urlParse

import (
	"net/url"
	"fmt"
	"net"
)

var p = fmt.Println
/**url解析*/
func Run() {
	str := "https://www.changeden.net:1234/index.html?token=444444#footer"

	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}

	p(u.Scheme)
	if u.User != nil {
		p(u.User)
		p(u.User.Username())
		psd, _ := u.User.Password()
		p(psd)
	} else {
		p("User is nil")
	}

	p(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)

	p(u.Path)
	p(u.Fragment)
	p(u.RawQuery)
	query, _ := url.ParseQuery(u.RawQuery)
	p(query)
}
