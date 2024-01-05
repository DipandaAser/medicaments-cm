package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"medicaments-cm/pkg/models"
	"net/http"
	"strings"
)

func SearchMedicaments(name string) ([]models.Medicament, error) {
	reqData := fmt.Sprintf("val=%s&choix=NomCommercial", name)
	request, err := http.NewRequest(http.MethodPost,
		"https://dpml.cm/repertoireDesAmm/func/recherche.php",
		strings.NewReader(reqData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status code error: %d %s", response.StatusCode, response.Status)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	medicaments := make([]models.Medicament, 0)

	doc.Find("table tbody tr").Each(func(i int, medicamentNode *goquery.Selection) {
		medicament := models.Medicament{
			AMM:                             medicamentNode.Find("td:nth-child(1)").Text(),
			ProductDesignation:              medicamentNode.Find("td:nth-child(2)").Text(),
			InternationalNonProprietaryName: medicamentNode.Find("td:nth-child(3)").Text(),
			LaboratoryAuthorizationHolder:   medicamentNode.Find("td:nth-child(4)").Text(),
			LaboratoryManufacturer:          medicamentNode.Find("td:nth-child(5)").Text(),
			AuthorizationStartDate:          medicamentNode.Find("td:nth-child(6)").Text(),
			AuthorizationEndDate:            medicamentNode.Find("td:nth-child(7)").Text(),
		}
		medicaments = append(medicaments, medicament)
	})

	return medicaments, nil
}
