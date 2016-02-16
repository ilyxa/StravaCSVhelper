package main
import (
	"fmt"
        "github.com/strava/go.strava"
	"os"
	"flag"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

var accessToken string
var filef string

func main() {

	flag.StringVar(&accessToken,"token","thisisbigathletetokengottenfromstravaspi","Access Token")
	flag.StringVar(&filef,"file","activities.csv","CSV File Location")
	flag.Parse()

	const trip_n = 100
	fmt.Printf("Connecting to Strava %s...\n",accessToken)
	client := strava.NewClient(accessToken)
	atlSrv:= strava.NewCurrentAthleteService(client)
	actList,err := atlSrv.ListActivities().Page(1).PerPage(trip_n).Do()

        if err != nil {
                println(err)
		fmt.Printf("service error: %v", err)
		check(err)
        }
	fmt.Printf("Opening file %s...\n", filef)
	f, err := os.Create(filef)
	check(err)
	defer f.Close()
	fmt.Fprintf(f,"Num(Id),Start Date,Avg Cadence(rpm),Moving Time(min),Distance(km),Avg speed(km/h)\n")
	fmt.Printf("Writing ")
	var tot_n float64 = 0
	var cad_n float64 = 0
	var tot_cad float64 = 0
	var tot_spd float64 = 0
	var tot_dst float64 = 0
	for _, a := range actList {
		if a.MovingTime/60 > 10 {
			tot_cad+=a.AverageCadence
			tot_spd+=a.Distance/1000
			tot_dst+=a.AverageSpeed*3.6

			fmt.Printf(".")
			fmt.Fprintf(f,"<a href=https://www.strava.com/activities/%d>%3.0f</a>,%s,%3.1f,%v,%2.1f,%3.2f\n", a.Id, tot_n, a.StartDateLocal, a.AverageCadence, a.MovingTime/60, a.Distance/1000, a.AverageSpeed*3.6)

			check(err)
			if a.AverageCadence > 0 && a.AverageCadence < 150 {
				cad_n++
			}
                        tot_n++

		}
	}
	fmt.Printf("\nFinish\n")
	fmt.Fprintf(f,"<a href=/>999999</a>,<strong>Average %3.0f trip stats</strong>,<strong>%3.1f</strong>,,<strong>%3.2f</strong>,<strong>%3.2f</strong>\n",tot_n,tot_cad/cad_n,tot_spd/tot_n,tot_dst/tot_n)
	f.Sync()
}
