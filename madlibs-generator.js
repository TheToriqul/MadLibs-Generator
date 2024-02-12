const fs = require('fs'); // Import the file system module

// Load the story from a file (replace "story.txt" with your actual file name)
const story = fs.readFileSync('story.txt', 'utf-8').toString();

// Create a set to store unique placeholder words
const words = new Set();

// Define the placeholder markers
const targetStart = "<";
const targetEnd = ">";

// Find placeholder words and add them to the set
let startOfWord = -1;
for (let i = 0; i < story.length; i++) {
    const char = story[i];

    if (char === targetStart) {
        startOfWord = i; // Mark the start of a potential placeholder
    } else if (char === targetEnd && startOfWord !== -1) {
        const word = story.slice(startOfWord, i + 1); // Extract the placeholder word
        words.add(word); // Add the unique word to the set
        startOfWord = -1; // Reset for the next placeholder
    }
}

// Create a dictionary to store user-provided words for placeholders
const answers = {};

// Prompt the user for words and fill the dictionary
for (const word of words) {
    const answer = prompt(`Enter a word for ${word}:`);
    answers[word] = answer;
}

// Replace placeholders with user-provided words
let updatedStory = story;
for (const word of words) {
    updatedStory = updatedStory.replaceAll(word, answers[word]);
}

// Print the completed story
console.log(updatedStory);
