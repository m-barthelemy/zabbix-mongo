package main 


import (
    "strings"
    "gopkg.in/cavaliercoder/g2z.v3"
)

func main() {
    panic("THIS_SHOULD_NEVER_HAPPEN")
}

func Echo(request *g2z.AgentRequest) (string, error) {
    return strings.Join(request.Params, " "), nil
}

func init() {
    g2z.RegisterStringItem("mongo.echo", "Hello world!", Echo)
    g2z.RegisterStringItem("mongo.run", "", queryDB)

    g2z.RegisterDiscoveryItem("mongo.discover", "", discover)
}

