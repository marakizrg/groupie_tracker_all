# 🎶 Groupie Tracker  

## Features  

### Artist Gallery

- Browse through a vibrant gallery of artists, complete with **photos** and **names**!  
- Each card links you to a treasure trove of artist-specific details. 

###  Artist Details
- Dive into the world of your favorite artists with:
  - **Name & Photo** 
  - **Band Members** 
  - **Creation Date & First Album** 
  - **Concert Locations & Dates** 
  - **Relations between locations and dates** 

###  Intuitive Design
- Seamlessly navigate through the app with user-friendly pages and elegant templates. 

---

## 💻 How It Works

###  Step 1: Fetching Artist Data
The app connects to a magical **API** 🪄 to retrieve detailed artist info, locations, dates, and relations.

###  Step 2: Dynamic Rendering
- The `index.html` page dynamically displays the artist gallery.
- Clicking on an artist leads to the `artist.html` page, showcasing their details.

###  Step 3: API Endpoints
- `/artists` - Fetches a list of artists.
- `/locations/<id>` - Retrieves concert locations for an artist.
- `/dates/<id>` - Retrieves concert dates.
- `/relation/<id>` - Shows relationships between dates and locations.

---


## 🚀 Getting Started

#  Prerequisites

- Go language installed on your device 

- An internet connection to fetch API data. 



##  Access the App

Open your browser and go to: http://localhost:8080


