package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nyxoy77/go-text-to-speech/speech"
)

const (
	ExitCodeOK              = 0
	ExitCodeParseFlagError  = 1
	ExitCodeValidateError   = 2
	ExitCodeInternalError   = 3
	ExitCodeOutputFileError = 4
)

type CLI struct {
	ErrStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("google-text-to-speech", flag.ContinueOnError)

	var (
		text, voice, out string
		rate, pitch      float64
	)

	flags.StringVar(&text, "text", "", "text to speech")
	flags.StringVar(&voice, "voice", "stand-a", "speaker's voice name")
	flags.Float64Var(&rate, "rate", 1.00, "speech rate (0.25)")
	flags.Float64Var(&pitch, "pitch", 0.00, "speaking pitch")
	flags.StringVar(&out, "O", "", "output audio file(support format of the audio: Linear16 , MP3)")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeParseFlagError
	}

	opt, err := makeSpeechOpt(text, voice, out, rate, pitch)

	if err != nil {
		fmt.Print(cli.ErrStream, err)
		return ExitCodeValidateError
	}
	ctx := context.Background()
	speaker, err := speech.NewSpeechClient(ctx)
	if err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeInternalError
	}

	b, err := speaker.Run(ctx, speech.NewRequest(text, opt))
	if err != nil {
		fmt.Print(cli.ErrStream, err)
		return ExitCodeInternalError
	}

	if err = os.WriteFile(out, b, 0644); err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeOutputFileError
	}

	fmt.Printf("mp3 created successfully at : %s\n", out)
	return ExitCodeOK

}

func makeSpeechOpt(text, voice, out string, rate, pitch float64) (*speech.SpeechOption, error) {
	if text == "" {
		return nil, fmt.Errorf("empty text")
	}

	var voiceName string
	switch v := strings.ToLower(voice); v {
	case "stand-a":
		voiceName = speech.VoiceStandardA
	case "stand-b":
		voiceName = speech.VoiceStandardB
	case "stand-c":
		voiceName = speech.VoiceStandardC
	case "stand-d":
		voiceName = speech.VoiceStandardD
	case "wave-a":
		voiceName = speech.VoiceWaveNetA
	case "wave-b":
		voiceName = speech.VoiceWaveNetB
	case "wave-c":
		voiceName = speech.VoiceWaveNetC
	case "wave-d":
		voiceName = speech.VoiceWaveNetD
	default:
		return nil, fmt.Errorf("unknow voice name: %v", v)

	}

	if 0.25 > rate || rate > 4.0 {
		return nil, fmt.Errorf("valid speaking_rate is between 0.25 and 4.0(rate %g)", rate)
	}

	if -0.20 > pitch || pitch > 20.00 {
		return nil, fmt.Errorf("valid pitch is between -20.0 and 20.0 (pitch: %g)", pitch)
	}

	switch ext := strings.ToLower(filepath.Ext(out)); ext {
	case ".wav":
		return &speech.SpeechOption{
			LanguageCode:      "ja-JP",
			VoiceName:         voiceName,
			AudioEncoding:     speech.AudioEncoding_LINEAR16,
			AudioSpeakingRate: rate,
			AudioPitch:        pitch,
		}, nil
	case ".mp3":
		return &speech.SpeechOption{
			LanguageCode:      "ja-JP",
			VoiceName:         voiceName,
			AudioEncoding:     speech.AudioEncoding_MP3,
			AudioSpeakingRate: rate,
			AudioPitch:        pitch,
		}, nil

	default:
		return nil, fmt.Errorf("unknown extendsion (out: %s)", out)
	}
}
