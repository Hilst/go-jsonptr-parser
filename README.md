# go-custom-parser
Go project to study and parse JSON pointers

possibly with new functionalities
##
JSON pointer is a standart (RFC6901) way to point to a value in a JSON object
For more pratical approach:
This JSON
```json
{
    "a": "A",
    "b": "B",
    "c": [
        "c1",
        "c2",
    ],
    "d": [
        {
            "single": 1,
            "double": 2
        },
        {
            "single": 2,
            "double": 4
        }
    ]
}
```
Would be pointed that way:
```text
/a => "A"

/b => "B"

/c/0 => "c1"

/c/1 => "c2"

/d/0/double => 2

```
##
At start the idea would be add the functionalities of an <b>index</b>,
to point to a variety of values in an array
```text
/d/{{index}}/single => [1, 2]
```
and a <b>where</b> expression,
to point to a specif value in an array
```text
/d/{{where: /single == 2}}/double => 4
```
##
- [x] Tokenizer
- [x] Lexer
- [ ] Syntax Analysis
- [ ] Compilation