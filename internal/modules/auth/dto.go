package auth

// RegisterDTO datos para registro
// @name RegisterDTO
type RegisterDTO struct {
	// Nombre del usuario
	Name string `json:"name" binding:"required,min=2" example:"John Doe"`
	// Email del usuario
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
	// Contraseña (mínimo 6 caracteres)
	Password string `json:"password" binding:"required,min=6" example:"secret123"`
}

// LoginDTO datos para login
// @name LoginDTO
type LoginDTO struct {
	// Email del usuario
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
	// Contraseña
	Password string `json:"password" binding:"required" example:"secret123"`
}

// RefreshDTO datos para refresh token
// @name RefreshDTO
type RefreshDTO struct {
	// Refresh token
	RefreshToken string `json:"refresh_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// TokenResponse respuesta con tokens
// @name TokenResponse
type TokenResponse struct {
	// Access token JWT
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	// Refresh token JWT
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	// Tiempo de expiración en segundos
	ExpiresIn int64 `json:"expires_in" example:"900"`
}

// UserInfo información del usuario autenticado
// @name UserInfo
type UserInfo struct {
	// ID del usuario
	ID string `json:"id" example:"507f1f77bcf86cd799439011"`
	// Nombre del usuario
	Name string `json:"name" example:"John Doe"`
	// Email del usuario
	Email string `json:"email" example:"john@example.com"`
}
