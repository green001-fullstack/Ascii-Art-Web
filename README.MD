# 🎨 ASCII Art Web Generator

A web-based ASCII art generator built with **Go (Golang)**.
It allows users to enter text, choose a banner style, generate ASCII art, and download the result.

---

## 🚀 Features

* Convert text into ASCII art
* Multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Multi-line text support (`\n`)
* Download generated ASCII art as a `.txt` file
* Clean and animated UI using HTML & CSS
* Error handling for invalid routes and inputs

---

## 🛠️ Tech Stack

* **Backend:** Go (net/http, html/template)
* **Frontend:** HTML, CSS
* **Architecture:** MVC-like structure

---

## 📁 Project Structure

```
.
├── main.go
├── ascii/
│   ├── generator.go
│   └── banners/
│       ├── standard.txt
│       ├── shadow.txt
│       └── thinkertoy.txt
├── templates/
│   ├── index.html
│   └── error.html
├── static/
│   └── style.css
└── README.md
```

---

## ⚙️ Installation & Setup

### 1. Clone the repository

```
git clone https://github.com/green001-fullstack/Ascii-Art-Web.git
cd ascii-art-web
```

### 2. Run the server

```
go run main.go
```

### 3. Open in browser

```
http://localhost:8080
```

---

## 🧠 How It Works

### 1. User Input

* User enters text and selects a banner style from the UI.

### 2. Request Handling

* `/ascii` route handles POST request.
* Text is processed and split into lines.

### 3. ASCII Generation

* The `GenerateAscii` function:

  * Reads the selected banner file
  * Maps each character to ASCII art
  * Builds output line-by-line

### 4. Display

* Result is rendered inside a `<pre>` block.

### 5. Download

* `/download` route sends ASCII art as a `.txt` file.

---

## 🧩 Key Functions

### `GenerateAscii(text, banner string) string`

* Reads banner file
* Converts characters into ASCII art
* Returns formatted string

### `asciiHandler`

* Handles form submission
* Processes text and generates ASCII

### `downloadHandler`

* Allows users to download generated ASCII art

---

## ⚠️ Error Handling

* Invalid HTTP methods → `405 Method Not Allowed`
* Invalid routes → Custom `404` page
* Empty input → User-friendly message
* Missing banner file → Graceful error return

---

## 🎯 Example

### Input:

```
Hello
```

### Output:

```
 _   _          _ _       
| | | |   ___  | | | ___  
| |_| |  / _ \ | | |/ _ \ 
|  _  | |  __/ | | | (_) |
|_| |_|  \___| |_|_|\___/ 
```

---

## 📌 Future Improvements

* Add text alignment (left, center, right, justify)
* Add more banner styles
* Allow file upload input
* Add copy-to-clipboard feature
* Add coloring to the art

---

## 👨‍💻 Author

Developed by **Bolarinwa Emmanuel**

---

## 📄 License

This project is open-source and available under the MIT License.
