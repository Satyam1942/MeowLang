# 🐱 Meow Lang

Meow Lang is an **interpreted** toy programming language with a cat-themed syntax. The interpreter is written in **Go**. The **sample.txt** file contains the sample code. The **interpreter.go** file has the interpreter
written. Currently you have to change the file name of your code in the **main function** of interpreter and run the **interpreter.go** itself. Find the entire documentation regarding syntax of language [here](https://docs.google.com/document/d/1v5oVxh43-RfygwwAfl4d848zmCxvvGUfmjWZL2DoPMQ/edit?tab=t.2prreufkgq0p). 

---

## 🏁 Program Structure
A Meow Lang program must:
- **Start with** `Hello Kitty!`
- **End with** `Bye Kitty!`
- Anything outside this will be **ignored**.

### Example:
```meow
Hello Kitty!

// This is a comment
purr "Hello, world!"

Bye Kitty!
```

---

## 📝 Comments
- Only **single-line comments** are supported.
- Syntax:
  ```meow
  // This is a comment
  ```

---

## 🖨️ Print Statements
Meow Lang provides two print functions:

### Print in the same line:
```meow
purr b 
purr "Hello"
```
**Output:**
```
b"Hello"
```

### Print on a new line:
```meow
purrln b
purr "Hello"
```
**Output:**
```
b
"Hello"
```

---

## 📦 Variables
- Only **Integer** and **String** assignments are supported.
- Syntax:
  ```meow
  meow a = 5
  meow b = "Hello"
  ```

---

## ➕ Arithmetic Operations
Supports basic **binary** mathematical operations:

| Operation | Symbol |
|-----------|--------|
| Addition  | `+`    |
| Subtraction | `-`  |
| Multiplication | `*` |
| Division  | `/`    |
| Bitwise AND | `&`  |
| Bitwise OR | `\|`   |
| Bitwise XOR | `^`  |

---

## 🔄 Conditionals
- Uses `meowIf` for **if statements**.
- Uses `otherwiseScratch` for **else statements**.
- **No** `if-else-if` chaining allowed.
- **Only one statement** allowed inside `if` or `else`.
- Variables defined inside `if` are accessible outside.

### Example:
```meow
meow a = 10
meowIf a > 5
    purr "Big Cat"
otherwiseScratch
    purr "Small Cat"
```

### Boolean Expressions Supported:
| Operator | Meaning |
|----------|---------|
| `>`  | Greater than |
| `>=` | Greater than or equal |
| `<`  | Less than |
| `<=` | Less than or equal |
| `==` | Equal to |
| `!=` | Not equal to |

---
## 🔁 Loops
🚧 **(Work in Progress)** 🚧

---
### Output of sample code 
Code: 
```
Hello Kitty!
meow a = 4
meow b = 5
// if a is not equal to b then add them else subtract them 
meowIf a != b 
    meow a = a + b
otherwiseScratch
    meow a = a - b
purrln a
Bye Kitty!

This will be ignored
```

Output: 
``` 9 ```

![image](https://github.com/user-attachments/assets/c2be3039-22e7-4e35-9b28-01c56883b449)

---
🐱 Happy Coding with Meow Lang! 🐾
