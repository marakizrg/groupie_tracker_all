// Handle search input (what happens as the user types in the search bar)
document.getElementById("search-bar").addEventListener("input", function() {
    if (this.value.trim().length === 0) {
        resetToDefaultView();  // If the input is empty, reset to default view
    }
    
    const query = this.value.trim();  // Get the trimmed input value
    const suggestionsDiv = document.getElementById("suggestions");  // Get the suggestions div
    
    // If the query is empty, clear the suggestions and search results
    if (query.length === 0) {
        suggestionsDiv.innerHTML = "";
        document.getElementById("search-results").innerHTML = "";
        document.querySelector(".default-artists").classList.remove("hidden");
    }

    // Fetch suggestions based on the query
    fetch(`/search?q=${encodeURIComponent(query)}`)
        .then(response => response.json())  // Parse the JSON response
        .then(data => {
            suggestionsDiv.innerHTML = "";  // Clear previous suggestions
            
            let currentType = "";  // To keep track of the current type (artist, album, etc.)
            
            // Loop through each suggestion and display it
            data.forEach(result => {
                // If the result type is different from the previous, create a new header
                if (result.type !== currentType) {
                    const header = document.createElement("div");
                    header.className = "suggestion-header";
                    header.textContent = result.type.toUpperCase();
                    suggestionsDiv.appendChild(header);
                    currentType = result.type;
                }

                // Create a link for the suggestion and add it to the suggestions list
                const suggestionLink = document.createElement("a");
                suggestionLink.className = "suggestion-link";
                suggestionLink.href = `/artist/${result.artistId || result.id}`;
                suggestionLink.textContent = `${result.text} (${result.artistName})`;
                
                suggestionsDiv.appendChild(suggestionLink);
            });
        });
});

// Function to display all search results (called after a search is made)
function displayAllResults(results) {
    // Get the element where search results will be displayed
    const resultsDiv = document.getElementById("search-results");
    
    // Get the element for the default artists section
    const defaultArtists = document.querySelector(".default-artists");
    
    // Clear any previous results
    resultsDiv.innerHTML = "";
    
    // Hide the default artists section
    defaultArtists.classList.add("hidden");

    // If no results are found, display a message and a button to reset to the default view
    if (!results || results.length === 0) {
        resultsDiv.innerHTML = `
            <div class="no-results-container">
                <p class="no-results">No results found for "${document.getElementById("search-bar").value}"</p>
                <button onclick="resetToDefaultView()" class="reset-button">
                    Show All Artists
                </button>
            </div>
        `;
        return;
    }

    // Group results by their 'type' (artist, album, etc.)
    const groupedResults = results.reduce((groups, result) => {
        const type = result.type;  // Type of the result (e.g., 'artist')
        if (!groups[type]) groups[type] = [];  // If group for this type doesn't exist, create it
        groups[type].push(result);  // Add the current result to the group
        return groups;
    }, {});

    // Loop through each group of results and display them
    for (const [type, items] of Object.entries(groupedResults)) {
        const section = document.createElement("div");
        section.className = "results-section";
        
        // Create a header for the group (e.g., 'Artists', 'Albums')
        const header = document.createElement("h3");
        header.textContent = `${type.charAt(0).toUpperCase() + type.slice(1)}s`;
        section.appendChild(header);

        // If the group is of type 'artist', display the results in a grid layout
        if (type === "artist") {
            const grid = document.createElement("div");
            grid.className = "container";
            
            // Create a card for each artist and add it to the grid
            items.forEach(item => {
                const card = document.createElement("div");
                card.className = "card";
                
                const link = document.createElement("a");
                link.href = `/artist/${item.id}`;  // Link to the artist's page
                link.className = "artist-link";
                
                const img = document.createElement("img");
                img.src = item.image; // Artist's image
                img.alt = item.name || item.artistName;  // Alt text for the image
                img.className = "artist-image";
                
                const title = document.createElement("h2");
                title.textContent = item.name || item.artistName;  // Artist's name
                
                link.appendChild(img);  // Add the image and title to the link
                link.appendChild(title);
                card.appendChild(link);  // Add the link to the card
                grid.appendChild(card);  // Add the card to the grid
            });
            
            section.appendChild(grid);  // Add the grid to the section
        } else {  // For other types (e.g., albums), display the results in a list
            const list = document.createElement("ul");
            list.className = "results-list";
            
            // Create a list item for each result
            items.forEach(item => {
                const listItem = document.createElement("li");
                const artistId = item.artistId || item.id;  // Use artistId or id for the link
                
                listItem.innerHTML = `
                    <a href="/artist/${artistId}" class="result-link">
                        <strong>${item.text}</strong> (${item.artistName})
                    </a>
                `;
                list.appendChild(listItem);  // Add the list item to the list
            });
            
            section.appendChild(list);  // Add the list to the section
        }
        
        // Add the section to the results display
        resultsDiv.appendChild(section);
    }
}

// Handle Enter key press in the search bar (initiates a search)
document.getElementById("search-bar").addEventListener("keydown", function(event) {
    if (event.key === "Enter") {
        event.preventDefault();  // Prevent default Enter key action (form submission)
        const query = this.value.trim();  // Get the trimmed input value
        
        if (query.length === 0) {
            resetToDefaultView();  // If the input is empty, reset to default view
            return;
        }

        // Fetch results based on the query and display them
        fetch(`/search?q=${encodeURIComponent(query)}`)
            .then(response => response.json())  // Parse the JSON response
            .then(data => displayAllResults(data))  // Display the results
            .catch(error => {
                console.error("Error:", error);  // Log any error that occurs during fetch
                document.getElementById("search-results").innerHTML = `
                    <div class="no-results-container">
                        <p class="no-results">Search failed</p>
                        <button onclick="resetToDefaultView()" class="reset-button">
                            Show All Artists
                        </button>
                    </div>
                `;
            });
    }
});

// Function to reset the search view to its default state
function resetToDefaultView() {
    // Clears the search bar value
    document.getElementById("search-bar").value = "";
    
    // Clears the search results and suggestions displayed on the page
    document.getElementById("search-results").innerHTML = "";
    document.getElementById("suggestions").innerHTML = "";
    
    // Shows the default artists section again (hidden earlier in the displayAllResults function)
    document.querySelector(".default-artists").classList.remove("hidden");
}
