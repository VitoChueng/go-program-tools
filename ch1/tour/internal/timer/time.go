package timer

import "time"

func GetNowTime() time.Time{
	loc,_ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

func GetCalTime(currTimer time.Time,d string)(time.Time,error){
	dur,err := time.ParseDuration(d)
	if err != nil{
		return time.Time{},err
	}
	return currTimer.Add(dur),nil
}