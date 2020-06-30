package main

//type/struct definations used in whole PROJECT

type Schedule struct {
	Start_time string `json:"start_time"`
	End_time string `json:"end_time"`

}

type Event struct {
	Userid string `json:"userid"`
	Name string `json:"name"`
	Eventid string `json:"eventid"`
	Description string `json:"description"`
	Status string `json:"status"`
	Schedule Schedule `json:"schedule"`
	
}

type Eventlist struct {
	Eventlist []Event  `json:"eventlist"`
}

type Deleteid struct{
	Id string `json:"eventid"`
}
