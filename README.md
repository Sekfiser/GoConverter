К сожалению полностью конвертатор поднять в докере - невозможно. Этот основывается на использовании Word и из-за этого нужна Windows, потому что LibreOffice выдает некорректные результаты в некоторых случаях (возможно из-за защищенного просомтра). Поэтому в планах выделить publisher и consumer в отдельные репозитории и поднять в docker.

В Docker запускается rabbitMQ, после запускается publisher и converter. После этого можно посылать из postman запросы на [localhost:8090](http://127.0.0.1:8090/file/FromFile) с содержанием в body файла для конвертации с именем FileToConvert
