package testing

import (
	"bytes"
	"encoding/json"
	"golang-base/controllers"
	"golang-base/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	SetUpMockDB()
	defer RemoveMockDB()

	mockResponse := models.User{Base: models.Base{ID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")}, Email: "test1@localhost.com", Mobile: "0801", Username: "test1", Password: "test1test1"}

	if err := models.DB.Create(&mockResponse).Error; err != nil {
		t.Fatal(err)
	}

	r := SetUpRouter()
	r.GET("/users/:id", controllers.GetUsers)
	req, _ := http.NewRequest("GET", "/users/"+mockResponse.ID.String(), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseDataBuf, _ := ioutil.ReadAll(w.Body)
	responseData := []models.User{}

	t.Log(string(responseDataBuf))

	if err := json.Unmarshal(responseDataBuf, &responseData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockResponse.ID, responseData[0].ID)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetUsers(t *testing.T) {
	SetUpMockDB()
	mockResponse := []models.User{
		{Base: models.Base{ID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")}, Email: "test1@localhost.com", Mobile: "0801", Username: "test1", Password: "test1test1"},
		{Base: models.Base{ID: uuid.FromStringOrNil("00000000-0000-0000-0000-000000000002")}, Email: "test2@localhost.com", Mobile: "0802", Username: "test2", Password: "test2test2"},
	}

	if err := models.DB.Create(&mockResponse[0]).Error; err != nil {
		t.Fatal(err)
	}

	if err := models.DB.Create(&mockResponse[1]).Error; err != nil {
		t.Fatal(err)
	}

	r := SetUpRouter()
	r.GET("/users", controllers.GetUsers)
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseDataBuf, _ := ioutil.ReadAll(w.Body)
	responseData := []models.User{}
	if err := json.Unmarshal(responseDataBuf, &responseData); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(mockResponse), len(responseData))
	assert.Equal(t, mockResponse[0].Username, responseData[0].Username)
	assert.Equal(t, mockResponse[1].Username, responseData[1].Username)
	assert.Equal(t, http.StatusOK, w.Code)

	models.DB.Where("1=1").Delete(&models.User{})
}

func TestPostUser(t *testing.T) {
	SetUpMockDB()
	defer RemoveMockDB()

	mockRequest := controllers.ReqNewUser{Email: "test1@localhost.com", Mobile: "0801", Username: "test1x", Password: "test1test1"}
	var mockRequestBuf bytes.Buffer
	if err := json.NewEncoder(&mockRequestBuf).Encode(mockRequest); err != nil {
		t.Fatal(err)
	}

	r := SetUpRouter()
	r.POST("/users", controllers.PostUser)
	req, _ := http.NewRequest("POST", "/users", &mockRequestBuf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	t.Log("responseData: ", string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
	models.DB.Where("1=1").Delete(&models.User{})
}
