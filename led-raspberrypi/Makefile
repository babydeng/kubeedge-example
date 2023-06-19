
# make led_light_mapper
.PHONY: default led_light_mapper push
led_light_mapper:
	env GOOS=linux GOARCH=arm64 go build main.go
	docker build -t mixiaochao/led-light-mapper:v1.1 .

push:
	docker push mixiaochao/led-light-mapper:v1.1

remove:
	rsync  main pi@10.177.21.124:/home/pi

