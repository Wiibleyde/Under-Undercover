package logger

import (
	"config"
	"io"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/fatih/color"
	"github.com/gtuk/discordwebhook"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	TrafficLogger *log.Logger
	DebugLogger   *log.Logger
	RpLogger      *log.Logger
)

func getDate() string {
	dt := time.Now()
	return dt.Format("2006-01-02")
}

func InitLogger() {
	if _, err := os.Stat("logs/"); os.IsNotExist(err) {
		os.Mkdir("logs/", 0777)
	}
	file, err := os.OpenFile("logs/logs-"+getDate()+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multi := io.MultiWriter(file, color.Output)

	WarningLogger = log.New(multi, color.YellowString("WARNING: "), log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(multi, color.BlueString("INFO: "), log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multi, color.RedString("ERROR: "), log.Ldate|log.Ltime|log.Lshortfile)
	TrafficLogger = log.New(multi, color.WhiteString("TRAFFIC: "), log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(multi, color.GreenString("DEBUG: "), log.Ldate|log.Ltime|log.Lshortfile)
	RpLogger = log.New(multi, color.MagentaString("RP: "), log.Ldate|log.Ltime|log.Lshortfile)

	// Ajout des appels à LogToDiscord pour chaque type de log
	WarningLogger.SetOutput(io.MultiWriter(file, color.Output, discordWriter{level: "WARNING"}))
	InfoLogger.SetOutput(io.MultiWriter(file, color.Output, discordWriter{level: "INFO"}))
	ErrorLogger.SetOutput(io.MultiWriter(file, color.Output, discordWriter{level: "ERROR"}))
	// TrafficLogger.SetOutput(io.MultiWriter(file, color.Output, discordWriter{level: "TRAFFIC"}))
	// DebugLogger.SetOutput(io.MultiWriter(file, color.Output, discordWriter{level: "DEBUG"}))
	// RpLogger.SetOutput(io.MultiWriter(file, color.Output, discordWriter{level: "RP"}))
}

type discordWriter struct {
	level string
}

func (writer discordWriter) Write(bytes []byte) (int, error) {
	LogToDiscord(string(bytes), writer.level)
	return len(bytes), nil
}

func removeANSICodes(message string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(message, "")
}

func removeLevel(message string) string {
	re := regexp.MustCompile(`(WARNING|INFO|ERROR|DEBUG|TRAFFIC): `)
	return re.ReplaceAllString(message, "")
}

func removeDate(message string) string {
	re := regexp.MustCompile(`\d{4}/\d{2}/\d{2} `)
	return re.ReplaceAllString(message, "")
}

func removeTime(message string) string {
	re := regexp.MustCompile(`\d{2}:\d{2}:\d{2} `)
	return re.ReplaceAllString(message, "")
}

func removeShortFile(message string) string {
	re := regexp.MustCompile(`\w+.go:\d+ `)
	return re.ReplaceAllString(message, "")
}

func LogToDiscord(message string, level string) {
	message = removeANSICodes(message)
	message = removeLevel(message)
	message = removeDate(message)
	message = removeTime(message)
	message = removeShortFile(message)

	var username = "Panel"
	var content = message
	var url = config.GetConfig().Webhooks.LogsWebhook
	var logo = "https://panel.hope-rp.com/static/img/logo-hope.png"

	thumbnail := discordwebhook.Thumbnail{
		Url: &logo,
	}

	footer := discordwebhook.Footer{
		Text: &username,
	}

	color := "65280"
	switch level {
	case "INFO":
		color = "65280"
	case "WARNING":
		color = "16776960"
	case "ERROR":
		color = "16711680"
	case "DEBUG":
		color = "16776960"
	case "TRAFFIC":
		color = "16776960"
	case "RP":
		color = "1127128"
	}

	if len(content) > 1900 {
		content = content[:1900] + "..."
	}

	messageType := discordwebhook.Message{
		Username: &username,
		Embeds: &[]discordwebhook.Embed{
			{
				Title:       &level,
				Description: &content,
				Thumbnail:   &thumbnail,
				Color:       &color,
				Footer:      &footer,
			},
		},
	}

	err := discordwebhook.SendMessage(url, messageType)
	if err != nil {
		ErrorLogger.Printf("Error sending message to discord: %s", err)
	}
}
