package tweeter

import (
    "log"
    core "../services"
)

func main() {
    client := core.GetTwitterClient()

    tweet, _, _ := client.Statuses.Update("testing tweet from other package", nil)
    log.Print("Posted Tweet: ", tweet.Text)
}