package libaladin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Book struct {
	Title              string `json:"title"`
	Link               string `json:"link"`
	Author             string `json:"author"`
	PubDate            string `json:"pubDate"`
	Description        string `json:"description"`
	Isbn               string `json:"isbn"`
	Isbn13             string `json:"isbn13"`
	ItemId             int    `json:"itemId"`
	PriceSales         int    `json:"priceSales"`
	PriceStandard      int    `json:"priceStandard"`
	MallType           string `json:"mallType"`
	StockStatus        string `json:"stockStatus"`
	Mileage            int    `json:"mileage"`
	Cover              string `json:"cover"`
	CategoryId         int    `json:"categoryId"`
	CategoryName       string `json:"categoryName"`
	Publisher          string `json:"publisher"`
	SalesPoint         int    `json:"salesPoint"`
	Adult              bool   `json:"adult"`
	FixedPrice         bool   `json:"fixedPrice"`
	CustomerReviewRank int    `json:"customerReviewRank"`
	AstomerReviewRank  int    `json:"astomerReviewRank"`
	SeriesInfo         struct {
		SeriesId   int    `json:"seriesId"`
		SeriesLink string `json:"seriesLink"`
		SeriesName string `json:"seriesName"`
	} `json:"seriesInfo"`
	SubInfo struct {
		EbookList []interface{} `json:"ebookList"`
		UsedList  struct {
			AladinUsed struct {
				ItemCount int    `json:"itemCount"`
				MinPrice  int    `json:"minPrice"`
				Link      string `json:"link"`
			} `json:"aladinUsed"`
			UserUsed struct {
				ItemCount int    `json:"itemCount"`
				MinPrice  int    `json:"minPrice"`
				Link      string `json:"link"`
			} `json:"userUsed"`
		} `json:"usedList"`
		SubTitle      string `json:"subTitle"`
		OriginalTitle string `json:"originalTitle"`
		ItemPage      int    `json:"itemPage"`
	} `json:"subInfo"`
}

type ItemResult struct {
	Version            string `json:"version"`
	Logo               string `json:"logo"`
	Title              string `json:"title"`
	Link               string `json:"link"`
	PubDate            string `json:"pubDate"`
	TotalResults       int    `json:"totalResults"`
	StartIndex         int    `json:"startIndex"`
	ItemsPerPage       int    `json:"itemsPerPage"`
	Query              string `json:"query"`
	SearchCategoryId   int    `json:"searchCategoryId"`
	SearchCategoryName string `json:"searchCategoryName"`
	Item               []Book `json:"item"`
}

type SearchResult struct {
	Version            string `json:"version"`
	Logo               string `json:"logo"`
	Title              string `json:"title"`
	Link               string `json:"link"`
	PubDate            string `json:"pubDate"`
	TotalResults       int    `json:"totalResults"`
	StartIndex         int    `json:"startIndex"`
	ItemsPerPage       int    `json:"itemsPerPage"`
	Query              string `json:"query"`
	SearchCategoryId   int    `json:"searchCategoryId"`
	SearchCategoryName string `json:"searchCategoryName"`
	Item               []Book `json:"item"`
}

var myttb = "ttbdcmichael12561543002"

var apis = map[string]string{
	"search": "https://www.aladin.co.kr/ttb/api/ItemSearch.aspx?" +
		"ttbkey={apikey}&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"SearchTarget=Book&" +
		"MaxResults=100&" +
		"sort=title&" +
		"QueryType=Title&" +
		"Start={page}&" +
		"Query=%v&",
	"searchAuthor": "https://www.aladin.co.kr/ttb/api/ItemSearch.aspx?" +
		"ttbkey={apikey}&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"SearchTarget=Book&" +
		"MaxResults=100&" +
		"sort=title&" +
		"QueryType=Author&" +
		"Start={page}&" +
		"Query=%v&",
	"searchPublisher": "https://www.aladin.co.kr/ttb/api/ItemSearch.aspx?" +
		"ttbkey={apikey}&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"SearchTarget=Book&" +
		"MaxResults=100&" +
		"sort=title&" +
		"QueryType=Publisher&" +
		"Start={page}&" +
		"Query=%v&",
	"list": "https://www.aladin.co.kr/ttb/api/ItemList.aspx?" +
		"ttbkey={apikey}&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"QueryType=ItemNewAll&" +
		"start=1&" +
		"MaxResults=10&" +
		"SearchTarget=Book&",
	"itemISBN": "https://www.aladin.co.kr/ttb/api/ItemLookUp.aspx?" +
		"ttbkey={apikey}&" +
		"itemIdType=ISBN&" +
		"ItemId=%v&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"OptResult=ebookList,usedList,reviewList,previewImgList,eventList,authors,reviewList,fulldescription,fulldescription2,Toc,Story",
	"itemISBN13": "https://www.aladin.co.kr/ttb/api/ItemLookUp.aspx?" +
		"ttbkey={apikey}&" +
		"itemIdType=ISBN13&" +
		"ItemId=%v&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"OptResult=ebookList,usedList,reviewList,previewImgList,eventList,authors,reviewList,fulldescription,fulldescription2,Toc,Story",
	"itemID": "https://www.aladin.co.kr/ttb/api/ItemLookUp.aspx?" +
		"ttbkey={apikey}&" +
		"itemIdType=ISBN&" +
		"ItemId=%v&" +
		"output=js&" +
		"includeKey=1&" +
		"Version=20131101&" +
		"OptResult=ebookList,usedList,reviewList,previewImgList,eventList,authors,reviewList,fulldescription,fulldescription2,Toc,Story",
}

