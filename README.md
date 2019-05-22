# Elasticsearchとfluentdを連携確認

ElastAlertのデモを試すためのリポジトリ

# 使い方
`$ docker-compose up` で立ち上がる

- 勝手に fluentd のimageがビルドされ、ルートディレクトリのfluent.confがコンテナに内蔵される
- fluentはhttpで待ち受けているので、`$ curl -i -X POST -d 'json={"action":"test","user":33333}' http://fluentd:8888/test.cycle` のようなリクエストでログを記録させればよい
- `docker-compose up` するとデフォルトのネットワークが構築されるので、curlでfluentdにリクエストするにはfront(see docker-compose.yml)内部で作業すると良い
