# gorm-logrus

Logrus logger for gorm v2
#### How to use

```
go get github.com/ttjio/gorm-logrus
```

```
import "github.com/ttjio/gorm-logrus"
import "github.com/sirupsen/logrus"


  logger := logrus.New()

  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
    Logger: gorm_logrus.New(logger),
  })
```
thinks [onrik/gorm-logrus](https://github.com/onrik/gorm-logrus) 
