package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "github.com/joho/godotenv"
)

func main() {

    err := godotenv.Load()
    if err != nil {
    fmt.Println("Error loading .env file")
    return
    }

    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
    fmt.Println("API_KEY is missing from .env file")
    return
    }

    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("Usage: ./upicheck <phone number>")
        return
    }

    phoneNumber := args[0]
    upiExtensions := []string{"apl", "ybl", "oksbi", "okhdfcbank", "axl", "paytm", "ibl", "upi", "icici", "sbi", "kotak", "postbank", "axisbank", "okicici", "okaxis", "dbs", "barodampay", "idfcbank"}

    for _, extension := range upiExtensions {
        vpa := phoneNumber + "@" + extension

        url := "https://upi-verification.p.rapidapi.com/v3/tasks/sync/verify_with_source/ind_vpa"

        payload := strings.NewReader("{\n    \"task_id\": \"UUID\",\n    \"group_id\": \"UUID\",\n    \"data\": {\n        \"vpa\": \"" + vpa + "\"\n    }\n}")

        req, err := http.NewRequest("POST", url, payload)
        if err != nil {
            fmt.Println("Error creating request:", err)
            return
        }

        req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", apiKey)
        req.Header.Add("X-RapidAPI-Host", "upi-verification.p.rapidapi.com")

        res, err := http.DefaultClient.Do(req)
        if err != nil {
            fmt.Println("Error sending request:", err)
            return
        }

        defer res.Body.Close()

        // read the response body into a []byte
        bodyBytes, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
            return
        }

        // create a new buffer to store only the JSON data
        var buf bytes.Buffer
        if err := json.Compact(&buf, bodyBytes); err != nil {
            fmt.Println("Error compacting JSON:", err)
            return
        }

        // print the JSON data as a string
        fmt.Println(buf.String())
    }
}

