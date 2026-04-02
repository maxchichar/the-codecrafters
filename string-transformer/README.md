# String Transformer

## 🛰️ Sentinel String Transformer - Design Notes

### 1. Function Separation
Each transformation is implemented as its own function, not inside `main()`.  
This keeps the code modular, readable, and reusable.

#### Implemented Functions:
- `ToUpper(text string) string`
- `ToLower(text string) string`
- `ToCapital(text string) string`
- `Title(text string) string`
- `snakeCase(text string) string`
- `Reverse(text string) string`
- `ReversedText(input string) string`

The `main()` function is responsible only for:
- Reading user input
- Parsing commands
- Routing execution to the correct function

---

### 2. Capitalization (`cap` and `title`)
To modify only the first letter of each word:

#### Approach:
1. Convert the entire string to lowercase:
   ```go
   text = strings.ToLower(text)

## What the program does ?
My program accepts a command followed by
a string and applies the correct transformation.
```
Type 'help' to start

> help

Sentinel String Transformer Help
──────────────────────────────────────


        Usage: 
                <command> <text>
        Command:
                > upper     <text> 
                > snake     <text> 
                > lower     <text> 
                > cap       <text> 
                > reverse   <text> 
                > title         <text>
                > exit
        Example:
                > upper sentinel online
                > SENTINEL ONLINE

                > exit
                > Shutting down String Transformer...

        > 

```

```
 upper <text> 

 → Convert every letter to UPPERCASE

 Input : "sentinel is online"
 Output: "SENTINEL IS ONLINE"
```

```
 lower <text> 

 → Convert every letter to lowercase 

 Input : "ALERT LEVEL FIVE DETECTED" 
 Output: "alert level five detected"
```

```
 cap <text>  

 → Capitalise the first letter of every word.
 All other letters go lowercase. 

 Input : "director adaeze okonkwo"  
 Output: "Director Adaeze Okonkwo"

 Input : "THREAT LEVEL elevated"
 Output: "Threat Level Elevated"
```

```
 title <text>   

 → Title Case like cap, but small connector
 words stay lowercase unless they are the
 first word.

 Small words: a, an, the, and, but, or,
 for, nor, on, at, to, by, in,
 of, up, as, is, it 

 Input : "the fall of the western power grid"  
 Output: "The Fall of the Western Power Grid"

```

```
 snake <text> 
 → Convert to snake_case.
 All lowercase, spaces replaced with _.
 Remove any character that is not a letter,
 digit, or underscore.

 Input : "Alert! Level 5 detected."
 Output: "alert_level_5_detected"
```

```
 reverse <text>
 → Reverse each word individually.
 Word order stays the same. 
 Spaces between words are preserved.

 Input : "Lagos Nigeria"
 Output: "sogaL airegiN"
```

## How to run it with examples per interation with the commands.

when you type help in the CLI tool it shows you how to run it.

To start program:

```bash
go run main.go
```

```zsh
➜  string-transformer git:(main) ✗ go run main.go
SENTINEL STRING TRANSFORMER — ONLINE
──────────────────────────────────────

S.S.T Loading...

Menu Loading...

Sentinel String Transformer Menu
────────────────────────────────
Input 'help' to start

> help
Sentinel String Transformer Help
──────────────────────────────────────


        Usage: 
                <command> <text>
        Command:
                > upper     <text> 
                > snake     <text> 
                > lower     <text> 
                > cap       <text> 
                > reverse   <text> 
                > title         <text>
                > exit
        Example:
                > upper sentinel online
                > SENTINEL ONLINE

                > exit
                > Shutting down String Transformer...

        > upper how to run it
 HOW TO RUN IT
> snake how to rUn! IT           
how_to_run_it
> lower HOW TO RUN IT
 how to run it
> cap how TO run IT
 How To Run It
> reverse how to run it
 woh ot nur ti
> title how to run it
How to Run it
> exit
Shutting down String Transformer...


Good bye
```