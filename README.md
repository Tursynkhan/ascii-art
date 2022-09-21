# ascii-art
### Objectives

Ascii-art is a program which consists in receiving a `string` as an argument and outputting the `string` in a graphic representation using ASCII. Time to write big.

What we mean by a graphic representation using ASCII, is to write the `string` received using ASCII characters, as you can see in the example below:

```console
@@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
@@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
@@@@@@""WW8888&&WW888888WW``@@@@$$
BB$$``&&&&WWWW8888&&&&8888&&``@@@@
$$``&&WW88&&88&&&&8888&&88WW88``$$
@@""&&&&&&&&88888888&&&&&&88&&``$$
``````^^``^^^^^^````""^^``^^``^^``
""WW^^@@@@^^``````^^BB@@^^``^^&&``
^^&&^^@@````^^``&&``@@````^^^^&&``
``WW&&^^""``^^WW&&&&""``^^^^&&88``
^^8888&&&&&&WW88&&88WW&&&&88&&WW``
@@``&&88888888WW&&WW88&&88WW88^^$$
@@""88&&&&&&&&888888&&``^^&&88``$$
@@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
@@@@@@^^888888&&88&&&&MM88^^BB$$$$
@@@@@@BB````&&&&&&&&88""``BB@@BB$$
$$@@$$$$$$$$``````````@@$$@@$$$$$$
```

- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Take a look at the ASCII manual.


### Usage

```console
Usage: go run . [STRING] [BANNER]
EX: go run . something standard

student$ go run . "" | cat -e
student$ go run . "\n" | cat -e
$
student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
student$ go run . "hello" | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . "HeLlO" | cat -e
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $

student$ go run . "Hello\nThere" | cat -e
 _    _           _    _           $
| |  | |         | |  | |          $
| |__| |   ___   | |  | |    ___   $
|  __  |  / _ \  | |  | |   / _ \  $
| |  | | |  __/  | |  | |  | (_) | $
|_|  |_|  \___|  |_|  |_|   \___/  $
                                   $
                                   $
 _______   _                              $
|__   __| | |                             $
   | |    | |__      ___    _ __     ___  $
   | |    |  _ \    / _ \  | '__|   / _ \ $
   | |    | | | |  |  __/  | |     |  __/ $
   |_|    |_| |_|   \___|  |_|      \___| $
                                          $
                                          $
student$ go run . "Hello\n\nThere" | cat -e
 _    _           _    _           $
| |  | |         | |  | |          $
| |__| |   ___   | |  | |    ___   $
|  __  |  / _ \  | |  | |   / _ \  $
| |  | | |  __/  | |  | |  | (_) | $
|_|  |_|  \___|  |_|  |_|   \___/  $
                                   $
                                   $
$
 _______   _                              $
|__   __| | |                             $
   | |    | |__      ___    _ __     ___  $
   | |    |  _ \    / _ \  | '__|   / _ \ $
   | |    | | | |  |  __/  | |     |  __/ $
   |_|    |_| |_|   \___|  |_|      \___| $
                                          $
                                          $
                                          
student$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
student$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$
student$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```
### Authors

+ [Tursynkhan](https://01.alem.school/git/Tursynkhan)