
# 📁 Step 1: Create folder structure

```
rectapp/
 ├── main.go
 └── rect/
      └── rect.go
```

---

# 📦 Step 2: Create the package (`rect/rect.go`)

```go
package rect

// Function to calculate area of rectangle
func Area(length, breadth int) int {
	return length * breadth
}
```

✔ This is your **user-defined package**

---

# 🧠 Step 3: Create main program (`main.go`)

```go
package main

import (
	"fmt"
	"rectapp/rect"
)

func main() {
	var l, b int

	fmt.Print("Enter length: ")
	fmt.Scan(&l)

	fmt.Print("Enter breadth: ")
	fmt.Scan(&b)

	area := rect.Area(l, b)

	fmt.Println("Area of rectangle:", area)
}
```

---

# ⚙️ Step 4: Run the program

Open terminal inside `rectapp` folder:

```bash
go mod init rectapp
go run main.go
```
