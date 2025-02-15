package services

import (
    "time"

    "github.com/go-redis/redis/v8"
    "gopkg.in/oauth2.v4"
    "gopkg.in/oauth2.v4/manage"
    "gopkg.in/oauth2.v4/models"
    "gopkg.in/oauth2.v4/server"
    "gopkg.in/oauth2.v4/store"
)

var OAuth2Server *server.Server

func InitOAuth2Service(redisClient *redis.Client) {
    manager := manage.NewDefaultManager()
    manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

    // Use Redis as the token store
    manager.MapTokenStorage(store.NewRedisStoreWithCli(redisClient))

    // Use in-memory client store for simplicity
    clientStore := store.NewClientStore()
    clientStore.Set("client_id_1", &models.Client{
        ID:     "client_id_1",
        Secret: "client_secret_1",
        Domain: "http://localhost:8081",
    })
    manager.MapClientStorage(clientStore)

    OAuth2Server = server.NewDefaultServer(manager)
    OAuth2Server.SetAllowGetAccessRequest(true)
    OAuth2Server.SetClientInfoHandler(server.ClientFormHandler)

    OAuth2Server.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
        // Handle user authentication here
        return "user_id", nil
    })

    OAuth2Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
        return
    })

    OAuth2Server.SetResponseErrorHandler(func(re *errors.Response) {
    })
}