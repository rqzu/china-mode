package china_mode

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/reujab/wallpaper"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func StartChinaMode() {
	err := wallpaper.SetFromURL("https://i.imgflip.com/3vsifd.jpg")
	if err != nil {
		return
	}
	out, err := os.Create("C:/Windows/Temp/lol.mp3")
	if err != nil {
		log.Println(err)
		return
	}
	defer out.Close()
	resp, err := http.Get("https://cdn.letoa.me/lol.mp3")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}

	f, _ := os.Open("C:/Windows/Temp/lol.mp3")
	defer f.Close()

	// Decode the audio file
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize the speaker
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	// Play the audio
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait until the audio is finished playing
	<-done
}
