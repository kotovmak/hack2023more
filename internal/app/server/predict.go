package server

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"hack2023/internal/app/model"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
)

func (s *server) RunOfficePredict(ctx context.Context, o []model.Office) ([]model.Office, error) {
	_, filename, _, _ := runtime.Caller(1)
	file, err := os.Create("/home/bitrix/www/hack2023/services/predict/data/offices.csv")
	if err != nil {
		log.Println(path.Dir(filename))
		return o, err
	}

	writer := csv.NewWriter(file)

	data := make([][]string, 0)
	headers := []string{"id", "salePointName"}
	for _, v := range o {
		str := []string{strconv.Itoa(v.ID), v.SalePointName}
		data = append(data, str)
	}

	writer.Write(headers)
	for _, row := range data {
		writer.Write(row)
	}

	writer.Flush()
	file.Close()

	log.Printf("[WORKER] Start python job")
	c := exec.Command(
		"python3",
		"/home/bitrix/www/hack2023/services/predict/offices.py",
	)

	if err := c.Run(); err != nil {
		log.Println(path.Join(path.Dir(filename), "../../../services/predict/offices.py"))
		return o, err
	}

	localFile, err := os.ReadFile("/home/bitrix/www/hack2023/services/predict/data/office_prediction.json")
	if err != nil {
		log.Println(path.Join(path.Dir(filename), "../../../services/predict/office_prediction.json"))
		return o, err
	}

	a := map[string]model.OfficePridict{}
	err = json.Unmarshal(localFile, &a)
	if err != nil {
		log.Printf("file parsing err %s", err)
		return o, err
	}

	newO := make(map[int]float64)
	for _, v := range a {
		newO[v.OfficeID] = v.Distance
	}

	for i, v := range o {
		k := 5 - newO[v.ID]
		k, err = strconv.ParseFloat(fmt.Sprintf("%.2f", k), 64)
		if err != nil {
			log.Printf("file parsing err %s", err)
			return o, err
		}
		o[i].Rating = k
	}

	return o, nil
}
