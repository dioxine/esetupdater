# esetupdater
Updater for ESET AV from custom mirros

edit `config.toml` accordingly comments in it

use `go run .` to see options, almost everything you need to start updating is to do `go run . update` 
or `go run . update -f` to force full downloading of updates base.

and of course you can build it to native executable for your system with `go build`

more carefull README soon as basic auth too. at the moment username and password do not work, in close TODO.
