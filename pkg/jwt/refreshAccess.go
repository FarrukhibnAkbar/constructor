package jwt

// func RefreshAccessToken(refreshTokenString string) (entities.Tokens, error) {
// 	var cfg = configs.Config()
// 	tokens := entities.Tokens{}
// 	// Parse and verify the refresh token
// 	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("error Parse refresh token. BadRequest")
// 		}
// 		return []byte(cfg.JWTSecretKey), nil
// 	})

// 	if err != nil || !refreshToken.Valid {
// 		return entities.Tokens{}, fmt.Errorf("invalid refresh token")
// 	}

// 	// Generate a new access token if refresh token is valid
// 	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
// 		id := claims["id"].(string)
// 		username := claims["username"].(string)
// 		role := claims["role"].(string)

// 		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 			"username": username,
// 			"role":     role,
// 			"id":       id,
// 			"exp":      time.Now().Add(time.Minute * 15).Unix(),
// 		})

// 		newAccessTokenString, err := newAccessToken.SignedString([]byte(cfg.JWTSecretKey))
// 		if err != nil {
// 			return entities.Tokens{}, fmt.Errorf("failed to generate new access token")
// 		}

// 		tokens.AccessToken = newAccessTokenString
// 		tokens.RefreshToken = refreshTokenString

// 	} else {
// 		return entities.Tokens{}, fmt.Errorf("invalid refresh token")
// 	}

// 	return tokens, nil
// }
