package main

//import rl "github.com/gen2brain/raylib-go/raylib"
import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

var(
	openedSongs []Song

	done chan bool // handles when the song/playlist ends
	endStream beep.Streamer = beep.Callback(func() {done <- true})

	globalSampleRate beep.SampleRate
	globalSpeedCoefficient float32 = 1;
	controlStreamer *beep.Ctrl
)

type Song struct {
	fullPath string
	fileName string
	data *os.File
	streamer beep.StreamSeekCloser
	format beep.Format
}

func readFile(path string) *os.File {
	data, err := os.Open(path);
	if err != nil {
		log.Fatal(err);
	}
	// we don't close the file since streamer.Close() should handle this for us (hopefully)
	return data;
}

func openSong(path string, song *Song) {
	song.fullPath = path;
	song.fileName = strings.Split(song.fullPath, "/") [len(strings.Split(song.fullPath, "/")) - 1];
	song.data = readFile(path);

	var err error;
	song.streamer, song.format, err = mp3.Decode(song.data);
	if err != nil {
		log.Fatal(err);
	}
}

func resampleStreamer(streamer beep.StreamSeekCloser, format beep.Format) *beep.Resampler {
	resampler := beep.Resample(4, format.SampleRate*beep.SampleRate(globalSpeedCoefficient), globalSampleRate, streamer)
	return resampler;	
}

func shutUpAndPlay() {
	var streamer beep.Streamer;
	fmt.Println(len(openedSongs))
	if len(openedSongs) == 1 {
		controlStreamer = &beep.Ctrl{Streamer: openedSongs[0].streamer, Paused: false};
		speaker.Play(controlStreamer);
		return;
	}
	for i := 0; i < len(openedSongs)-1; i++ {
		streamerX := resampleStreamer(openedSongs[i].streamer, openedSongs[i].format);
		streamer = beep.Seq(streamerX, resampleStreamer(openedSongs[i+1].streamer, openedSongs[i+1].format));
		fmt.Println(i);
	}
	controlStreamer = &beep.Ctrl{Streamer: streamer, Paused: false};
	speaker.Play(beep.Seq(controlStreamer, endStream));
}

func musicProccessing() {
	for {
		select{
		case <-done:
			done <- false;
			return;
		case <-time.After(time.Second):
			speaker.Lock()
			fmt.Println(openedSongs[0].format.SampleRate.D(openedSongs[0].streamer.Position()).Round(time.Second))
			speaker.Unlock()
		}
	}
}

func main() {
	if len(os.Args) == 2 {
		openedSongs = append(openedSongs, Song{});
		openSong(os.Args[1], &openedSongs[0]);
		defer openedSongs[0].streamer.Close();
		globalSampleRate = openedSongs[0].format.SampleRate;
	} else if len(os.Args) > 2 {
		fmt.Println("gmusic: [USAGE]: gmusic [FILE]")
	}
	//openedSongs = append(openedSongs, Song{});
	//openSong("music2.mp3", &openedSongs[1]);
	defer openedSongs[1].streamer.Close();
	err := speaker.Init(globalSampleRate, int(globalSampleRate.N(time.Second/10)))
	if err != nil {
		log.Fatal(err);
	}

	go musicProccessing();
	shutUpAndPlay();
	renderLoop();
}

