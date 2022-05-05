/*

Does:
	All web vulnerabilitiy testers

Alot of this was previously broken tools moved to new modules, updated to the newest standard, and modernized
*/

package web

import (
	"fmt"
	"io/ioutil"
	httpreqconst "main/modg/constants/webconst"
	httprequeests "main/modg/requests"
	httpreqtypes "main/modg/types"
	"net/http"
	"strings"
)

func Find(i []string) []httpreqtypes.Result {
	for i, d := range i {
		httpreqconst.Request_limit <- d
		httpreqconst.Wg.Add(1)
		go func(i int, da string) {
			defer httpreqconst.Wg.Done()
			defer func() {
				<-httpreqconst.Request_limit
			}()
			response, e := http.Get(da)
			httpreqconst.Mutmod.Lock()
			if e == nil {
				body, err := ioutil.ReadAll(response.Body)
				if err == nil && len(body) != 0 {
					sb := string(body)
					results := CheckSinks(sb, da)
					httpreqconst.Content_results = append(httpreqconst.Content_results, results...)
				}
				response.Body.Close()
			}
			httpreqconst.Mutmod.Unlock()
		}(i, d)
	}
	httpreqconst.Wg.Wait()
	return httpreqconst.Content_results
}

func CheckSinks(response_body string, url string) []httpreqtypes.Result {
	var r []httpreqtypes.Result
	t1 := strings.ToLower(response_body)
	t2 := strings.ReplaceAll(t1, " ", "")
	for _, sink := range httpreqconst.DOM {
		if strings.Contains(t2, sink) {
			res := httpreqtypes.Result{
				Result_Sink: sink, Url: url}
			r = append(r, res)
		}
	}
	return r
}

func Call_sink(target []string) {
	results := Find(httprequeests.Remove_URL_vals(target))
	for _, k := range results {
		fmt.Println("<RR6> Sink-> Got sink on url ", k.Url)
		fmt.Printf("<RR6> Sink->  < %s > \n", k.Result_Sink)
	}
}
