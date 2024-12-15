# qngplex

qngplex is an open source [Bitcoin](http://bitcoin.org/) block chain browser written in [Go](http://golang.org/), it allows you to search and navigate the [block chain](https://en.bitcoin.it/wiki/Block_chain). Also provides APIs to access data pragmatically.

**This is an early release, you might expect some bugs.**

## Requirements

- A [qng](https://github.com/Qitmeer/qng) instance
- Go >=1.2
- [Redis](http://redis.io/) 2.6+
- [SSDB](https://github.com/ideawu/ssdb)
- [LevelDB](https://code.google.com/p/leveldb/)
- 150+GB disk space / 4+GB RAM

Building qngplex database takes **1+ week** on a small server (dual core 1.2GHz/6GB RAM) and **28 hours** on a more powerful server (i5/16GB RAM).

## Installation

Assuming you have:

- a working Go workspace (and `$GOPATH` already set)
- [Redis 2.6+](http://redis.io/)
- [SSDB](https://github.com/ideawu/ssdb)
- [LevelDB](https://code.google.com/p/leveldb/) ([nice tutorial](http://techoverflow.net/blog/2012/12/14/compiling-installing-leveldb-on-linux/))
- [Snappy](http://code.google.com/p/snappy/)

You will also need to export `CGO_LDFLAGS`, needed to install [levigo](https://github.com/jmhodges/levigo).

    $ git clone https://github.com/tsileo/qngplex.git
    $ cd qngplex
    $ export CGO_LDFLAGS="-L/usr/local/lib -L/usr/local/lib -lsnappy"
    $ ./build.sh
    $ cp -r config.sample.json config.json
    $ vim config.json

Start the initial import (the example use nohup, but you should use a tool like [supervisord](http://supervisord.org/)):

    $ nohup ./bin/qngplex-import > import.log&

And once the process is done, you will have to restart you bitcoind with the ` -blocknotify``` parameter:  `-blocknotify="/home/thomas/qngplex/bin/qngplex-blocknotify -c /home/thomas/qngplex/config.json %s"`. Now you can start `qngplex-prod``:

    $ nohup ./bin/qngplex-prod > prod.log&

Even while importing, you can start the webserver:

    $ ./bin/qngplex-server

## Roadmap

- Stabilize everything for 1.0 release.
- Make sure a transaction can't be processed multiples times
- A receive payment API

Some features that are on my TODO list:

- An easy way to monitor Bitcoin address via API (maybe using Webhooks)
- Convert BTC to fiat money easily
- An official Python module to interact with the API and offer a reliable way to track address
- An official JS lib to interact with the API
- A Watch-only addresses page
- Display unconfirmed transactions on address page
- Live notification on a unconfirmed transaction page when it actually get included in a block
- Parse the coinbase to extract which pool mined the block
- An admin interface to monitor bitcoind/qngplex
- New SSE endoind: utxin/utxout for a given address
- Escrow transaction handling
- Docker build
- Provides supervisord config
- ... (don't hesitate to request features!)

## Documentation

The documentation is written in Markdown and is available in the docs directory, it's also available online (powered by [MkDocs](http://www.mkdocs.org/)) on [docs.qngplex.com](http://docs.qngplex.com).

## Contribution

Contribution are welcome, see [HACKING.md](HACKING.md) and [DESIGN.md](DESIGN.md) to get started.

## Feedback / Support

You can ping me @trucsdedev/contact@qngplex.com/thomas.sileo@gmail.com if you have any feedback/issue.

## Donation

BTC: 16obt7HXb3PmyDb1wZMA2X7HYPUPHp45GB

## License (MIT)

Copyright (c) 2014 Thomas Sileo

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
