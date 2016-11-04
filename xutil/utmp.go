package xutil

const (
	LineSize = 32
	NameSize = 32
	HostSize = 256
)

type exit_status struct {
	E_termination int16
	E_exit        int16
}

type ut_tv struct {
	Tv_sec  int32
	Tv_usec int32
}

type Utmp struct {
	Ut_type			int32
	Ut_pid 			int32
	Ut_line  		[LineSize]byte
	Ut_id    		[4]byte
	Ut_user    		[NameSize]byte
	Ut_host    		[HostSize]byte
	Exit_status    	exit_status
	Ut_session 		int32
	Time    		ut_tv
	Yt_addr_v6    	[4]int32
	Unused  		[20]byte
}

//key-value map , C의 hash table, string과 float 타입만 사용가능.
//map[key]value
//interface 함수집합, 비어있느 interface -> 모든타입 사용가능(Dynamic Type)
func UtmpMap(u *Utmp) map[string]interface{} {
	mapWtmp := map[string]interface{}{}
	mapWtmp["Type"] = u.Ut_type
	mapWtmp["Pid"] = u.Ut_pid
	mapWtmp["Device"] = Trim(u.Ut_line[:])
	mapWtmp["ID"] = Trim(u.Ut_id[:])
	mapWtmp["User"] = Trim(u.Ut_user[:])
	mapWtmp["Host"] = Trim(u.Ut_host[:])
	mapWtmp["Status"] = u.Exit_status
	mapWtmp["Session"] = u.Ut_session
	mapWtmp["Time"] = GetDate64(u.Time)
	mapWtmp["Addr"] = GetIpStr(u.Yt_addr_v6)
	mapWtmp["Unused"] = Trim(u.Unused[:])

	return mapWtmp
}