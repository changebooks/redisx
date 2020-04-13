# redisx
redis + log
==

<pre>
profile, err := log.NewProfile(map[string]string{
    log.ProfileDirectory: "./log",
    log.ProfileChannel:   "redis",
})

if err != nil {
    fmt.Println(err)
    return
}

stream, err := log.NewStream(profile)
if err != nil {
    fmt.Println(err)
    return
}

logger, err := log.NewLogger(stream, "test", 1)
if err != nil {
    fmt.Println(err)
    return
}

idRegister := &log.IdRegister{}
idRegister.SetTraceId("trace-id-10001")
idRegister.SetBizId("biz-id-20002")

redis2, err := redisx.New(idRegister, logger, map[string]string{redis.ProfileHost: "127.0.0.1"})
if err != nil {
    fmt.Println(err)
    return
}

result, err := redis2.SetEx(idRegister, "id", 10001, 3600)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)

value, err := redis2.Get(idRegister, "id")
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(value)
</pre>
