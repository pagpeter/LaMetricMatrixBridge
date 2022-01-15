package main

type Success struct {
	Success bool `json:"success"`
}

// STATE

type DeviceState struct {
	Audio        Audio     `json:"audio"`
	Bluetooth    Bluetooth `json:"bluetooth"`
	Display      Display   `json:"display"`
	ID           string    `json:"id"`
	Mode         string    `json:"mode"`
	Model        string    `json:"model"`
	Name         string    `json:"name"`
	OsVersion    string    `json:"os_version"`
	SerialNumber string    `json:"serial_number"`
	Wifi         Wifi      `json:"wifi"`
}
type VolumeLimit struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
type VolumeRange struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
type Audio struct {
	Volume      int         `json:"volume"`
	VolumeLimit VolumeLimit `json:"volume_limit"`
	VolumeRange VolumeRange `json:"volume_range"`
}
type LowEnergy struct {
	Active      bool `json:"active"`
	Advertising bool `json:"advertising"`
	Connectable bool `json:"connectable"`
}
type Bluetooth struct {
	Active       bool      `json:"active"`
	Address      string    `json:"address"`
	Available    bool      `json:"available"`
	Discoverable bool      `json:"discoverable"`
	LowEnergy    LowEnergy `json:"low_energy"`
	Name         string    `json:"name"`
	Pairable     bool      `json:"pairable"`
}
type BrightnessLimit struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
type BrightnessRange struct {
	Max int `json:"max"`
	Min int `json:"min"`
}
type TimeBased struct {
	Enabled        bool   `json:"enabled"`
	EndTime        string `json:"end_time"`
	LocalEndTime   string `json:"local_end_time"`
	LocalStartTime string `json:"local_start_time"`
	StartTime      string `json:"start_time"`
}
type WhenDark struct {
	Enabled bool `json:"enabled"`
}
type Modes struct {
	TimeBased TimeBased `json:"time_based"`
	WhenDark  WhenDark  `json:"when_dark"`
}
type Screensaver struct {
	Enabled bool   `json:"enabled"`
	Modes   Modes  `json:"modes"`
	Widget  string `json:"widget"`
}
type Display struct {
	Brightness      int             `json:"brightness"`
	BrightnessLimit BrightnessLimit `json:"brightness_limit"`
	BrightnessMode  string          `json:"brightness_mode"`
	BrightnessRange BrightnessRange `json:"brightness_range"`
	Height          int             `json:"height"`
	Screensaver     Screensaver     `json:"screensaver"`
	Type            string          `json:"type"`
	Width           int             `json:"width"`
}
type Wifi struct {
	Active     bool   `json:"active"`
	Address    string `json:"address"`
	Available  bool   `json:"available"`
	Encryption string `json:"encryption"`
	Essid      string `json:"essid"`
	IP         string `json:"ip"`
	Mode       string `json:"mode"`
	Netmask    string `json:"netmask"`
	Strength   int    `json:"strength"`
}

// NOTIFICATIONS
// https://lametric-documentation.readthedocs.io/en/latest/reference-docs/device-notifications.html

type Notification struct {
	Priority       string `json:"priority"`
	IconType       string `json:"icon_type"`
	LifeTime       int    `json:"lifeTime"`
	Model          Model  `json:"model"`
	Created        string `json:"created"`
	ID             string `json:"id"`
	ExpirationDate string `json:"expiration_date"`
}
type GoalData struct {
	Start   int    `json:"start"`
	Current int    `json:"current"`
	End     int    `json:"end"`
	Unit    string `json:"unit"`
}
type Frames struct {
	Icon      string   `json:"icon,omitempty"`
	Text      string   `json:"text,omitempty"`
	GoalData  GoalData `json:"goalData,omitempty"`
	ChartData []int    `json:"chartData,omitempty"`
}
type Sound struct {
	Category string `json:"category"`
	ID       string `json:"id"`
	Repeat   int    `json:"repeat"`
}
type Model struct {
	Frames []Frames `json:"frames"`
	Sound  Sound    `json:"sound"`
	Cycles int      `json:"cycles"`
}

// APPS

type Widget struct {
	Index   int    `json:"index"`
	Package string `json:"package"`
}

type Trigger interface{}
type Action interface{}

type App struct {
	Actions     map[string]Action `json:"actions"`
	Package     string            `json:"package"`
	Title       string            `json:"title"`
	Triggers    []Trigger         `json:"triggers"`
	Vendor      string            `json:"vendor"`
	Version     string            `json:"version"`
	VersionCode string            `json:"versionCode"`
	Widgets     map[string]Widget `json:"widgets"`
}

type Apps map[string]App
