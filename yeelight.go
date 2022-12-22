package yeelight

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Config struct {
	conn      net.Conn
	IpAddress string
	Port      int
}

type Request struct {
	ID     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type Response struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result"`
}

type FlowExpression struct {
	Duration   int
	Mode       int
	Value      int
	Brightness int
}

type Scene struct {
	Action           string
	Color            int
	ColorTemperature int
	Hue              int
	Saturation       int
	ColorFlow        []FlowExpression
	Mode             int
	Brightness       int
	Duration         int
}

func generateID() int {
	rand.Seed(time.Now().UTC().UnixNano())
	id := math.Floor((rand.Float64() * 100000) + 1)

	return int(id)
}

func New(c *Config) Config {
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%v", c.IpAddress, c.Port))

	return Config{
		conn,
		c.IpAddress,
		c.Port,
	}
}

/*
After you run the tcp you should close the tcp connection.
*/
func (c *Config) Close() {
	c.conn.Close()
}

/*
This function is used to generate RGB to decimal integer to represent the color.
You should fill red, green and blue with integer.
*/
func (c *Config) GenerateRGB(red, green, blue int) int {
	color := (red * 65536) + (green * 256) + blue

	return color
}

/*
This function is used to retrieve current property. The allowed values is "power", "bright", "ct", "rgb", "hue",
"sat", "color_mode", "flowing", "delayoff", "flow_params", "music_on", "name", "bg_power", "bg_flowing", "bg_flow_params",
"bg_ct", "bg_lmode", "bg_bright", "bg_rgb", "bg_hue", "bg_sat", "nl_br", "active_mode"
*/
func (c *Config) GetProps(p ...interface{}) (Response, error) {
	id := generateID()

	request := Request{
		ID:     int(id),
		Method: "get_prop",
		Params: p,
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to change the color temperature. The allowed value for temp is in range 1700 ~ 6500.
The allowed value effect is "sudden" and "smooth". For the duration action, it's should be more than 30 milliseconds.
*/
func (c *Config) SetColorTemp(temp int, effect string, duration int) (Response, error) {
	id := generateID()

	if effect != "smooth" && effect != "sudden" {
		return Response{}, fmt.Errorf("effect values is wrong, yeelight only supports effects 'smooth' and 'sudden'")
	}

	if temp < 1700 {
		temp = 1700
	}

	if temp > 6500 {
		temp = 6500
	}

	if duration < 30 {
		duration = 30
	}

	request := Request{
		ID:     int(id),
		Method: "set_ct_abx",
		Params: []interface{}{temp, effect, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to change the color. The allowed value effect is "sudden" and "smooth".
For the duration action, it's should be more than 30 milliseconds.
*/
func (c *Config) SetRGB(red int, green int, blue int, effect string, duration int) (Response, error) {
	id := generateID()

	if red < 0 || red > 255 {
		return Response{}, fmt.Errorf("rgb should be in range 0-255")
	}

	if green < 0 || green > 255 {
		return Response{}, fmt.Errorf("rgb should be in range 0-255")
	}

	if blue < 0 || blue > 255 {
		return Response{}, fmt.Errorf("rgb should be in range 0-255")
	}

	if effect != "smooth" && effect != "sudden" {
		return Response{}, fmt.Errorf("effect values is wrong, yeelight only supports effects 'smooth' and 'sudden'")
	}

	if duration < 30 {
		duration = 30
	}

	color := (red * 65536) + (green * 256) + blue

	request := Request{
		ID:     int(id),
		Method: "set_rgb",
		Params: []interface{}{color, effect, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to change the color with hue and saturation.
The allowed value hue is in range 0 ~ 359.
The allowed value sat is in range 0 ~ 100.
The allowed value effect is "sudden" and "smooth".
For the duration action, it's should be more than 30 milliseconds.
*/
func (c *Config) SetHueSaturation(hue int, sat int, effect string, duration int) (Response, error) {
	id := generateID()

	if hue < 0 || hue > 359 {
		return Response{}, fmt.Errorf("hue value should be in range 0-359")
	}

	if sat < 0 || sat > 100 {
		return Response{}, fmt.Errorf("saturation value should be in range 0-100")
	}

	if effect != "smooth" && effect != "sudden" {
		return Response{}, fmt.Errorf("effect values is wrong, yeelight only supports effects 'smooth' and 'sudden'")
	}

	if duration < 30 {
		duration = 30
	}

	request := Request{
		ID:     int(id),
		Method: "set_hsv",
		Params: []interface{}{hue, sat, effect, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to change the brightness.
The allowed value brightness is in range 1 ~ 100.
The allowed value effect is "sudden" and "smooth".
For the duration action, it's should be more than 30 milliseconds.
*/
func (c *Config) SetBright(brightness int, effect string, duration int) (Response, error) {
	id := generateID()

	if effect != "smooth" && effect != "sudden" {
		return Response{}, fmt.Errorf("effect values is wrong, yeelight only supports effects 'smooth' and 'sudden'")
	}

	if brightness < 1 {
		brightness = 1
	}

	if brightness > 100 {
		brightness = 100
	}

	if duration < 30 {
		duration = 30
	}

	request := Request{
		ID:     int(id),
		Method: "set_bright",
		Params: []interface{}{brightness, effect, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to switch on or off. The allowed value effect is "sudden" and "smooth".
For the duration action, it's should be more than 30 milliseconds.
*/
func (c *Config) SetPower(power bool, effect string, duration int) (Response, error) {
	id := generateID()

	p := "off"
	if power {
		p = "on"
	}

	if duration < 30 {
		duration = 30
	}

	if effect != "smooth" && effect != "sudden" {
		return Response{}, fmt.Errorf("effect values is wrong, yeelight only supports effects 'smooth' and 'sudden'")
	}

	request := Request{
		ID:     int(id),
		Method: "set_power",
		Params: []interface{}{p, effect, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to save current state.
*/
func (c *Config) SetDefault() (Response, error) {
	id := generateID()

	request := Request{
		ID:     int(id),
		Method: "set_default",
		Params: []interface{}{},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to start a color flow. "count" is the number of visible state changing before color flow stopped.
"action" is the action taken after the flow is stopped.
0 means smart LED recover to the state before the color flow started.
1 means smart LED stay at the state when the flow is stopped.
2 means turn off the smart LED after the flow is stopped.

"exprs" is the expression of the state changing series. Fill with "mode" 1 - color, 2 - color temperature,
"duration" for the duration in milliseconds, "value" is following the "mode", color or color temperature.
*/
func (c *Config) SetColorFlow(count, action int, exprs []FlowExpression) (Response, error) {
	id := generateID()

	if action < 0 && action > 2 {
		return Response{}, fmt.Errorf("action should be in range 0-2")
	}

	var exprStrArr []string
	for _, expr := range exprs {
		if expr.Duration < 30 {
			expr.Duration = 30
		}

		if expr.Mode != 1 && expr.Mode != 2 && expr.Mode != 7 {
			return Response{}, fmt.Errorf("flow expression mode should be 1, 2 or 7. 1 - color, 2 - color temperature, 7 - sleep")
		}

		if expr.Brightness < 1 {
			expr.Brightness = 1
		} else if expr.Brightness > 100 {
			expr.Brightness = 100
		}

		exprStrArr = append(exprStrArr, fmt.Sprintf("%d,%d,%d,%d", expr.Duration, expr.Mode, expr.Value, expr.Brightness))
	}

	request := Request{
		ID:     int(id),
		Method: "start_cf",
		Params: []interface{}{count, action, strings.Join(exprStrArr, ",")},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
The function is used to stop current color flow.
*/
func (c *Config) StopColorFlow() (Response, error) {
	id := generateID()

	request := Request{
		ID:     int(id),
		Method: "stop_cf",
		Params: []interface{}{},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to set scene with color, hue saturation, color temperature or color flow.
*/
func (c *Config) SetScene(scene Scene) (Response, error) {
	id := generateID()

	var params []interface{}

	if scene.Action == "color" {
		params = []interface{}{scene.Action, scene.Color, scene.Brightness}
	} else if scene.Action == "hsv" {
		params = []interface{}{scene.Action, scene.Hue, scene.Saturation, scene.Brightness}
	} else if scene.Action == "ct" {
		params = []interface{}{scene.Action, scene.ColorTemperature, scene.Brightness}
	} else if scene.Action == "cf" {
		params = []interface{}{scene.Action, scene.Duration, scene.Mode, scene.ColorFlow}
	}

	request := Request{
		ID:     id,
		Method: "set_scene",
		Params: params,
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to added a cron job to turn off the lamp.
*/
func (c *Config) CronAdd(timer int) (Response, error) {
	id := generateID()

	request := Request{
		ID:     id,
		Method: "cron_add",
		Params: []interface{}{0, timer},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to get cron jobs in queue.
*/
func (c *Config) CronGet() (Response, error) {
	id := generateID()

	request := Request{
		ID:     id,
		Method: "cron_get",
		Params: []interface{}{0},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to delete cron job in queue.
*/
func (c *Config) CronDelete() (Response, error) {
	id := generateID()

	request := Request{
		ID:     id,
		Method: "cron_del",
		Params: []interface{}{0},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to adjust brightness, color tempterature or color.
The allowed value "action" is increase, decrease and circle.
The allowed value "prop" is bright, ct and color.
*/
func (c *Config) SetAdjust(action, prop string) (Response, error) {
	id := generateID()

	if action != "increase" && action != "decrease" && action != "circle" {
		return Response{}, fmt.Errorf("action should be increase, decrease or circle")
	}

	if prop != "bright" && prop != "ct" && prop != "color" {
		return Response{}, fmt.Errorf("prop should be bright, ct or color")
	}

	if prop == "color" && action != "circle" {
		return Response{}, fmt.Errorf("when props is color, the action can only be circle")
	}

	request := Request{
		ID:     id,
		Method: "set_adjust",
		Params: []interface{}{action, prop},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
The function is used to change the device name, stored in device not cloud.
*/
func (c *Config) SetName(name string) (Response, error) {
	id := generateID()

	request := Request{
		ID:     id,
		Method: "set_name",
		Params: []interface{}{name},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to adjust the brightness by specified bright percentage within specified duration.
"bright" should fill with range -100 ~ 100.
"duration" set the action duration with milisecond.
*/
func (c *Config) AdjustBright(bright, duration int) (Response, error) {
	id := generateID()

	request := Request{
		ID:     id,
		Method: "adjust_bright",
		Params: []interface{}{bright, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}

/*
This function is used to adjust the color temperature by specified bright percentage within specified duration.
"bright" should fill with range -100 ~ 100.
"duration" set the action duration with milisecond.
*/
func (c *Config) AdjustColorTemperature(bright, duration int) (Response, error) {
	id := generateID()

	request := Request{
		ID:     id,
		Method: "adjust_ct",
		Params: []interface{}{bright, duration},
	}

	bytes, _ := json.Marshal(request)

	_, err := c.conn.Write([]byte(string(bytes) + "\r\n"))
	if err != nil {
		return Response{}, err
	}

	message, _ := bufio.NewReader(c.conn).ReadString('\n')

	var r Response
	err = json.Unmarshal([]byte(message), &r)
	if err != nil {
		return Response{}, err
	}

	return r, nil
}
