// // package main

// // import (
// //     "bufio"
// //     "fmt"
// //     "os"
// //     "strings"
// // )

// // func main() {
// //     reader := bufio.NewReader(os.Stdin)

// //     fmt.Print("What is your name? ")

// //     name, _ := reader.ReadString('\n')
// //     name = strings.TrimSpace(name)

// //     fmt.Printf("Hello, %s! Welcome to the Go Greeter CLI.\n", name)
// // }


// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const version = "1.0.0"

// func main() {
// 	if len(os.Args) < 2 {
// 		printHelp()
// 		return
// 	}

// 	command := os.Args[1]

// 	switch command {
// 	case "hello":
// 		helloCommand(os.Args[2:])
// 	case "bye":
// 		byeCommand(os.Args[2:])
// 	case "version":
// 		fmt.Println("Greeter CLI version", version)
// 	default:
// 		fmt.Printf("Unknown command: %s\n\n", command)
// 		printHelp()
// 	}
// }

// func helloCommand(args []string) {
// 	name := flag.NewFlagSet("hello", flag.ExitOnError)
// 	user := name.String("name", "friend", "Your name")
// 	loud := name.Bool("loud", false, "Shout the greeting")

// 	name.Parse(args)

// 	greeting := fmt.Sprintf("Hello, %s!", *user)
// 	if *loud {
// 		greeting = strings.ToUpper(greeting)
// 	}
// 	fmt.Println(greeting)
// }

// func byeCommand(args []string) {
// 	name := flag.NewFlagSet("bye", flag.ExitOnError)
// 	user := name.String("name", "friend", "Your name")

// 	name.Parse(args)

// 	fmt.Printf("Goodbye, %s!\n", *user)
// }

// func printHelp() {
// 	fmt.Println("Greeter CLI — Advanced Version")
// 	fmt.Println()
// 	fmt.Println("Usage:")
// 	fmt.Println("  greeter <command> [--flags]")
// 	fmt.Println()
// 	fmt.Println("Commands:")
// 	fmt.Println("  hello     Greet someone")
// 	fmt.Println("  bye       Say goodbye")
// 	fmt.Println("  version   Show version info")
// 	fmt.Println()
// 	fmt.Println("Examples:")
// 	fmt.Println("  greeter hello --name Brian")
// 	fmt.Println("  greeter hello --name Brian --loud")
// 	fmt.Println("  greeter bye --name Brian")
// }

// main.go
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Greeting represents a single greeting entry
type Greeting struct {
	Name     string
	Message  string
	Language string
	Time     string
}

// Shared in-memory store for greetings (last 5 only)
var (
	greetings []Greeting
	mu        sync.RWMutex
)

// HomePageData sent to template
type HomePageData struct {
	Greetings []Greeting
}

// Language map
var langGreetings = map[string]string{
	"en": "Hello",
	"es": "¡Hola",
	"fr": "Bonjour",
	"de": "Hallo",
	"ja": "こんにちは",
	"hi": "नमस्ते",
}

func init() {
	// Preload sample greeting
	greetings = []Greeting{{
		Name:     "Guest",
		Message:  "Hello",
		Language: "en",
		Time:     time.Now().Format("15:04:05"),
	}}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/api/greetings", apiGreetingsHandler)

	fmt.Println("🚀 Go Greetings App running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>Go Greetings ✨</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(120deg, #89f7fe 0%, #66a6ff 100%);
            min-height: 100vh;
            padding: 20px;
            color: #2c3e50;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background: rgba(255, 255, 255, 0.85);
            border-radius: 20px;
            padding: 30px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.15);
        }
        h1 {
            text-align: center;
            margin-bottom: 20px;
            color: #2980b9;
        }
        .live-clock {
            text-align: center;
            font-size: 1.2rem;
            margin-bottom: 30px;
            color: #7f8c8d;
            font-weight: bold;
        }
        form {
            display: flex;
            flex-direction: column;
            gap: 15px;
            margin-bottom: 30px;
        }
        input, select, button {
            padding: 12px;
            font-size: 1rem;
            border: 2px solid #bdc3c7;
            border-radius: 10px;
        }
        button {
            background: #3498db;
            color: white;
            cursor: pointer;
            font-weight: bold;
            transition: background 0.3s;
        }
        button:hover {
            background: #2980b9;
        }
        .greeting-history h2 {
            margin-bottom: 15px;
            color: #2c3e50;
        }
        .history-item {
            background: #ecf0f1;
            padding: 12px;
            margin-bottom: 10px;
            border-radius: 8px;
            display: flex;
            justify-content: space-between;
        }
        .history-item .msg { font-weight: bold; color: #2980b9; }
        .history-item .meta { color: #7f8c8d; font-size: 0.9rem; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🌍 Multilingual Greetings</h1>
        
        <div class="live-clock" id="clock">--:--:--</div>

        <form id="greetForm" method="POST" action="/greet">
            <input type="text" id="name" name="name" placeholder="Enter your name" required />
            <select name="lang" id="lang">
                <option value="en">English</option>
                <option value="es">Spanish</option>
                <option value="fr">French</option>
                <option value="de">German</option>
                <option value="ja">Japanese</option>
                <option value="hi">Hindi</option>
            </select>
            <button type="submit">Send Greeting 💌</button>
        </form>

        <div class="greeting-history">
            <h2>Recent Greetings</h2>
            {{range .Greetings}}
            <div class="history-item">
                <span class="msg">{{.Message}} {{.Name}}!</span>
                <span class="meta">{{.Language}} • {{.Time}}</span>
            </div>
            {{end}}
        </div>
    </div>

    <script>
        // Live clock
        function updateClock() {
            const now = new Date();
            document.getElementById('clock').textContent = 
                now.toLocaleTimeString();
        }
        setInterval(updateClock, 1000);
        updateClock();

        // Form submission with validation
        document.getElementById('greetForm').addEventListener('submit', function(e) {
            const name = document.getElementById('name').value.trim();
            if (!name) {
                e.preventDefault();
                alert('Please enter your name!');
                return false;
            }
        });
    </script>
</body>
</html>`

	data := HomePageData{
		Greetings: getGreetings(),
	}

	t, err := template.New("home").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	lang := r.FormValue("lang")

	if name == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	greeting, ok := langGreetings[lang]
	if !ok {
		greeting = "Hello"
		lang = "en"
	}

	// Add to history (thread-safe)
	mu.Lock()
	greetings = append([]Greeting{{
		Name:     name,
		Message:  greeting,
		Language: lang,
		Time:     time.Now().Format("15:04:05"),
	}}, greetings...)
	// Keep only last 5
	if len(greetings) > 5 {
		greetings = greetings[:5]
	}
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func apiGreetingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getGreetings())
}

func getGreetings() []Greeting {
	mu.RLock()
	defer mu.RUnlock()
	// Return a copy to avoid race conditions
	result := make([]Greeting, len(greetings))
	copy(result, greetings)
	return result
}
