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
BenchmarkHelloSerial/0k/flatc/marshal-10        186544239               18.97 ns/op            0 B/op          0 allocs/op
BenchmarkHelloSerial/0k/flatc/unmarshal-10              1000000000               0.09372 ns/op         0 B/op          0 allocs/op
BenchmarkHelloSerial/0k/proto/marshal-10                56678445                66.80 ns/op          152 B/op          2 allocs/op
BenchmarkHelloSerial/0k/proto/unmarshal-10              55199026                65.86 ns/op          144 B/op          4 allocs/op
BenchmarkHelloSerial/4k/flatc/marshal-10                97855113                36.38 ns/op            0 B/op          0 allocs/op
BenchmarkHelloSerial/4k/flatc/unmarshal-10              1000000000               0.1003 ns/op          0 B/op          0 allocs/op
BenchmarkHelloSerial/4k/proto/marshal-10                 4541742               806.7 ns/op          4992 B/op          2 allocs/op
BenchmarkHelloSerial/4k/proto/unmarshal-10               4192128               848.0 ns/op          4240 B/op          5 allocs/op
BenchmarkHelloSerial/64k/flatc/marshal-10               11267330               452.3 ns/op             0 B/op          0 allocs/op
BenchmarkHelloSerial/64k/flatc/unmarshal-10             1000000000               0.1024 ns/op          0 B/op          0 allocs/op
BenchmarkHelloSerial/64k/proto/marshal-10                 325908             10368 ns/op           73856 B/op          2 allocs/op
BenchmarkHelloSerial/64k/proto/unmarshal-10               403837              9088 ns/op           65682 B/op          5 allocs/op
BenchmarkHelloSerial/1m/flatc/marshal-10                  514702              7344 ns/op               0 B/op          0 allocs/op
BenchmarkHelloSerial/1m/flatc/unmarshal-10              1000000000               0.09854 ns/op         0 B/op          0 allocs/op
BenchmarkHelloSerial/1m/proto/marshal-10                   83457             43622 ns/op         1056897 B/op          2 allocs/op
BenchmarkHelloSerial/1m/proto/unmarshal-10                135021             26625 ns/op         1048799 B/op          5 allocs/op
BenchmarkHelloSerial/4m/flatc/marshal-10                  106922             34388 ns/op               0 B/op          0 allocs/op
BenchmarkHelloSerial/4m/flatc/unmarshal-10              1000000000               0.09977 ns/op         0 B/op          0 allocs/op
BenchmarkHelloSerial/4m/proto/marshal-10                   43941             90742 ns/op         4202624 B/op          2 allocs/op
BenchmarkHelloSerial/4m/proto/unmarshal-10                 49084             72060 ns/op         4195306 B/op          5 allocs/op
BenchmarkHelloGRPC/0k/flatc/send_data-10                  216490             16336 ns/op            9732 B/op        170 allocs/op
BenchmarkHelloGRPC/0k/flatc/recv_data-10                  218889             16204 ns/op            9751 B/op        170 allocs/op
BenchmarkHelloGRPC/0k/flatc/sendrecv-10                   219512             16292 ns/op            9751 B/op        170 allocs/op
BenchmarkHelloGRPC/0k/proto/send_data-10                  208886             17308 ns/op            9569 B/op        174 allocs/op
BenchmarkHelloGRPC/0k/proto/recv_data-10                  210439             17328 ns/op            9569 B/op        174 allocs/op
BenchmarkHelloGRPC/0k/proto/sendrecv-10                   210094             17294 ns/op            9569 B/op        174 allocs/op
BenchmarkHelloGRPC/4k/flatc/send_data-10                  210930             17002 ns/op           19159 B/op        171 allocs/op
BenchmarkHelloGRPC/4k/flatc/recv_data-10                  210501             16903 ns/op           19177 B/op        171 allocs/op
BenchmarkHelloGRPC/4k/flatc/sendrecv-10                   202954             17907 ns/op           24005 B/op        172 allocs/op
BenchmarkHelloGRPC/4k/proto/send_data-10                  187312             20962 ns/op           23529 B/op        176 allocs/op
BenchmarkHelloGRPC/4k/proto/recv_data-10                  174826             19915 ns/op           23523 B/op        176 allocs/op
BenchmarkHelloGRPC/4k/proto/sendrecv-10                   161869             21689 ns/op           37635 B/op        178 allocs/op
BenchmarkHelloGRPC/64k/flatc/send_data-10                 107774             34295 ns/op          157624 B/op        173 allocs/op
BenchmarkHelloGRPC/64k/flatc/recv_data-10                 117202             32430 ns/op          157608 B/op        174 allocs/op
BenchmarkHelloGRPC/64k/flatc/sendrecv-10                   81314             44679 ns/op          232038 B/op        177 allocs/op
BenchmarkHelloGRPC/64k/proto/send_data-10                  77266             46374 ns/op          226539 B/op        180 allocs/op
BenchmarkHelloGRPC/64k/proto/recv_data-10                  81044             45650 ns/op          226542 B/op        181 allocs/op
BenchmarkHelloGRPC/64k/proto/sendrecv-10                   55699             66994 ns/op          446236 B/op        187 allocs/op
BenchmarkHelloGRPC/1m/flatc/send_data-10                    9674            365887 ns/op         2130517 B/op        202 allocs/op
BenchmarkHelloGRPC/1m/flatc/recv_data-10                    8600            354168 ns/op         2131897 B/op        194 allocs/op
BenchmarkHelloGRPC/1m/flatc/sendrecv-10                     5498            602067 ns/op         3207609 B/op        254 allocs/op
BenchmarkHelloGRPC/1m/proto/send_data-10                    8845            406737 ns/op         3190562 B/op        211 allocs/op
BenchmarkHelloGRPC/1m/proto/recv_data-10                    8656            392143 ns/op         3193544 B/op        201 allocs/op
BenchmarkHelloGRPC/1m/proto/sendrecv-10                     4878            631097 ns/op         6422014 B/op        263 allocs/op
BenchmarkHelloGRPC/4m/flatc/send_data-10                    2451           1366027 ns/op         8434029 B/op        256 allocs/op
BenchmarkHelloGRPC/4m/flatc/recv_data-10                    2676           1344151 ns/op         8434663 B/op        251 allocs/op
BenchmarkHelloGRPC/4m/flatc/sendrecv-10                     1477           2357601 ns/op        12685399 B/op        412 allocs/op
BenchmarkHelloGRPC/4m/proto/send_data-10                    2451           1498205 ns/op        12632983 B/op        280 allocs/op
BenchmarkHelloGRPC/4m/proto/recv_data-10                    2484           1453387 ns/op        12636255 B/op        254 allocs/op
BenchmarkHelloGRPC/4m/proto/sendrecv-10                     1450           2267929 ns/op        25381501 B/op        392 allocs/op
PASS
ok      github.com/icefed/flatctest     177.058s
```
