# gorm-logrus

Logrus logger for gorm v2
#### how use

```
go get github.com/xiaohao0416/gorm-logrus
```

```
import "github.com/xiaohao0416/gorm-logrus"
import "github.com/sirupsen/logrus"


  logger := logrus.New()

  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
    Logger: gorm_logrus.New(logger),
  })
```