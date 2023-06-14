docker build \
  --build-arg INFLUXDB_DATABASE=demo \
  --build-arg INFLUXDB_URL=https://us-east-1-1.aws.cloud2.influxdata.com \
  --build-arg INFLUXDB_TOKEN=WhUifTNPds_0Jt1RCoo-aGwKXi0nBbvoCsS4k_WSQ0U-MsifW9r8DEzZzLOcFHQAybdR8ur5lZgh3mSGwXe96A== \
  -t goflight .
