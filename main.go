package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "encoding/json"
        "flag"
        "strings"
)

func main() {
        user := flag.String("u", "", "Twitch Username")
        clientid := flag.String("c", "", "Twitch API Client ID")
        bearerfile := flag.String("b", "", "Twitch API Bearer Token")
        game := flag.String("g", "", "Game substring")
        flag.Parse()

        arr_token, err := ioutil.ReadFile(*bearerfile)
        if err != nil {
                fmt.Print(err.Error())
        }
        str_token := "Bearer " + string(arr_token)
        REST := rest(*user, *clientid, str_token)
        parsedData := parse(REST)

        if len(parsedData.Data) > 0 {
                parsedData.Data[0].Title = strings.ReplaceAll(parsedData.Data[0].Title, "|", "/")
                if strings.Compare(*game, "none") != 0 {
                        if strings.Contains(parsedData.Data[0].GameName, *game) {
                                // User ONLINE with game filter and plays it
                                fmt.Printf("CRITICAL - %s ist live!\n\nTitel: %s \nGame: %s | viewer=%d \n", parsedData.Data[0].UserName, parsedData.Data[0].Title, parsedData.Data[0].GameName, parsedData.Data[0].ViewerCount)
                                os.Exit(2)
                        } else {
                                // User ONLINE with game filter but plays a different game (for viewer monitoring purpose)
                                fmt.Printf("OK - %s ist live!\n\nTitel: %s \nGame: %s | viewer=%d \n", parsedData.Data[0].UserName, parsedData.Data[0].Title, parsedData.Data[0].GameName, parsedData.Data[0].ViewerCount)
                                os.Exit(1)
                        }
                } else {
                        // User ONLINE without filter
                        fmt.Printf("CRITICAL - %s ist live!\n\nTitel: %s \nGame: %s | viewer=%d \n", parsedData.Data[0].UserName, parsedData.Data[0].Title, parsedData.Data[0].GameName, parsedData.Data[0].ViewerCount)
                        os.Exit(2)
                }
        } else {
                // User OFFLINE
                fmt.Printf("OK - %s ist offline. | viewer=0\n", *user)
                os.Exit(0)
        }
}

func rest(user string, clientid string, str_token string) (arr_resp []byte){
        client := &http.Client{}
        req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_login=" + user, nil)
        req.Header.Add("client-id", clientid)
        req.Header.Add("Authorization", str_token)
        resp, err := client.Do(req)
        if err != nil {
                fmt.Print(err.Error())
                os.Exit(3)
        }
        arr_resp, err = ioutil.ReadAll(resp.Body)
        if err != nil {
                fmt.Print(err.Error())
                os.Exit(3)
        }
        if (resp.StatusCode != 200) {
                s := string(arr_resp)
                e := strings.ReplaceAll(s, "\n", "")
                fmt.Printf("UNKNOWN - %d %s\n", resp.StatusCode, e)
                os.Exit(3)
        }
        return
}

func parse(arr_resp []byte) (JD JSON) {
        json.Unmarshal(arr_resp, &JD)
        return
}
