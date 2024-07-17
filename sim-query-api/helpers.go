package simqueryapi

type Bundle struct {
	ID       string `json:"id"`
	Ul       string `json:"ul"`
	Dl       string `json:"dl"`
	Quota    string `json:"quota"`
	Duration string `json:"duration"`
	Label    string `json:"label"`
	Type     string `json:"type"`
}

type SimCard struct {
	Msisdn   string `json:"msisdn"`
	Imsi     string `json:"imsi"`
	Iccid    string `json:"iccid"`
	Secret   string `json:"secret"`
	Tac      string `json:"tac"`
	Eid      string `json:"eid"`
	Cid      string `json:"cid"`
	Imei     string `json:"imei"`
	BundleID Bundle `json:"bundle"`
}

func ResponseFormatter(simData map[string]string, bundle map[string]string) SimCard {
	return SimCard{
		Msisdn: simData["msisdn"],
		Imsi:   simData["imsi"],
		Iccid:  simData["iccid"],
		Secret: simData["secret"],
		Tac:    simData["tac"],
		Eid:    simData["eid"],
		Cid:    simData["cid"],
		Imei:   simData["imei"],
		BundleID: Bundle{
			ID:       bundle["id"],
			Ul:       bundle["ul"],
			Dl:       bundle["dl"],
			Quota:    bundle["quota"],
			Duration: bundle["duration"],
			Label:    bundle["label"],
			Type:     bundle["type"],
		},
	}
}
