package timeDemo

import (
	"fmt"
	"time"
)

/**时间类*/

func Run() {
	p := fmt.Println

	now := time.Now()
	p(now)

	/**自定义一个日期*/
	then := time.Date(2017, 11, 3, 14, 13, 44, 7490547, time.Local)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	/**获取星期几*/
	p(then.Weekday())

	/**自定义时间是否前于现在*/
	p(then.Before(now))
	/**自定义时间是否后于现在*/
	p(then.After(now))
	/**自定义时间是否与现在相等*/
	p(then.Equal(now))
	/**计算两个时间之间的差异*/
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	/**获取当前毫秒*/
	p("---------------------")
	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	p(secs, millis, nanos)

	/**秒转日期*/
	p(time.Unix(secs, 0))
	/**纳米秒转日期*/
	p(time.Unix(0, nanos))

	/**格式化/解析时间日期*/
	p("---------------------")
	p(now.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2005-12-25T10:30:44+00:00")
	p(t1)

	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	t2, _ := time.Parse("3 04 PM", "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
