package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/apcera/termtables"
	"github.com/araddon/dateparse"
)

var examples = []string{
	"May 8, 2009 5:57:51 PM",
	"Mon Jan  2 15:04:05 2006",
	"Mon Jan  2 15:04:05 MST 2006",
	"Mon Jan 02 15:04:05 -0700 2006",
	"Monday, 02-Jan-06 15:04:05 MST",
	"Mon, 02 Jan 2006 15:04:05 MST",
	"Tue, 11 Jul 2017 16:28:13 +0200 (CEST)",
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"Mon Aug 10 15:44:11 UTC+0100 2015",
	"Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)",
	"12 Feb 2006, 19:17",
	"2013-Feb-03",
	//   mm/dd/yy
	"3/31/2014",
	"03/31/2014",
	"08/21/71",
	"8/1/71",
	"4/8/2014 22:05",
	"04/08/2014 22:05",
	"04/2/2014 03:00:51",
	"8/8/1965 12:00:00 AM",
	"8/8/1965 01:00:01 PM",
	"8/8/1965 01:00 PM",
	"8/8/1965 1:00 PM",
	"8/8/1965 12:00 AM",
	"4/02/2014 03:00:51",
	"03/19/2012 10:11:59",
	"03/19/2012 10:11:59.3186369",
	// yyyy/mm/dd
	"2014/3/31",
	"2014/03/31",
	"2014/4/8 22:05",
	"2014/04/08 22:05",
	"2014/04/2 03:00:51",
	"2014/4/02 03:00:51",
	"2012/03/19 10:11:59",
	"2012/03/19 10:11:59.3186369",
	//   yyyy-mm-ddThh
	"2006-01-02T15:04:05+0000",
	"2009-08-12T22:15:09-07:00",
	"2009-08-12T22:15:09",
	"2009-08-12T22:15:09Z",
	//   yyyy-mm-dd hh:mm:ss
	"2014-04-26 17:24:37.3186369",
	"2012-08-03 18:31:59.257000000",
	"2014-04-26 17:24:37.123",
	"2013-04-01 22:43:22",
	"2014-12-16 06:20:00 UTC",
	"2014-12-16 06:20:00 GMT",
	"2014-04-26 05:24:37 PM",
	"2014-04-26 13:13:43 +0800",
	"2014-04-26 13:13:44 +09:00",
	"2012-08-03 18:31:59.257000000 +0000 UTC",
	"2015-09-30 18:48:56.35272715 +0000 UTC",
	"2015-02-18 00:12:00 +0000 GMT",
	"2015-02-18 00:12:00 +0000 UTC",
	"2017-07-19 03:21:51+00:00",
	"2014-04-26",
	"2014-04",
	"2014",
	"2014-05-11 08:20:13,787",
	//  yyyymmdd and similar
	"20140601",
	// unix seconds, ms
	"1332151919",
	"1384216367189",
}

var (
	timezone = ""
)

func main() {
	flag.StringVar(&timezone, "timezone", "UTC", "Timezone aka `America/Los_Angeles` formatted time-zone")
	flag.Parse()

	if timezone != "" {
		// NOTE:  This is very, very important to understand
		// time-parsing in go
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			panic(err.Error())
		}
		time.Local = loc
	}

	table := termtables.CreateTable()

	table.AddHeaders("Input", "Parsed, and Output as %v")
	for _, dateExample := range examples {
		t, err := dateparse.ParseLocal(dateExample)
		if err != nil {
			panic(err.Error())
		}
		table.AddRow(dateExample, fmt.Sprintf("%v", t))
	}
	fmt.Println(table.Render())
}

