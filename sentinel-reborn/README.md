# THE GO ROUTINES SENTINEL REBORN 
    
    - What the program does
    - How to run it with examples
    - Each transformation listed and explained
    - What your personal contribution was
    - One thing you found hardest today
    - One thing you understand now that
      you did not understand this morning

## What the program does ?


```
The full program is a string and base transformation program. It receives the input, process it and then write it to the output
```

## How to run it with examples ?

you can run the code with the guide below:
```
Usage: go run . <input.txt> <output.txt>
```

Example of usage:
```bash
go run . input.txt output.txt
```

## Each transformation listed and explained

#### Transformations: 

* **(hex)**             :
>> `This converts hexadecimal to decimal`
* **(bin)**             :
>> `This converts binary to decimal`
* **(up)**              : 
>> `This converts lowercase words to uppercase`
* **(low)**             : 
>> `This converts capitalised or uppercase words to lowercase`
* **(up, n)**           : 
>> `This converts the 'n' words in a string before the transformer to uppercase`
* **(low, n)**          : 
>> `This converts the 'n' words in a string before the transformer to lowercase`
* **(cap)**             : 
>> `This converts either lower and uppercase to capitalised words e.g "my name is diRector (cap) x" >>> "my name is Director x"`
* **Punctuations  '**   : 
>> `Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").`
`Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".`
* **Articles**          : 
>> `Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").`
* **Quotes**            : 
>> `This fix a problem like this in a text if there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")`

## What your personal contribution was ?

* **Chibueze Maxwell**:
>> My contribution to this project helping with the planning session with my secatary blessing and also working on the helper function called FixQoute and error handling of same files in the main.go. Also helped in leading the team and in creating of branches and git pushing. also help in merging
* **Janai Egeonu**:
>> my contibution to the project was build the code executor(main.go) and compiling all the helper functions into a  file function called Complier and sending it to the main for proper execution of the program.
* **Agene Okoh**:
>> My own contribution to the project writing the helper function on Punctuation that handles fixing incorrect spacing before punctuation in a text and check if a line contains punctuation.
* **Blessing Anebi**:
>> my personal contribution was writing the helper function for "fix article", helped in the integration of all the functions and debugging.
* **Faith Ejembi**:
>> My own contribution to this project is writing the helperfunction for the transformer "(cap)" , this function scans the texts before the (cap) and capitalise the word before it. then removes the keyword and return the transformed text.
* **Ruth Agi**:
>> My personal contribution was writing the upper case function. ToUpper convert every lowercase letter in a string into uppercase. it goes through the text charater by charater and change. it does not change the original string(GO)
* **Ummulqulsal Musa**:
>> My contribution is i wrote a helper function that scans a sentence and coonverts any binary number that appears before (bin) into it decimal form
* **Edwin Ejembi**:
>> My contribution to this project is writing a helper function called hex to decimal. the first line split the string into the word (slice of string), then split by space , then i llop through every element in a slice of s. Then i check if the current word is equals to hex and also check that there is a word before it i > 0 to avoid the crash. Then convert the previous word from hex to decimal. then the text is transformed.
* **Emmanuel Unogwu**:
>> He helped with the FixArticle helper function using function main and helped in debugging and compilation.

## One thing you found hardest today

>>> merging our files to the repo.
>>> we also encountered error like merging conflicts and we refined the code in compilation because it was not compiling.

## One thing you understand now that you did not understand this morning
>>> pull request and merging