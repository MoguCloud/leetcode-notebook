// Your runtime beats 60 % of golang submissions (146 ms)
// Your memory usage beats 55 % of golang submissions (15 MB)
//
// 使用2个hash table
// hash table 1 的 key是用户id value是进站站名和进入时间；
// hash table 2 的 key 是行程名称 value是总时间和总人数

type CheckInLog struct {
	station string
	t       int
}

type TimeLog struct {
	total int
	count int
}

type UndergroundSystem struct {
	CheckInLogMap map[int]CheckInLog
	TimeMap       map[string]TimeLog
}

func getName(in string, out string) string {
	return fmt.Sprintf("%s>%s", in, out)
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		CheckInLogMap: make(map[int]CheckInLog),
		TimeMap:       make(map[string]TimeLog),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.CheckInLogMap[id] = CheckInLog{
		station: stationName,
		t:       t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkInLog := this.CheckInLogMap[id]
	name := getName(checkInLog.station, stationName)
	timeLog, ok := this.TimeMap[name]
	if !ok {
		timeLog = TimeLog{
			total: 0,
			count: 0,
		}
	}
	timeLog.total += t - checkInLog.t
	timeLog.count++
	this.TimeMap[name] = timeLog
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	name := getName(startStation, endStation)
	timeLog := this.TimeMap[name]
	return float64(timeLog.total) / float64(timeLog.count)
}
