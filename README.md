            func InitMongodb() (err error) {
            // 1.连接mongodb
            client, err = mongo.Connect(context.Background(),
            options.Client().ApplyURI("mongodb://mongodb:").
            SetConnectTimeout(5*time.Second))
            if err != nil {
            fmt.Println(err)
            return err
            }
            
                if err = client.Ping(context.Background(), nil); err != nil {
                    fmt.Println(err)
                    return err
                }
            
                // 由于都是对同一个集合进行操作，所以在初始化mongodb时就选择了集合，防止后续出现大量重复代码
                coll = client.Database("").Collection("stu")
                return
            }