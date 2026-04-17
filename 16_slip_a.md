\# 📁 Step 1: Create folder structure



```

rectapp/

&#x20;├── main.go

&#x20;└── rect/

&#x20;     └── rect.go

```



\---



\# 📦 Step 2: Create the package (rect/rect.go)



```go 

package rect



// Function to calculate area of rectangle

func Area(length, breadth int) int {

&#x09;return length \* breadth

}

```



✔ This is your \*\*user-defined package\*\*



\---



\# 🧠 Step 3: Create main program (main.go)



```go 

package main



import (

&#x09;"fmt"

&#x09;"rectapp/rect"

)



func main() {

&#x09;var l, b int



&#x09;fmt.Print("Enter length: ")

&#x09;fmt.Scan(\&l)



&#x09;fmt.Print("Enter breadth: ")

&#x09;fmt.Scan(\&b)



&#x09;area := rect.Area(l, b)



&#x09;fmt.Println("Area of rectangle:", area)

}

```



\---



\# ⚙️ Step 4: Run the program



Open terminal inside `rectapp` folder:



```bash

go run main.go

```





