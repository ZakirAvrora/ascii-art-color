# Ascii-art-color
## Authors
@ZakirAvrora
@sudo_root

## Features
The project works in a similar way that the original ascii-art,  but with implementation of __flag__ `--color=<color>`:
- Color representation has to be in format of (`ANSI, RGB, HSL or hex`).
- Ability to select coloring between single letters, words and set of words.
- If the letter is not specified, the whole text should be colored.
- The flag must have exactly the same format as above, any other formats must return the following usage message:
```bash
Usage: go run . [STRING] [OPTION]

EX: go run . something --color=<color>
```


## Notes
1. Template that was used for ascii-art is default standard.txt
2. The RGB format should be given as `--color=rgb($R, $G, $B)` where __`$R`__ , __`$G`__  and __`$B`__ are intesity of red, green and blue colors accodingly with a integer value between 0 and 255. Example:
```bash
go run . "Hello Wolrd" "--color=rgb(255, 0, 0)"
```
3. The HSL format should be given as `--color=rgb($H, $S%, $L%)` where __`$H`__ , __`$S`__  and __`$L`__ are hue, saturation and lightness respectively. The integer value of hue  is between 0 and 360, whereas the saturation and lightness presented in form of percentage from 0 to 100. Example:
```bash
go run . "Hello Wolrd" "-color=hsl(0, 100%, 50%)"
```
4. The Hex cas has to be given as `--color=#$HEX`, where __`$HEX`__ is six digit hexadecimal representation of color.Example:
```bash
go run . "Hello Wolrd" "-color=#ff0000"
```

## Coloring a single letter or a set of letters
Specification of character coloring can be conducted by two ways:
1. __Coloring using the index__: In this case user can specify indexes in the given string, the given index range `[start:end]` has the same working priciple as indexes for array. However, it should be noted, the negative or high index than length of string always take values of starting index or end of string. Example:
```bash
go run . "Hello World" "--color=hsl(0, 100 %,50%)" "[1:5]"
```
2. __Coloring using the sets__: The sets of words should be provided into brackets in the form of `{...}`. The set of words or chars have to be seperated by the comma.In the case of absence of any word, nothing will be colored. Example:
```bash
go run . "Hello World" "--color=yellow" "{Hello}"
```