func SortBooksByTitle(books []Book) {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})
}

func SortBooksByISBN(books []Book) {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].Isbn < books[j].Isbn
	})
}

func SortBooksByISBN13(books []Book) {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].Isbn13 < books[j].Isbn13
	})
}
func SortBooksByID(books []Book) {
	sort.SliceStable(books, func(i, j int) bool {
		return books[i].ItemId < books[j].ItemId
	})
}

func makeQuery(querytype string, key string, page int) string {
	ret := strings.Replace(apis[querytype], "{apikey}", key, 1)
	ret = strings.Replace(ret, "{page}", fmt.Sprint(page), 1)
	return ret
}

func GetBook(isbn string) Book {
	fmt.Println(fmt.Sprintf(makeQuery("itemISBN", myttb, 1), isbn))
	resp, err := http.Get(fmt.Sprintf(makeQuery("itemISBN", myttb, 1), isbn))
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	//bookjson := fmt.Sprintf("%s", data)
	//fmt.Println(bookjson)
	var result ItemResult
	error := json.Unmarshal(data, &result)
	if error != nil {
		fmt.Printf("err: %v", err)
	}
	//b, _ := json.MarshalIndent(result, "", "  ")
	//fmt.Printf("book: %v\n", string(b))
	book := result.Item[0]
	return book
}

func SearchBookAuthor(id int) []Book {
	resp, err := http.Get(fmt.Sprintf(makeQuery("searchAuthor", myttb, 1), id))
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	//bookjson := fmt.Sprintf("%s", string(data))
	var result ItemResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	if len(result.Item) == result.ItemsPerPage {
		Item := SearchBookAuthors(id, 2)
		result.Item = append(result.Item, Item...)
	}
	//books_ :=
	SortBooksByTitle(result.Item)
	return result.Item
}

func SearchBook(name string) []Book {
	resp, err := http.Get(fmt.Sprintf(makeQuery("search", myttb, 1), url.QueryEscape(name)))
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	//bookjson := fmt.Sprintf("%s", string(data))
	var result ItemResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	if len(result.Item) == result.ItemsPerPage {
		Item := SearchBooks(name, 2)
		result.Item = append(result.Item, Item...)
	}
	//books_ :=
	SortBooksByTitle(result.Item)
	return result.Item
}

func SearchBooks(name string, page int) []Book {
	fmt.Println(fmt.Sprintf(makeQuery("search", myttb, page), url.QueryEscape(name)))
	resp, err := http.Get(fmt.Sprintf(makeQuery("search", myttb, page), url.QueryEscape(name)))
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	//bookjson := fmt.Sprintf("%s", string(data))
	var result ItemResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	if len(result.Item) == result.ItemsPerPage && page <= 4 {
		Item := SearchBooks(name, page+1)
		result.Item = append(result.Item, Item...)
	}
	return result.Item
}

func SearchBookAuthors(id int, page int) []Book {
	fmt.Println(fmt.Sprintf(makeQuery("searchAuthor", myttb, page), id))
	resp, err := http.Get(fmt.Sprintf(makeQuery("searchAuthor", myttb, page), id))
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	//bookjson := fmt.Sprintf("%s", string(data))
	var result ItemResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	if len(result.Item) == result.ItemsPerPage && page <= 4 {
		Item := SearchBookAuthors(id, page+1)
		result.Item = append(result.Item, Item...)
	}
	return result.Item
}
func Hello(name string) string {
	s := fmt.Sprintf("Hello, %v", name)
	return s
}
