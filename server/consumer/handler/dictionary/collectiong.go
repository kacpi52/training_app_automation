package dictionary

import (
	params_data "myInternal/consumer/data"
	dictionary_data "myInternal/consumer/data/dictionary"
	database "myInternal/consumer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionDictionary struct {
	Collection []dictionary_data.ResponseCollection `json:"collection"`
	Status     int                          `json:"status"`
	Error      string                       `json:"error"`
}

func HandlerCollectionDictionary(c *gin.Context){

	params := params_data.Params{
		AppLanguage: c.GetHeader("AppLanguage"),
	}
	
	collection, err := CollectionDictionary(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollectionDictionary{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCollectionDictionary{
		Collection: collection.Collection,
		Status: collection.Status,
		Error: collection.Error,
	})
}


func CollectionDictionary(params params_data.Params)(ResponseCollectionDictionary, error){

	var dictionaryCollection []dictionary_data.Collection
	var dictionaryMultiCollection []dictionary_data.MultiCollection
	var finalDictionaryCollection []dictionary_data.ResponseCollection
	var dictionaryId = params.AppLanguage

	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionDictionary{}, err
    }
	defer db.Close()

	query := `SELECT * FROM dictionary`
	rows, err := db.Query(query)
    if err != nil {
        return ResponseCollectionDictionary{}, err
    }
	defer rows.Close()

	for rows.Next() {
		var dictionary dictionary_data.Collection
		if err := rows.Scan(&dictionary.Id, &dictionary.CreatedUp, &dictionary.UpdateUp); err != nil {
			return ResponseCollectionDictionary{}, err
		}
		dictionaryCollection = append(dictionaryCollection, dictionary)
	}

	query = `SELECT * FROM dictionary_multi_language WHERE "dictionaryId" = $1`
	rows, err = db.Query(query, dictionaryId)
    if err != nil {
        return ResponseCollectionDictionary{}, err
    }
	defer rows.Close()

	for rows.Next() {
		var dictionary dictionary_data.MultiCollection
		if err := rows.Scan(&dictionary.Id, &dictionary.DictionaryId, &dictionary.Key, &dictionary.Translation); err != nil {
			return ResponseCollectionDictionary{}, err
		}
		dictionaryMultiCollection = append(dictionaryMultiCollection, dictionary)
	}

	for index, value := range dictionaryCollection {
		dictionary := dictionary_data.ResponseCollection{
			Id:          value.Id,
			Key:         dictionaryMultiCollection[index].Key,
			Translation: dictionaryMultiCollection[index].Translation,
			CreatedUp:   value.CreatedUp,
			UpdateUp:    value.UpdateUp,
		}
		finalDictionaryCollection = append(finalDictionaryCollection, dictionary)
	}


	return ResponseCollectionDictionary{
		Collection: finalDictionaryCollection,
		Status: http.StatusOK,
		Error: "",
	}, nil
}