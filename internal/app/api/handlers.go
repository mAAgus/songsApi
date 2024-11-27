package api

import (
	"encoding/json"
	"fmt"
	"go/scr/hhruxongs/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content_Type", "application/json")
}
func (api *API) DeleteSongsById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Song by ID DELET /api/v1/song/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok, ers := api.storage.Songs().FindById(id)
	if ers != nil {
		api.logger.Info("Troubles while accessing database table(songs) with id. Error:", ers)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can't find song with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Song with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, erd := api.storage.Songs().DeleteSongById(id)
	if erd != nil {
		api.logger.Info("Troubles while deleting database table(songs) with id. Error:", ers)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Song with ID %d successfully deleted", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}
func (api *API) PutSongs(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Put Song PUT /api/v1/song/{id}")
	var song models.Songs
	err := json.NewDecoder(req.Body).Decode(&song)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	upA, ok, err := api.storage.Songs().UpdteSongs(id, &song)
	if err != nil {
		api.logger.Info("Troubles while creating new song", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accesting database. Try agin.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	if !ok {
		api.logger.Info("Can't find song with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Song with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(upA)

}
func (api *API) PostSong(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Song POST /api/v1/song")
	var song models.Songs
	err := json.NewDecoder(req.Body).Decode(&song)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, ers := api.storage.Songs().Create(&song)
	if ers != nil {
		api.logger.Info("Troubles while creating new song", ers)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accesting database. Try agin.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}
func (api *API) GetSongsIdPart(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get article by ID GET /api/v1/song/{id}/{page}/{pageSize}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use ID as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	page, err := strconv.Atoi(mux.Vars(req)["page"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {page} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use Page as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	pageSize, err := strconv.Atoi(mux.Vars(req)["pageSize"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {pageSize} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use PageSize as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	song, ok, ers := api.storage.Songs().GetSongsParts(id, page, pageSize)
	if ers != nil {
		api.logger.Info("Troubles while accessing database table(articles) with id. Error:", ers)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can't find song with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(song)
}
func (api *API) GetSongs(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get article by ID GET /api/v1/song/{filter}/{page}/{pageSize}")
	filter := mux.Vars(req)["filter"]
	page, err := strconv.Atoi(mux.Vars(req)["page"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {page} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use Page as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	pageSize, err := strconv.Atoi(mux.Vars(req)["pageSize"])
	if err != nil {
		api.logger.Info("Troubles wile parsing {pageSize} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. Don't use PageSize as uncasting to int value",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	songs, err := api.storage.Songs().SelectAllFiltrPangination(filter, page, pageSize)
	if err != nil {
		api.logger.Info("Error while Songs.SelectAllFiltrPangination: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trobles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(songs)
}
