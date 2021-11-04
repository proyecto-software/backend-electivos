package models

type Maxus struct {
	License_plate                string  `json:"License_plate"`
	Make                         string  `json:"Model"`
	Id_event                     int64   `json:"id"`
	Label                        string  `json:"label"`
	Vehicle_ID                   int64   `json:"vehicle_id"`
	Ecu_neutral_state            float64 `json:"ecu_neutral_state"`
	Ecu_speed                    float64 `json:"ecu_speed"`
	Altitude                     float64 `json:"al"`
	Ecu_ignition_state           float64 `json:"ecu_ignition_state"`
	Ecu_stability_control        float64 `json:"ecu_stability_control"`
	Ecu_throttle                 float64 `json:"ecu_throttle"`
	Ecu_ac_state                 float64 `json:"ecu_ac_state"`
	Ecu_rpm                      float64 `json:"ecu_rpm"`
	Event_time                   string  `json:"event_time"`
	Longitude                    float64 `json:"lon"`
	Latitude                     float64 `json:"lat"`
	Ecu_ready_state              float64 `json:"ecu_ready_state"`
	Ecu_handbrake                float64 `json:"ecu_handbrake"`
	Ecu_reverse_state            float64 `json:"ecu_reverse_state"`
	Valid_position               bool    `json:"valid_position"`
	Event_hour                   float32 `json:"event_hour"`
	System_time                  string  `json:"system_time"`
	Ecu_state_of_charge          float64 `json:"ecu_state_of_charge"`
	Ecu_distance                 float64 `json:"ecu_distance"`
	Age                          int     `json:"Age"`
	Ecu_drive_state              float64 `json:"ecu_drive_state"`
	Ecu_regenerative_brake_level float64 `json:"ecu_regenerative_brake_level"`
	Io_ignition                  bool    `json:"io_ign"`
}
type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Home struct {
	Domain     string `json:"Domain"`
	Name       string `json:"Name"`
	Version    string `json:"Version"`
	Endpoint_1 string `json:"Endpoint_1"`
	Endpoint_2 string `json:"Endpoint_2"`
	Endpoint_3 string `json:"Endpoint_3"`
	Endpoint_4 string `json:"Endpoint_4"`
	Endpoint_5 string `json:"Endpoint_5"`
	Endpoint_6 string `json:"Endpoint_6"`
}

type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func HomeData(domain string, name string, ver string, end1 string, end2 string, end3 string, end4 string, end5 string, end6 string) Home {
	var d = &Home{
		domain,
		name,
		ver,
		end1,
		end2,
		end3,
		end4,
		end5,
		end6,
	}
	return *d
}

func MessageLoad(tipo string, msg string) Message {
	var m = &Message{tipo, msg}
	return *m
}
