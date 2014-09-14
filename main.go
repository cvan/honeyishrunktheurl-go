package main

import (
    "github.com/bitly/go-simplejson"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "os"
)

func main() {
    http.HandleFunc("/", Redirect)

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    log.Printf("Server running at 0.0.0.0:" + port)

    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
    // TODO: Consider moving outside out of this function, since `sites.json`
    // is reopened on every request. On the plus side, the binary doesn't need
    // to be recompiled every time the `sites.json` file changes.
    file, e := ioutil.ReadFile("./sites.json")
    if e != nil {
        log.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    sites, err := simplejson.NewJson(file)
    if err != nil {
        log.Fatalln(err)
    }

    key := r.URL.Path[1:]  // Strip the slash from the short URL.

    // Look up the destination from the short URL.
    destValue := sites.Get(key).MustString()
    if err != nil {
        log.Fatalln(err)
    }

    // If the short URL isn't registered, bail now.
    if destValue == "" {
        http.NotFound(w, r)
        return
    }

    // Parse the destination URL so we can extract its query-string
    // parameters.
    destUrl, err := url.Parse(destValue)
    if err != nil {
        log.Fatal(err)
    }

    // Parse the short URL so we can extract its query-string parameters.
    shortUrl, err := url.Parse(r.URL.String())  // Parse `/<key>`.
    if err != nil {
        log.Fatal(err)
    }

    params, _ := url.ParseQuery(destUrl.RawQuery)
    newParams, _ := url.ParseQuery(shortUrl.RawQuery)

    // Overwrite/add to the destination URL each query-string parameter found
    // in the short URL.
    for key, values := range newParams {
        for _, value := range values {
            params.Add(key, value)
        }
    }

    // Replace the original query-string parameters with the new URL-encoded
    // parameters.
    destUrl.RawQuery = params.Encode()

    log.Printf("%s ➡ %s\n", shortUrl, destUrl)

    // Finally, redirect.
    http.Redirect(w, r, destUrl.String(), http.StatusFound)
}
