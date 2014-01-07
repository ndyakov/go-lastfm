#go-lastfm

* * *
<img
src="http://th02.deviantart.net/fs70/PRE/f/2013/134/2/1/i_m_not_ready_yet____pancham_by_avidmc-d65aa8b.jpg" style="margin:20px auto; width: 450px; display:block;" />
* * * 
[Lastfm's api](http://www.last.fm/api) wrapper in [golang](http://golang.org).

Documentation available on [godoc.org](http://godoc.org/github.com/ndyakov/go-lastfm).

[![BuildStatus](https://travis-ci.org/ndyakov/go-lastfm.png)](https://travis-ci.org/ndyakov/go-lastfm)
[![GoDoc](https://godoc.org/github.com/ndyakov/go-lastfm?status.png)](https://godoc.org/github.com/ndyakov/go-lastfm)

## Instalation

```
    go get github.com/ndyakov/go-lastfm
```

## Usage

###1. Import the package.
```
import "github.com/ndyakov/go-lastfm"
```
###2. Create new api client object.

```
    lfm := lastfm.New("api-key")
```

You can optain api key after registration [here](http://www.last.fm/api/account/create).

###3. Browse the documentation for supported methods.
Browse [here](http://godoc.org/github.com/ndyakov/go-lastfm) for available
endpoints or take a look at ``example_test.go`` for examples.


## TODO

* Implement all methods that don't need authentication.
* Make authentication work.
* Implement the rest of the API methods.

## License
   Copyright 2014 Nedyalko Dyakov

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

