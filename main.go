package main

import (
	"encoding/json"
	"github.com/stvp/pager"
	"log"
	"net/http"
	"os"
	"time"
)

type SidekiqStats struct {
	Sidekiq struct {
		Processed      int `json:"processed"`
		Failed         int `json:"failed"`
		Busy           int `json:"busy"`
		Processes      int `json:"processes"`
		Enqueued       int `json:"enqueued"`
		Scheduled      int `json:"scheduled"`
		Retries        int `json:"retries"`
		Dead           int `json:"dead"`
		DefaultLatency int `json:"default_latency"`
	} `json:"sidekiq"`
	Redis struct {
		RedisVersion        string `json:"redis_version"`
		UptimeInDays        string `json:"uptime_in_days"`
		ConnectedClients    string `json:"connected_clients"`
		UsedMemoryHuman     string `json:"used_memory_human"`
		UsedMemoryPeakHuman string `json:"used_memory_peak_human"`
	}
}

func getStats() (*SidekiqStats, error) {
	log.Println("Getting stats")
	username := os.Getenv("HTTP_USERNAME")
	password := os.Getenv("HTTP_PASSWORD")
	response, err := http.Get("https://" + username + ":" + password + "@api.dispatch.me/sidekiq/stats")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	stats := &SidekiqStats{}
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(stats)
	return stats, err
}

func main() {
	pager.ServiceKey = os.Getenv("PAGERDUTY_KEY")

	log.Println("Starting monitoring loop")
	for {
		stats, err := getStats()
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%d messages enqueued\n", stats.Sidekiq.Enqueued)
			if stats.Sidekiq.Enqueued > 500 {
				pager.Trigger("Sidekiq queue backed up")
			}
		}

		time.Sleep(60 * time.Second)
	}
}
