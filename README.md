# 🎶 Groupie Tracker 🎶

Welcome to **Groupie Tracker**, your whimsical one-stop app for discovering artists, exploring their music journeys, and keeping tabs on their concerts and appearances! ✨ Whether you're a die-hard fan or a casual listener, this app will make sure you're always in tune. 🎤🎵

---
## 🌟 Features
### 📸 Artist Gallery

- Browse through a vibrant gallery of artists, complete with **photos** and **names**! 💃🕺
- Each card links you to a treasure trove of artist-specific details. 🔗

### 📜 Artist Details
- Dive into the world of your favorite artists with:
  - **Name & Photo** 🌟📷
  - **Band Members** 👩‍🎤👨‍🎤👩‍🎤
  - **Creation Date & First Album** 🗓️🎶
  - **Concert Locations & Dates** 📍📅
  - **Relations between locations and dates** 🗺️➡️🎤

### 🌈 Intuitive Design
- Seamlessly navigate through the app with user-friendly pages and elegant templates. ✨
- Optimized for both desktop and mobile viewing. 📱💻

---

## 💻 How It Works

### 🎯 Step 1: Fetching Artist Data
The app connects to a magical **API** 🪄 to retrieve detailed artist info, locations, dates, and relations.

### 🎯 Step 2: Dynamic Rendering
- The `index.html` page dynamically displays the artist gallery.
- Clicking on an artist leads to the `artist.html` page, showcasing their details.

### 🎯 Step 3: API Endpoints
- `/artists` - Fetches a list of artists.
- `/locations/<id>` - Retrieves concert locations for an artist.
- `/dates/<id>` - Retrieves concert dates.
- `/relation/<id>` - Shows relationships between dates and locations.

---

## 📂 Project Structure

```bash

/project-root
├── handlers/            # Contains the logic for handling requests and responses
│   ├── api.go           # API request handlers, typically for routing and interaction with the backend
│   ├── errors.go        # Handles error responses, such as 400, 404, and 500 status codes
│   ├── utils.go         # Utility functions that may be used across various handlers
├── static/              
│   ├── styles.css      
│   ├── favicon.ico          
├── structs/             # Structs for the app
│   ├── artists.go  
│   ├── dates.go 
│   ├── locations.go 
│   ├── relations.go       
├── templates/           # HTML templates
│   ├── 400.html       
│   ├── 404.html      
│   ├── 500.html       
│   ├── artist.html
│   ├── index.html
├── go.mod               # Go module file, used for managing dependencies
├── main.go              # Entry point of the application, initializes the server and routes
├── README.md   


   

```



## 🚀 Getting Started

# 📦 Prerequisites

- Go language installed on your device 🐹

- An internet connection to fetch API data. 🌐

## ⚡ Installation

# Clone this repository
``` bash
$ git clone https://platform.zone01.gr/git/mzargani/groupie-tracker.git

# Navigate to the project directory
$ cd groupie-tracker

# Run our application
$ go run .
```

## 🌍 Access the App

Open your browser and go to: http://localhost:8080

## 🌈 Whimsy Mode 🌈

"Music gives a soul to the universe, wings to the mind, flight to the imagination, and life to everything." - Plato ✨🎶


## 📞 Support

Need help? Have suggestions? Reach out to us:

- Gitea: astas and mzargani
- Discord: <a href="https://discordapp.com/users/780150798927134740" target="_blank">astas(Andriana)</a> - <a href="https://discordapp.com/users/768008367192014869" target="_blank"> mzargani(Maria)</a>

