package utils

import (
	"errors"
	"gopkg.in/olivere/elastic.v3"
	"log"
)

type ES struct {
	es_host      string
	es_index     string
	es_Type      string
	logHandle    *log.Logger
	ES_Client    *elastic.Client
	bool_query   *elastic.BoolQuery
	aggregations map[string]elastic.Aggregation
	size         int
	form         int
	sort         map[string]bool
}

// FinderResponse is the outcome of calling Finder.Find.
type ESResponse struct {
	Total  int64
	Result []*string
	Aggs   elastic.Aggregations
}

//New ES
func NewES(host string, logHandle *log.Logger) *ES {
	//retry
	retry := elastic.NewBackoffRetrier(elastic.ZeroBackoff{})

	//connect
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetInfoLog(logHandle),
		elastic.SetRetrier(retry),
	)
	if err != nil {
		panic(err)
	}
	return &ES{
		es_host:      host,
		logHandle:    logHandle,
		ES_Client:    client,
		bool_query:   elastic.NewBoolQuery(),
		aggregations: make(map[string]elastic.Aggregation),
		size:         10000,
		form:         0,
		sort:         make(map[string]bool),
	}
}

//Set ES Index
func (e *ES) Index(es_index string) *ES {
	e.es_index = es_index
	return e
}

//Set ES Type
func (e *ES) Type(es_type string) *ES {
	e.es_Type = es_type
	return e
}

//ES Where Condition search
func (e *ES) Where(field string, value interface{}) *ES {
	e.bool_query = e.bool_query.Must(elastic.NewTermQuery(field, value))
	return e
}

//Set Size
func (e *ES) Take(size int) *ES {
	e.size = size
	return e
}

//Set Search From
func (e *ES) Page(page int) *ES {
	e.form = e.size * (page - 1)
	return e
}

//ES Range Condition search
func (e *ES) Range(field string, values map[string]interface{}) *ES {
	range_query := elastic.NewRangeQuery(field)
	for k, v := range values {
		if k == "gt" {
			range_query.Gt(v)
		}
		if k == "gte" {
			range_query.Gte(v)
		}
		if k == "lt" {
			range_query.Lt(v)
		}
		if k == "lte" {
			range_query.Lte(v)
		}
	}
	e.bool_query = e.bool_query.Must(range_query)
	return e
}

//ES Avg Aggregation
func (e *ES) Avg(field string) *ES {
	avg := elastic.NewAvgAggregation().Field(field)
	e.aggregations["avg"] = avg
	return e
}

//ES Max Aggregation
func (e *ES) Max(field string) *ES {
	max := elastic.NewMaxAggregation().Field(field)
	e.aggregations["max"] = max
	return e
}

//ES Min Aggregation
func (e *ES) Min(field string) *ES {
	min := elastic.NewMinAggregation().Field(field)
	e.aggregations["min"] = min
	return e
}

//ES Sum Aggregation
func (e *ES) Sum(field string) *ES {
	sum := elastic.NewSumAggregation().Field(field)
	e.aggregations["sum"] = sum
	return e
}

//ES Count Aggregation
func (e *ES) Count(field string) *ES {
	count := elastic.NewValueCountAggregation().Field(field)
	e.aggregations["count"] = count
	return e
}

//ES GroupBy Aggregation
func (e *ES) GroupBy(field string) *ES {
	groupBy := elastic.NewTermsAggregation().Field(field).Size(e.size)
	e.aggregations["groupBy"] = groupBy
	return e
}

//ES OrderBy Query
func (e *ES) OrderBy(field string, order string) *ES {
	var dir bool
	if order == "desc" {
		dir = false
	} else {
		dir = true
	}
	e.sort[field] = dir
	return e
}

//ES Create Index
func (e *ES) CreateIndex(index string) (err error) {
	exists, err := e.ES_Client.IndexExists(index).Do()
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := e.ES_Client.CreateIndex(index).Do()
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("create index failed")
		}
	}
	return nil
}

//ES Delete Index
func (e *ES) DeleteIndex(index string) (err error) {
	exists, err := e.ES_Client.IndexExists(index).Do()
	if err != nil {
		return err
	}
	if exists {
		_, err = e.ES_Client.DeleteIndex(index).Do()
		if err != nil {
			return err
		}
	} else {
		return errors.New("index: " + index + " not exist")
	}
	return nil
}

//ES Get Response
func (e *ES) Search() (ESResponse, error) {
	var resp ESResponse

	// Create service and use query, aggregations, filter
	search := e.ES_Client.Search().
		Index(e.es_index).
		Type(e.es_Type).
		Query(e.bool_query).
		Size(e.size).
		From(e.form)

	// Add aggregation
	if len(e.aggregations) > 0 {
		for k, v := range e.aggregations {
			search = search.Aggregation(k, v)
		}
	}

	// Add SortBy
	if len(e.sort) > 0 {
		for k, v := range e.sort {
			search = search.Sort(k, v)
		}
	}

	// Execute query
	sr, err := search.Do()
	if err != nil {
		return resp, err
	}

	// Decode response
	rs, aggs, err := e.decodeResponse(sr)
	if err != nil {
		return resp, err
	}
	resp.Result = rs
	resp.Total = sr.Hits.TotalHits
	resp.Aggs = aggs
	return resp, nil
}

//DecodeLogs takes a search result and deserializes the response.
func (e *ES) decodeResponse(res *elastic.SearchResult) ([]*string, elastic.Aggregations, error) {
	if res == nil || res.TotalHits() == 0 {
		return nil, nil, nil
	}

	var rss []*string
	for _, hit := range res.Hits.Hits {
		tmp,_ := hit.Source.MarshalJSON()
		tmps := string(tmp)
		rss = append(rss, &tmps)
	}
	return rss, res.Aggregations, nil
}
