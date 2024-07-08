package currency

import (
	"encoding/xml"
	"errors"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/application/core/domain"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/cache"
	"github.com/aknEvrnky/currency-api-hexogonal/internal/ports"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	DEFAULT_ENDPOINT = "https://www.tcmb.gov.tr/kurlar/today.xml"
)

// define errors objects
var (
	ErrCurrencyNotFound = errors.New("currency not found")
	ErrApiError         = errors.New("api error")
)

type Adapter struct {
	cache    ports.CachePort
	endpoint string
}

func NewAdapter(cache ports.CachePort) *Adapter {
	return &Adapter{
		cache:    cache,
		endpoint: DEFAULT_ENDPOINT,
	}
}

type currency struct {
	Unit            uint    `xml:"Unit"`
	Name            string  `xml:"Isim"`
	CurrencyCode    string  `xml:"CurrencyCode,attr"`
	CurrencyName    string  `xml:"CurrencyName"`
	ForexBuying     string  `xml:"ForexBuying"`
	ForexSelling    string  `xml:"ForexSelling"`
	BanknoteBuying  float64 `xml:"BanknoteBuying"`
	BanknoteSelling float64 `xml:"BanknoteSelling"`
}

type currencies struct {
	Currencies []currency `xml:"Currency"`
}

func (a *Adapter) GetList() ([]domain.Currency, error) {
	currencyEntries, err := a.cache.Remember("currencies", 10*time.Second, func() (cache.Value, error) {
		log.Println("fetching from api")

		var currencyEntries []domain.Currency

		// make a http request to the endpoint
		res, err := http.Get(a.endpoint)
		defer res.Body.Close()
		if err != nil {
			return nil, ErrApiError
		}

		// check the response status code
		if res.StatusCode != http.StatusOK {
			return nil, ErrApiError
		}

		// parse the response body
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, ErrApiError
		}

		// unmarshal the response body
		var c currencies
		err = xml.Unmarshal(data, &c)

		if err != nil {
			return nil, ErrApiError
		}

		// map the response to domain.Currency
		for _, cur := range c.Currencies {
			currencyEntries = append(currencyEntries, domain.Currency{
				Code:        cur.CurrencyCode,
				Title:       cur.CurrencyName,
				Unit:        cur.Unit,
				BuyingRate:  cur.BanknoteBuying,
				SellingRate: cur.BanknoteSelling,
			})
		}

		return currencyEntries, nil
	})

	if err != nil {
		return nil, err
	}

	return currencyEntries.([]domain.Currency), nil
}

func (a *Adapter) GetByCurrencyCode(currencyCode string) (domain.Currency, error) {
	currencies, err := a.GetList()
	if err != nil {
		return domain.Currency{}, err
	}

	for _, currency := range currencies {
		if currency.Code == currencyCode {
			return currency, nil
		}
	}

	return domain.Currency{}, ErrCurrencyNotFound
}