/*
+-------------------------------------------------------+----------------------------------------+
| Input                                                 | Parsed, and Output as %v               |
+-------------------------------------------------------+----------------------------------------+
| May 8, 2009 5:57:51 PM                                | 2009-05-08 17:57:51 +0000 UTC          |
| Mon Jan  2 15:04:05 2006                              | 2006-01-02 15:04:05 +0000 UTC          |
| Mon Jan  2 15:04:05 MST 2006                          | 2006-01-02 15:04:05 +0000 MST          |
| Mon Jan 02 15:04:05 -0700 2006                        | 2006-01-02 15:04:05 -0700 -0700        |
| Monday, 02-Jan-06 15:04:05 MST                        | 2006-01-02 15:04:05 +0000 MST          |
| Mon, 02 Jan 2006 15:04:05 MST                         | 2006-01-02 15:04:05 +0000 MST          |
| Tue, 11 Jul 2017 16:28:13 +0200 (CEST)                | 2017-07-11 16:28:13 +0200 +0200        |
| Mon, 02 Jan 2006 15:04:05 -0700                       | 2006-01-02 15:04:05 -0700 -0700        |
| Mon Aug 10 15:44:11 UTC+0100 2015                     | 2015-08-10 15:44:11 +0000 UTC          |
| Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time) | 2015-07-03 18:04:07 +0100 GMT          |
| 12 Feb 2006, 19:17                                    | 2006-02-12 19:17:00 +0000 UTC          |
| 2013-Feb-03                                           | 2013-02-03 00:00:00 +0000 UTC          |
| 3/31/2014                                             | 2014-03-31 00:00:00 +0000 UTC          |
| 03/31/2014                                            | 2014-03-31 00:00:00 +0000 UTC          |
| 08/21/71                                              | 1971-08-21 00:00:00 +0000 UTC          |
| 8/1/71                                                | 1971-08-01 00:00:00 +0000 UTC          |
| 4/8/2014 22:05                                        | 2014-04-08 22:05:00 +0000 UTC          |
| 04/08/2014 22:05                                      | 2014-04-08 22:05:00 +0000 UTC          |
| 04/2/2014 03:00:51                                    | 2014-04-02 03:00:51 +0000 UTC          |
| 8/8/1965 12:00:00 AM                                  | 1965-08-08 00:00:00 +0000 UTC          |
| 8/8/1965 01:00:01 PM                                  | 1965-08-08 13:00:01 +0000 UTC          |
| 8/8/1965 01:00 PM                                     | 1965-08-08 13:00:00 +0000 UTC          |
| 8/8/1965 1:00 PM                                      | 1965-08-08 13:00:00 +0000 UTC          |
| 8/8/1965 12:00 AM                                     | 1965-08-08 00:00:00 +0000 UTC          |
| 4/02/2014 03:00:51                                    | 2014-04-02 03:00:51 +0000 UTC          |
| 03/19/2012 10:11:59                                   | 2012-03-19 10:11:59 +0000 UTC          |
| 03/19/2012 10:11:59.3186369                           | 2012-03-19 10:11:59.3186369 +0000 UTC  |
| 2014/3/31                                             | 2014-03-31 00:00:00 +0000 UTC          |
| 2014/03/31                                            | 2014-03-31 00:00:00 +0000 UTC          |
| 2014/4/8 22:05                                        | 2014-04-08 22:05:00 +0000 UTC          |
| 2014/04/08 22:05                                      | 2014-04-08 22:05:00 +0000 UTC          |
| 2014/04/2 03:00:51                                    | 2014-04-02 03:00:51 +0000 UTC          |
| 2014/4/02 03:00:51                                    | 2014-04-02 03:00:51 +0000 UTC          |
| 2012/03/19 10:11:59                                   | 2012-03-19 10:11:59 +0000 UTC          |
| 2012/03/19 10:11:59.3186369                           | 2012-03-19 10:11:59.3186369 +0000 UTC  |
| 2006-01-02T15:04:05+0000                              | 2006-01-02 15:04:05 +0000 UTC          |
| 2009-08-12T22:15:09-07:00                             | 2009-08-12 22:15:09 -0700 -0700        |
| 2009-08-12T22:15:09                                   | 2009-08-12 22:15:09 +0000 UTC          |
| 2009-08-12T22:15:09Z                                  | 2009-08-12 22:15:09 +0000 UTC          |
| 2014-04-26 17:24:37.3186369                           | 2014-04-26 17:24:37.3186369 +0000 UTC  |
| 2012-08-03 18:31:59.257000000                         | 2012-08-03 18:31:59.257 +0000 UTC      |
| 2014-04-26 17:24:37.123                               | 2014-04-26 17:24:37.123 +0000 UTC      |
| 2013-04-01 22:43:22                                   | 2013-04-01 22:43:22 +0000 UTC          |
| 2014-12-16 06:20:00 UTC                               | 2014-12-16 06:20:00 +0000 UTC          |
| 2014-12-16 06:20:00 GMT                               | 2014-12-16 06:20:00 +0000 UTC          |
| 2014-04-26 05:24:37 PM                                | 2014-04-26 17:24:37 +0000 UTC          |
| 2014-04-26 13:13:43 +0800                             | 2014-04-26 13:13:43 +0800 +0800        |
| 2014-04-26 13:13:44 +09:00                            | 2014-04-26 13:13:44 +0900 +0900        |
| 2012-08-03 18:31:59.257000000 +0000 UTC               | 2012-08-03 18:31:59.257 +0000 UTC      |
| 2015-09-30 18:48:56.35272715 +0000 UTC                | 2015-09-30 18:48:56.35272715 +0000 UTC |
| 2015-02-18 00:12:00 +0000 GMT                         | 2015-02-18 00:12:00 +0000 UTC          |
| 2015-02-18 00:12:00 +0000 UTC                         | 2015-02-18 00:12:00 +0000 UTC          |
| 2017-07-19 03:21:51+00:00                             | 2017-07-19 03:21:51 +0000 UTC          |
| 2014-04-26                                            | 2014-04-26 00:00:00 +0000 UTC          |
| 2014-04                                               | 2014-04-01 00:00:00 +0000 UTC          |
| 2014                                                  | 2014-01-01 00:00:00 +0000 UTC          |
| 2014-05-11 08:20:13,787                               | 2014-05-11 08:20:13.787 +0000 UTC      |
| 20140601                                              | 2014-06-01 00:00:00 +0000 UTC          |
| 1332151919                                            | 2012-03-19 10:11:59 +0000 UTC          |
| 1384216367189                                         | 2013-11-12 00:32:47.189 +0000 UTC      |
+-------------------------------------------------------+----------------------------------------+
*/
