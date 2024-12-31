#!/bin/bash
#  --include_imports
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/common.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/channel/channel.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/amendment/amendment.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/hotel/hotel.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/inout_authorization/inout_authorization.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/order/order.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/park/park.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/portrait/portrait.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/product/product.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/report/report.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/section/section.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/temporary_booking/temporary_booking.proto
protoc -I .  --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative cmd/odas/grpc/tourist/tourist.proto

# grpcurl -H 'appid:abcdefg'  -d '{"start": "2023-06-01T00:00:00Z","end":"2023-08-23T00:00:00Z"}' -plaintext 127.0.0.1:50051 proto.ChannelService.OrderChannelBySid
