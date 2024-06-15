# FlatBuffers VS ProtocolBuffer in Golang

This project is a benchmark comparing FlatBuffers and ProtocolBuffer in Golang for serializing and grpc, also including FlatBuffers examples.

I used FlatBuffers builder pool to reduce memory usage and make it faster.

## run bench
``` bash
go test -bench=. -benchtime=5s -benchmem
```

## bench results
```bash
goos: darwin
goarch: arm64
pkg: github.com/icefed/flatctest
BenchmarkSerial/0k/flatc/marshal-10     32386426               183.7 ns/op           344 B/op          3 allocs/op
BenchmarkSerial/0k/flatc/unmarshal-10           1000000000               0.09336 ns/op         0 B/op          0 allocs/op
BenchmarkSerial/0k/proto/marshal-10             36719302               163.0 ns/op           152 B/op          2 allocs/op
BenchmarkSerial/0k/proto/unmarshal-10           136989554               44.96 ns/op          144 B/op          4 allocs/op
BenchmarkSerial/4k/flatc/marshal-10             11720715               510.6 ns/op      8021.69 MB/s        4952 B/op          3 allocs/op
BenchmarkSerial/4k/flatc/unmarshal-10           1000000000               0.09285 ns/op  44113826.25 MB/s               0 B/op          0 allocs/op
BenchmarkSerial/4k/proto/marshal-10             11292236               531.7 ns/op      7704.30 MB/s        4992 B/op          2 allocs/op
BenchmarkSerial/4k/proto/unmarshal-10            8511976               711.7 ns/op      5755.01 MB/s        4240 B/op          5 allocs/op
BenchmarkSerial/64k/flatc/marshal-10             1713183              3497 ns/op        18740.08 MB/s      73816 B/op          3 allocs/op
BenchmarkSerial/64k/flatc/unmarshal-10          1000000000               0.09183 ns/op  713695376.05 MB/s              0 B/op          0 allocs/op
BenchmarkSerial/64k/proto/marshal-10             1724004              3538 ns/op        18523.45 MB/s      73856 B/op          2 allocs/op
BenchmarkSerial/64k/proto/unmarshal-10           1826797              3267 ns/op        20058.99 MB/s      65680 B/op          5 allocs/op
BenchmarkSerial/1m/flatc/marshal-10               209402             29682 ns/op        35326.96 MB/s    1056856 B/op          3 allocs/op
BenchmarkSerial/1m/flatc/unmarshal-10           1000000000               0.09219 ns/op  11374183234.74 MB/s            0 B/op          0 allocs/op
BenchmarkSerial/1m/proto/marshal-10               185460             32231 ns/op        32532.70 MB/s    1056896 B/op          2 allocs/op
BenchmarkSerial/1m/proto/unmarshal-10             195446             30848 ns/op        33991.33 MB/s    1048722 B/op          5 allocs/op
BenchmarkSerial/4m/flatc/marshal-10                84526             70602 ns/op        59407.96 MB/s    4202585 B/op          3 allocs/op
BenchmarkSerial/4m/flatc/unmarshal-10           1000000000               0.09704 ns/op  43221254577.56 MB/s            0 B/op          0 allocs/op
BenchmarkSerial/4m/proto/marshal-10                74522             82227 ns/op        51008.63 MB/s    4202625 B/op          2 allocs/op
BenchmarkSerial/4m/proto/unmarshal-10              74306             78487 ns/op        53439.47 MB/s    4194451 B/op          5 allocs/op
BenchmarkGRPC/0k/flatc/write-10                   356440             16762 ns/op            9799 B/op        179 allocs/op
BenchmarkGRPC/0k/flatc/read-10                    361366             16824 ns/op            9977 B/op        174 allocs/op
BenchmarkGRPC/0k/flatc/writestream-10             649605              9422 ns/op            1301 B/op         35 allocs/op
BenchmarkGRPC/0k/flatc/readstream-10              627028              9295 ns/op            1477 B/op         30 allocs/op
BenchmarkGRPC/0k/proto/write-10                   370357             16043 ns/op            9134 B/op        164 allocs/op
BenchmarkGRPC/0k/proto/read-10                    377991             16106 ns/op            9166 B/op        164 allocs/op
BenchmarkGRPC/0k/proto/writestream-10             684652              8664 ns/op             716 B/op         22 allocs/op
BenchmarkGRPC/0k/proto/readstream-10              684654              8763 ns/op             748 B/op         22 allocs/op
BenchmarkGRPC/4k/flatc/write-10                   326216             18163 ns/op         225.52 MB/s       19384 B/op        180 allocs/op
BenchmarkGRPC/4k/flatc/read-10                    314350             18859 ns/op         217.19 MB/s       24194 B/op        175 allocs/op
BenchmarkGRPC/4k/flatc/writestream-10             518239             11761 ns/op         348.28 MB/s       10862 B/op         37 allocs/op
BenchmarkGRPC/4k/flatc/readstream-10              491422             12158 ns/op         336.89 MB/s       15673 B/op         32 allocs/op
BenchmarkGRPC/4k/proto/write-10                   322810             18590 ns/op         220.34 MB/s       23118 B/op        170 allocs/op
BenchmarkGRPC/4k/proto/read-10                    330999             18381 ns/op         222.84 MB/s       23146 B/op        170 allocs/op
BenchmarkGRPC/4k/proto/writestream-10             506856             12021 ns/op         340.73 MB/s       14675 B/op         29 allocs/op
BenchmarkGRPC/4k/proto/readstream-10              498494             12206 ns/op         335.59 MB/s       14705 B/op         29 allocs/op
BenchmarkGRPC/64k/flatc/write-10                  150900             40364 ns/op        1623.63 MB/s      159731 B/op        183 allocs/op
BenchmarkGRPC/64k/flatc/read-10                   141702             43450 ns/op        1508.31 MB/s      234130 B/op        180 allocs/op
BenchmarkGRPC/64k/flatc/writestream-10            168550             35898 ns/op        1825.61 MB/s      151045 B/op         41 allocs/op
BenchmarkGRPC/64k/flatc/readstream-10             152575             39431 ns/op        1662.05 MB/s      225296 B/op         36 allocs/op
BenchmarkGRPC/64k/proto/write-10                  135118             45533 ns/op        1439.31 MB/s      225703 B/op        174 allocs/op
BenchmarkGRPC/64k/proto/read-10                   136824             43882 ns/op        1493.46 MB/s      225736 B/op        175 allocs/op
BenchmarkGRPC/64k/proto/writestream-10            152743             38901 ns/op        1684.68 MB/s      216906 B/op         33 allocs/op
BenchmarkGRPC/64k/proto/readstream-10             154670             40225 ns/op        1629.23 MB/s      216802 B/op         33 allocs/op
BenchmarkGRPC/1m/flatc/write-10                    15668            372921 ns/op        2811.79 MB/s     2137121 B/op        209 allocs/op
BenchmarkGRPC/1m/flatc/read-10                     16165            366897 ns/op        2857.95 MB/s     3195781 B/op        202 allocs/op
BenchmarkGRPC/1m/flatc/writestream-10              16540            360812 ns/op        2906.16 MB/s     2130089 B/op         65 allocs/op
BenchmarkGRPC/1m/flatc/readstream-10               16203            367039 ns/op        2856.85 MB/s     3194432 B/op         55 allocs/op
BenchmarkGRPC/1m/proto/write-10                    15099            378915 ns/op        2767.31 MB/s     3196282 B/op        190 allocs/op
BenchmarkGRPC/1m/proto/read-10                     14916            384893 ns/op        2724.33 MB/s     3192524 B/op        205 allocs/op
BenchmarkGRPC/1m/proto/writestream-10              17311            347154 ns/op        3020.49 MB/s     3197955 B/op         51 allocs/op
BenchmarkGRPC/1m/proto/readstream-10               16072            370094 ns/op        2833.27 MB/s     3187025 B/op         54 allocs/op
BenchmarkGRPC/4m/flatc/write-10                     3906           1306072 ns/op        3211.39 MB/s     8436711 B/op        243 allocs/op
BenchmarkGRPC/4m/flatc/read-10                      3913           1427557 ns/op        2938.10 MB/s    12635649 B/op        274 allocs/op
BenchmarkGRPC/4m/flatc/writestream-10               4426           1354595 ns/op        3096.35 MB/s     8425319 B/op         85 allocs/op
BenchmarkGRPC/4m/flatc/readstream-10                3776           1418248 ns/op        2957.38 MB/s    12635373 B/op        117 allocs/op
BenchmarkGRPC/4m/proto/write-10                     3993           1415338 ns/op        2963.47 MB/s    12635738 B/op        235 allocs/op
BenchmarkGRPC/4m/proto/read-10                      4076           1422215 ns/op        2949.14 MB/s    12635129 B/op        246 allocs/op
BenchmarkGRPC/4m/proto/writestream-10               4338           1270769 ns/op        3300.60 MB/s    12634890 B/op         78 allocs/op
BenchmarkGRPC/4m/proto/readstream-10                4341           1372446 ns/op        3056.08 MB/s    12631841 B/op         79 allocs/op
PASS
ok      github.com/icefed/flatctest     385.807s
```
