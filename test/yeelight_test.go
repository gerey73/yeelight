package test

import (
	"testing"
	"time"

	"github.com/LordAur/yeelight"
)

type Yeelight struct {
	IpAddress string
}

func TestGetProps(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.GetProps("bright", "power", "ct")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)
}

func TestSetRGB(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetRGB(58, 50, 153, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(5 * time.Second)

	r, err = y.SetRGB(255, 255, 255, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)
}

func TestSetTemp(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetColorTemp(5000, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(5 * time.Second)

	r, err = y.SetColorTemp(3200, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)
}

func TestSetHueSaturation(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	time.Sleep(3 * time.Second)

	r, err := y.SetHueSaturation(255, 45, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestSetBright(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetBright(10, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(5 * time.Second)

	r, err = y.SetBright(80, "smooth", 500)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestSetDefault(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetDefault()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestSetPower(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetPower(false, "smooth", 1000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(5 * time.Second)

	r, err = y.SetPower(true, "smooth", 1000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)
}

func TestStartColorFlow(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetColorFlow(0, 0, []yeelight.FlowExpression{
		{
			Duration:   1000,
			Mode:       2,
			Value:      2000,
			Brightness: 100,
		},
		{
			Duration:   1000,
			Mode:       2,
			Value:      1000,
			Brightness: 100,
		},
		{
			Duration:   1000,
			Mode:       2,
			Value:      4000,
			Brightness: 100,
		},
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestStartColorFlowRGB(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetColorFlow(0, 0, []yeelight.FlowExpression{
		{
			Duration:   2000,
			Mode:       1,
			Value:      y.GenerateRGB(135, 62, 35),
			Brightness: 100,
		},
		{
			Duration:   2000,
			Mode:       1,
			Value:      y.GenerateRGB(118, 181, 197),
			Brightness: 100,
		},
		{
			Duration:   2000,
			Mode:       1,
			Value:      y.GenerateRGB(234, 182, 118),
			Brightness: 100,
		},
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestStopColorFlow(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.StopColorFlow()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestSetScene(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetScene(yeelight.Scene{
		Action:     "color",
		Color:      65280,
		Brightness: 70,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestCronAdd(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.CronAdd(1)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestCronGet(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.CronGet()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestCronDelete(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.CronDelete()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestSetAdjust(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetAdjust("increase", "bright")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)

	r, err = y.SetAdjust("decrease", "bright")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestSetName(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.SetName("Bed Bulb")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestAdjustBright(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.AdjustBright(-10, 100)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}

func TestAdjustColorTemperature(t *testing.T) {
	y := yeelight.New(&yeelight.Config{
		IpAddress: "192.168.100.7",
		Port:      55443,
	})

	defer y.Close()

	r, err := y.AdjustColorTemperature(-10, 100)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(r)

	time.Sleep(3 * time.Second)
}
