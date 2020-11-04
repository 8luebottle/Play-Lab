// Reference : [Book] Go in Action
package TestCode

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Vim Commands</title>
	<description>Github : https://github.com/8luebottle/</description>
	<link>https://babytiger.netlify.app/</link>
	<item>
		<pubDate>Sunday, 30-Aug-20 00:00:00 UTC</pubDate>
		<title>Vim Commands</title>
		<description>
			<p>Normal Mode → Insert mode<br>
				<code>i</code> : insert<br>
				Enter insert mode
			</p>
		</description>
		<link>https://babytiger.netlify.app/posts/vim/</link>
	</item>
</channel>
</rss>`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Println(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("start testing content download  feature.")
	{
		t.Logf("\tURL \"%s\" check status code \"%d\"", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal(
					"\t\t check up HTTP GET request",
					ballotX, err,
				)
			}
			t.Log(
				"\t\tcheck HTTP GET request",
				checkMark,
			)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf(
					"\t\t check status code is \"%d\" %v %v",
					statusCode, ballotX, resp.StatusCode,
				)
			}
			t.Logf("\t\t check status code is \"%d\" %v", statusCode, checkMark)
		}
	}
}
