package cmd

import (
    "encoding/json"
    "errors"
    // "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/olekukonko/tablewriter"
)

type Password struct {
    Id int `json:"id"`
    Name string`json:"name"`
    AccessInfo string`json:"access_info"`
    Username string`json:"username"`
    Email string`json:"updated_on"`
    UpdatedOn string`json:"updated_on"`
}

type Passwords []Password

var (
    tag string
)

var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "A brief description of your command",
    Long: ``,
    RunE: func(cmd *cobra.Command, args []string) error {
        if len(args) == 0 {
            return errors.New("You need to provide a search query")
        }

        url := viper.GetString("base_url") + "/api/v4/passwords/search/" + strings.Join(args, " ") + "/page/1.json"

        client := &http.Client{
            Timeout: time.Second * 10,
        }

        req, reqError := http.NewRequest(http.MethodGet, url, nil)
        if reqError != nil {
            log.Fatal(reqError)
            return reqError
        }

        req.SetBasicAuth(viper.GetString("username"), viper.GetString("password"))
        req.Header.Add("Content-Type", `application/json; charset=utf-8`)
        req.Header.Set("User-Agent", "tmp vXXX")

        res, resError := client.Do(req)
        if resError != nil {
            log.Fatal(resError)
            return resError
        }

        body, bodyError := ioutil.ReadAll(res.Body)
        if bodyError != nil {
            log.Fatal(bodyError)
            return bodyError
        }

        passwords := Passwords{}
        jsonError := json.Unmarshal(body, &passwords)
        if jsonError != nil {
            log.Fatal(jsonError)
            return jsonError
        }

        table := tablewriter.NewWriter(os.Stdout)
        table.SetAlignment(tablewriter.ALIGN_LEFT)
        table.SetColumnSeparator(" ")
        table.SetBorder(false)
        table.SetRowLine(false)

        for _, password := range passwords {
            table.Append([]string{strconv.FormatInt(int64(password.Id), 10), password.Name, password.AccessInfo, password.Username})
        }

        table.Render()

        return nil
    },
}

func init() {
    RootCmd.AddCommand(searchCmd)

    searchCmd.Flags().StringVar(&tag, "tag", "", "Filter search results by the given tag")
}
