package config

var DBMS = "mysql"
var User = "root"
var Password = "password"
var DbName = "goApi"

var QueryRateNew = "select currency, rate from goApi.Envelope where time = (select time from goApi.Envelope order by time DESC limit 1)"
var QueryRateByDay = "select currency, rate from goApi.Envelope where time like '%s'"
var QueryMaxMin = "select currency, max(rate) as max, min(rate) as min, avg(rate) as avg FROM goApi.Envelope group by currency"
