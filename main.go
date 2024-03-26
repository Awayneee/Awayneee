package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
    <title>Happy Birthday!</title>
    <style>
        body {
            background-color: white;
            color: black;
            text-align: center;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            display: none; /* Initially hide the container */
            background-color: pink;
            padding: 20px;
            width: 80%;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.5);
            position: relative;
        }
        .heart {
            color: red;
            animation: pulse 1s infinite;
            font-size: 20px;
            position: absolute;
        }
        @keyframes pulse {
            0% {
                transform: scale(0.9);
            }
            50% {
                transform: scale(1);
            }
            100% {
                transform: scale(0.9);
            }
        }
        .btn {
            background-color: black;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            transition: background-color 0.3s ease;
        }
        .btn:hover {
            background-color: #333;
        }
    </style>
</head>
<body>
    <button class="btn" onclick="openContainer()">Open</button>
    <div class="container" id="container"> <!-- Show the container when triggered -->
        <div class="hearts">
            <!-- Hearts will be dynamically added here -->
        </div>
        <div class="sentence">
            <h1>Dayı, doğum gününüz kutlu olsun.</h1>
            <p>Nice sağlıklı, mutlu, huzurlu yaşlarınız olsun.</p>
        </div>
        <div class="signature">
            <p>Ahmet Zeki Uslu</p>
        </div>
    </div>
    
    <script>
        function openContainer() {
            var container = document.getElementById("container");
            container.style.display = "block";
            var btn = document.querySelector(".btn");
            btn.style.display = "none"; // Hide the button
        }
    </script>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}



