package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/volatiletech/null.v6"
)

func (im *ItemModule) AddItem(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ItemRequest{ItemRequested: Item{}, CookieString: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.ItemRequested.Name) == 0 ||
		len(request.ItemRequested.Manufacturer) == 0 ||
		len(request.ItemRequested.Category) == 0 ||
		len(request.ItemRequested.SubCategory) == 0 {
		out.Msg = "name, manufacturer/creator and category/sub_category cannot be zero."
		return
	}

	_item, e := models.Items(im.DataBase, qm.Where("name = ? AND manufacturer = ? AND category = ? AND sub_category = ?",
		request.ItemRequested.Name,
		request.ItemRequested.Manufacturer,
		request.ItemRequested.Category,
		request.ItemRequested.SubCategory)).
		One()

	if e == nil && _item != nil {
		out.Msg = "Item exist in our database"
		out.Response = map[string]interface{}{
			"item": Item{
				ID:                _item.ID,
				Name:              _item.Name.String,
				Manufacturer:      _item.Manufacturer.String,
				Category:          _item.Category.String,
				SubCategory:       _item.SubCategory.String,
				SubSubCategory:    _item.SubSubCategory.String,
				SubSubSubCategory: _item.SubSubSubCategory.String,
				RegionCountry:     _item.RegionCountry.String,
				RegionState:       _item.RegionCity.String,
				RegionCity:        _item.RegionCity.String,
				RegionPin:         _item.RegionPin.String,
				AliasName:         _item.AliasName.String,
				ItemUrl:           _item.Itemurl.String,
				Owner:             _item.OwnerName.String,
				CreatedOn:         _item.CreatedOn.Time,
				ExpiredOn:         _item.ExpiryOn.Time,
				IsExpired:         _item.HasExpired.Valid,
			},
		}
		return
	}

	/* Else, we are here to add this item that user has requested */

	item := models.Item{
		ID:                request.ItemRequested.ID,
		Name:              null.StringFrom(request.ItemRequested.Name),
		Manufacturer:      null.StringFrom(request.ItemRequested.Manufacturer),
		Category:          null.StringFrom(request.ItemRequested.Category),
		SubCategory:       null.StringFrom(request.ItemRequested.SubCategory),
		SubSubCategory:    null.StringFrom(request.ItemRequested.SubSubCategory),
		SubSubSubCategory: null.StringFrom(request.ItemRequested.SubSubSubCategory),
		RegionCountry:     null.StringFrom(request.ItemRequested.RegionCountry),
		RegionState:       null.StringFrom(request.ItemRequested.RegionCity),
		RegionCity:        null.StringFrom(request.ItemRequested.RegionCity),
		RegionPin:         null.StringFrom(request.ItemRequested.RegionPin),
		AliasName:         null.StringFrom(request.ItemRequested.AliasName),
		Itemurl:           null.StringFrom(request.ItemRequested.ItemUrl),
		OwnerName:         null.StringFrom(request.ItemRequested.Owner),
		CreatedOn:         null.TimeFrom(request.ItemRequested.CreatedOn),
		ExpiryOn:          null.TimeFrom(request.ItemRequested.ExpiredOn),
	}

	if request.ItemRequested.IsExpired {
		item.HasExpired = null.Int8From(1)
	} else {
		item.HasExpired = null.Int8From(0)
	}

	if e := item.Insert(im.DataBase); e != nil {
		out.Msg = e.Error()
		return
	}

	/* Done , Add success */
	out.Code = 0
	out.Msg = "created"
	out.Response = map[string]interface{}{
		"item_id": item.ID,
	}

	return

}
