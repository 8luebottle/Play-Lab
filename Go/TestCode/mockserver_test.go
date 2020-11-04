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

/*
	=== RUN   TestDownload
		TestDownload: mockserver_test.go:52: start testing content download  feature.
		TestDownload: mockserver_test.go:54: 	URL "http://127.0.0.1:59987" check status code "200"
	&{0xc000122be0 0xc000162100 {} 0x40edaa0 true false false false 0xc0000920c0 {0xc0000b2000 map[] false false} map[Content-Type:[application/xml]] true 0 -1 200 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc0000b0000 0} <?xml version="1.0" encoding="UTF-8"?>
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
	</rss>
		TestDownload: mockserver_test.go:63: 		check HTTP GET request ✓
		TestDownload: mockserver_test.go:76: 		 check status code is "200" ✓
	--- PASS: TestDownload (0.00s)
	PASS

	Process finished with exit code 0
*/
