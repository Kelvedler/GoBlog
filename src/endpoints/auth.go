package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const googleAuthScope = "email https://www.googleapis.com/auth/userinfo.email openid"
const googleAuthPrompt = "consent"

type GoogleResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

func GoogleAuthURL(ginCtx *gin.Context) {
	req, err := http.NewRequest(http.MethodGet, "https://accounts.google.com/o/oauth2/v2/auth", nil)
	query := req.URL.Query()
	query.Add("client_id", os.Getenv("OAUTH_CLIENT_ID"))
	query.Add("redirect_uri", fmt.Sprintf("%v/api/v1/jwt", os.Getenv("SERVER_ROOT_URI")))
	query.Add("access_type", "offline")
	query.Add("response_type", "code")
	query.Add("scope", googleAuthScope)
	query.Add("prompt", googleAuthPrompt)
	req.URL.RawQuery = query.Encode()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{"url": req.URL.String()})
}

func JWTToken(ginCtx *gin.Context) {
	ctx, ok := ginCtx.MustGet("topContext").(context.Context)
	if !ok {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		return
	}
	code := ginCtx.Query("code")
	scope := ginCtx.Query("scope")
	prompt := ginCtx.Query("prompt")
	if code == "" || scope != googleAuthScope || prompt != googleAuthPrompt {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	client := http.DefaultClient
	var respBody GoogleResp
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", os.Getenv("OAUTH_CLIENT_ID"))
	data.Set("redirect_uri", fmt.Sprintf("%v/api/v1/jwt", os.Getenv("SERVER_ROOT_URI")))
	data.Set("client_secret", os.Getenv("OAUTH_CLIENT_SECRET"))
	data.Set("grant_type", "authorization_code")
	req, err := http.NewRequest(http.MethodPost, "https://oauth2.googleapis.com/token", strings.NewReader(data.Encode()))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	payload, err := idtoken.Validate(ctx, respBody.IDToken, os.Getenv("OAUTH_CLIENT_ID"))
	if err != nil {
		ginCtx.JSON(http.StatusForbidden, gin.H{"error": err.Error})
		return
	}
	ginCtx.JSON(http.StatusOK, payload.Claims)
}
