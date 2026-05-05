package middleware

import ("net/http"
		"context"
		"github.com/a-h/templ"
		"github.com/gattini0928/Equilibrium/internal/services/auth"
	)


func IsAuthenticated(r *http.Request) bool {
    return r.Context().Value("user_id") != nil
}

func AuthMiddleware(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		userID, err := auth.ValidateJWT(secret, cookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) {
    _ = c.Render(r.Context(), w)
}


