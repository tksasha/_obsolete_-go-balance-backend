package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/tksasha/date"

	. "github.com/tksasha/balance/config/test"
	. "github.com/tksasha/balance/models"
)

func TestBuildErrors(t *testing.T) {
	item := new(Item)

	item.Build(Values)

	assert.NotPanics(t, func() {
		item.Errors.Add("foo", "is invalid")
	})
}

func TestMarshalJSON(t *testing.T) {
	assert := assert.New(t)

	category := Category{
		ID:   36,
		Name: "Food",
	}

	item := Item{
		ID:          57,
		Date:        date.New("1982-05-17").Time(),
		Sum:         10.52,
		Description: "Red Dry Wine",
		Category:    category,
	}

	m, _ := json.Marshal(item)

	var d ItemDecorator

	json.Unmarshal(m, &d)

	assert.Equal(57, d.ID)

	assert.Equal("1982-05-17", d.Date)

	assert.Equal(10.52, d.Sum)

	assert.Equal("Red Dry Wine", d.Description)

	assert.Equal(36, d.Category.ID)

	assert.Equal("Food", d.Category.Name)
}

func TestBuildWithoutParams(t *testing.T) {
	item := new(Item)

	item.Build(Values)

	assert.True(t, item.Date.IsZero())

	assert.Zero(t, item.CategoryID)

	assert.Zero(t, item.Sum)

	assert.Empty(t, item.Description)
}

func TestBuildWithParams(t *testing.T) {
	item := new(Item)

	Values.Set("item[date]", "2007-01-28")

	Values.Set("item[category_id]", "42")

	Values.Set("item[sum]", "14.44")

	Values.Set("item[description]", "Red Dry Wine")

	item.Build(Values)

	assert.Equal(t, "2007-01-28", item.Date.Format("2006-01-02"))

	assert.Equal(t, 42, item.CategoryID)

	assert.Equal(t, 14.44, item.Sum)

	assert.Equal(t, "Red Dry Wine", item.Description)
}

func TestBuildWithEmptyValues(t *testing.T) {
	item := new(Item)

	Values.Set("item[date]", "")

	Values.Set("item[category_id]", "")

	Values.Set("item[sum]", "")

	Values.Set("item[description]", "")

	item.Build(Values)

	assert.Equal(t, new(time.Time).Format("2006-01-02"), item.Date.Format("2006-01-02"))

	assert.Equal(t, 0, item.CategoryID)

	assert.Equal(t, 0.0, item.Sum)

	assert.Equal(t, "", item.Description)
}

func TestValidateSumGreaterThanZero(t *testing.T) {
	item := new(Item)

	Values.Set("item[sum]", "")

	item.Build(Values)

	assert.False(t, item.IsValid())

	assert.Equal(t, []string{"must be greater than 0"}, item.Errors.Get("sum"))
}

func TestValidatePresenceOfCategory(t *testing.T) {
	item := new(Item)

	Values.Set("item[category_id]", "")

	item.Build(Values)

	assert.False(t, item.IsValid())

	assert.Equal(t, []string{"can't be blank"}, item.Errors.Get("category_id"))
}

func TestValidatePresenceOfDate(t *testing.T) {
	item := new(Item)

	delete(Values, "item[date]")

	item.Build(Values)

	assert.False(t, item.IsValid())

	assert.Equal(t, []string{"can't be blank"}, item.Errors.Get("date"))
}

func TestValidatePresenceOfDateWithEmptyString(t *testing.T) {
	item := Item{Date: time.Now()}

	Values.Set("item[date]", "")

	item.Build(Values)

	assert.False(t, item.IsValid())

	assert.Equal(t, []string{"can't be blank"}, item.Errors.Get("date"))
}
