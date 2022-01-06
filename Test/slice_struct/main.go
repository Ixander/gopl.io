package main

import "fmt"

func main() {
	type PayReqInfo struct {
		Code       int     `json:"code"`
		Name       string  `json:"name"`
		StrCode    string  `json:"strCode"`
		SumWithNds float64 `json:"sumWithNds"`
	}

	type PayChInfo struct {
		ChId  int    `json:"chId"`
		Brmn  string `json:"brmn"`
		State string `json:"state"`
	}

	info1 := PayChInfo{
		ChId:  9,
		Brmn:  "brmn1",
		State: "stateInfo1",
	}

	info2 := PayChInfo{
		ChId:  10,
		Brmn:  "brmn2",
		State: "stateInfo2",
	}

	sliceChInfo := []PayChInfo{info1, info2}

	type PayInfo struct {
		ReqId          int          `json:"reqId"`
		CreateDate     string       `json:"createDate"`
		AuthorLdap     string       `json:"authorLdap"`
		State          int          `json:"state"`
		ReqSum         float64      `json:"reqSum"`
		Pdoc           string       `json:"pdoc"`
		PayWhy         string       `json:"payWhy"`
		PayPeriodFrom  string       `json:"payPeriodFrom"`
		PayPeriodTo    string       `json:"payPeriodTo"`
		AccPayDate     string       `json:"accPayDate"`
		NomenklReqInfo []PayReqInfo `json:"nomenklReqInfo"`
		NomenklChInfo  []PayChInfo  `json:"nomenklChInfo"`
	}

	c := PayInfo{
		ReqId:         9,
		ReqSum:        500,
		NomenklChInfo: sliceChInfo,
	}
	//fmt.Println(c.ReqId)

	for _, y := range c.NomenklChInfo {
		//fmt.Println(x)
		fmt.Println(y.ChId)
		fmt.Println(y.Brmn)
	}

	//fmt.Println(c.NomenklChInfo)

}
