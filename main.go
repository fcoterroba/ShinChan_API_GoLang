package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv"
)

type Personaje struct {
    Nombre string `json:"personaje"`
    Edad int `json:"edad"`
    Imagen string `json:"imagen"`
}

func main() {
    file, err := os.Open("characters.json")
    if err != nil {
        fmt.Println(err)
        return
    }

    var personajes []Personaje
    err = json.NewDecoder(file).Decode(&personajes)
    if err != nil {
        fmt.Println(err)
        return
    }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/characters", func(w http.ResponseWriter, r *http.Request) {
        idStr := r.URL.Query().Get("id")

        if idStr == "" {
            data, err := json.Marshal(personajes)
            if err != nil {
                fmt.Println(err)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.Write(data)
            return
        }

        id, err := strconv.Atoi(idSt
        if err != nil {
            fmt.Println(err)
            return
        }

        if id < 0 || id >= len(personajes) {
            fmt.Fprintf(w, "Error: index %d out of range", id)
            return
        }

        data, err := json.Marshal(personajes[id])
        if err != nil {
            fmt.Println(err)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(data)
    })

    http.ListenAndServe(":8080", nil)
}