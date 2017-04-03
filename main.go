package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port = os.Getenv("PORT")
)

func AnnounceHandler(svc *polly.Polly) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		announce := r.URL.Query().Get("a")

		if announce == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "No accouncement included. Add ?a=<message>")
			return
		}

		params := &polly.SynthesizeSpeechInput{
			OutputFormat: aws.String("mp3"),
			Text:         aws.String(announce),
			VoiceId:      aws.String("Brian"),
			SampleRate:   aws.String("22050"),
			TextType:     aws.String("text"),
		}
		resp, err := svc.SynthesizeSpeech(params)
		if err != nil {
			log.Fatalf("Failed to synthesize speach: %v", err)
		}

		fmt.Println("Resp: ", resp)

		f, err := os.Create("./out.mp3")
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer f.Close()

		audio, err := ioutil.ReadAll(resp.AudioStream)
		if err != nil {
			log.Fatalf("Failed to read audio stream: %v", err)
		}

		_, err = f.Write(audio)
		if err != nil {
			log.Fatalf("Failed to write out to file: %v", err)
		}

		err = f.Sync()
		if err != nil {
			log.Fatalf("Failed to sync file output: %v", err)
		}

		fmt.Fprintf(w, announce)
	}

}

func main() {
	// must set env variables
	// * Access Key ID:     AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY
	// * Secret Access Key: AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY
	config := aws.NewConfig()
	sess := session.Must(session.NewSession(config.WithCredentials(credentials.NewEnvCredentials()).WithRegion("us-east-1")))
	svc := polly.New(sess)

	http.HandleFunc("/", AnnounceHandler(svc))

	log.Printf("Started announce server at http://localhost:%v\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
