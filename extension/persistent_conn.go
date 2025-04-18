package extension

import "net/http"

func ClosePersistentConn(req *http.Request) error {
	req.Header.Set("Connection", "close")
	return nil
}
