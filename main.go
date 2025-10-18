// =============================================================
// 🧩 API Key Scanner v2.0 by HeJo-1
// -------------------------------------------------------------
// 🔍 Description:
//    Scans given URLs for potential exposed API keys (Google, AWS, GitHub, Slack).
//    Supports Turkish 🇹🇷 and English 🇬🇧 language options.
//    Beautiful colored console output for an aesthetic experience.
//
// ⚙️ Usage:
//    go run main.go --lang en https://example.com
//    go run main.go --lang tr https://ornek.com
// =============================================================

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/fatih/color"
)

// =============================================================
// 🎨 Language Packs
// =============================================================
var messages = map[string]map[string]string{
	"tr": {
		"banner":          "🚀 API Anahtar Tarayıcı v2.0 - HeJo-1 Tarafından",
		"scanning":        "🔍 Taranıyor: %s",
		"url_failed":      "❌ URL alınamadı %s: %v",
		"read_failed":     "⚠️  İçerik okunamadı %s: %v",
		"found_key":       "✅ Potansiyel %s Anahtarı Bulundu: %s",
		"found_item":      "   ➜ %s",
		"not_found":       "😕 Anahtar bulunamadı: %s",
		"done":            "🎯 Tarama tamamlandı. Sonuçlar 'found_keys.txt' dosyasına kaydedildi.",
		"usage":           "💡 Kullanım: go run main.go [--lang tr|en] <url1> <url2> ...",
		"no_url":          "⚠️  Lütfen taranacak bir veya daha fazla URL belirtin.",
		"file_error":      "📁 Dosya açılamadı: %v",
		"write_error":     "📝 Dosyaya yazılamadı: %v",
	},
	"en": {
		"banner":          "🚀 API Key Scanner v2.0 - Created by HeJo-1",
		"scanning":        "🔍 Scanning: %s",
		"url_failed":      "❌ Failed to fetch URL %s: %v",
		"read_failed":     "⚠️  Failed to read content %s: %v",
		"found_key":       "✅ Potential %s Key Found: %s",
		"found_item":      "   ➜ %s",
		"not_found":       "😕 No keys found: %s",
		"done":            "🎯 Scan complete. Results saved to 'found_keys.txt'.",
		"usage":           "💡 Usage: go run main.go [--lang tr|en] <url1> <url2> ...",
		"no_url":          "⚠️  Please specify one or more URLs to scan.",
		"file_error":      "📁 Failed to open file: %v",
		"write_error":     "📝 Failed to write to file: %v",
	},
}

// =============================================================
// 🧠 Global Variables
// =============================================================
var lang string

// Supported API key patterns
var apiRegex = map[string]string{
	"Google": `AIza[0-9A-Za-z\-_]{35}`,
	"AWS":    `AKIA[0-9A-Z]{16}`,
	"GitHub": `[a-zA-Z0-9_-]+@github\.com:[a-zA-Z0-9_-]+\/[a-zA-Z0-9_-]+\.git`,
	"Slack":  `xox[baprs]-[0-9a-zA-Z]{10,48}`,
}

// =============================================================
// 🗣️ Language Helper
// =============================================================
func t(key string) string {
	if val, ok := messages[lang][key]; ok {
		return val
	}
	return key
}

// =============================================================
// 🔍 Core Scanner Function
// =============================================================
func findApiKeys(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	cyan := color.New(color.FgCyan).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()
	red := color.New(color.FgRed).PrintfFunc()

	cyan(t("scanning")+"\n", url)

	resp, err := http.Get(url)
	if err != nil {
		red(t("url_failed")+"\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		red(t("read_failed")+"\n", url, err)
		return
	}

	content := string(body)
	found := false

	for key, pattern := range apiRegex {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(content, -1)

		if len(matches) > 0 {
			found = true
			green(t("found_key")+"\n", key, url)
			for _, match := range matches {
				yellow(t("found_item")+"\n", match)
				saveToFile(url, key, match)
			}
		}
	}

	if !found {
		yellow(t("not_found")+"\n", url)
	}
}

// =============================================================
// 💾 Save Found Keys
// =============================================================
func saveToFile(url, keyType, key string) {
	file, err := os.OpenFile("found_keys.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf(t("file_error"), err)
	}
	defer file.Close()

	line := fmt.Sprintf("URL: %s\nType: %s\nKey: %s\n\n", url, keyType, key)
	if _, err := file.WriteString(line); err != nil {
		log.Printf(t("write_error"), err)
	}
}

// =============================================================
// 🚀 Main Function
// =============================================================
func main() {
	flag.StringVar(&lang, "lang", "tr", "🌐 Dil seçimi / Language (tr/en)")
	flag.Parse()

	banner := color.New(color.FgMagenta, color.Bold)
	banner.Println("\n==============================================")
	banner.Println(messages[lang]["banner"])
	banner.Println("==============================================\n")

	if len(flag.Args()) < 1 {
		color.Red(t("no_url"))
		color.Yellow(t("usage"))
		return
	}

	urls := flag.Args()
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go findApiKeys(url, &wg)
	}

	wg.Wait()
	color.New(color.FgGreen, color.Bold).Println("\n" + t("done"))
}

