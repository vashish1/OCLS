package v1

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/vashish1/OnlineClassPortal/backend/api/utility"
	"github.com/vashish1/OnlineClassPortal/backend/pkg/models"
)

func TeachersLogin(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json")
	var login models.Login
	var t models.LoginResponse

	err := c.BodyParser(&login)
	fmt.Println("user trying to login \n", login)
	if err != nil {
		t.Success = false
		t.Error = "Body not parsed"
		c.Status(400).JSON(t)
		return
	}

	token, err := utility.GenerateJwtForTeacher(login.Email, login.Password)
	if err != nil {
		t.Success = false
		t.Error = err.Error()
		c.Status(400).JSON(t)
		return
	}
	t.Success = true
	t.Token = token
	c.Status(200).JSON(t)
}

// func Tdashboard() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var dash models.Dash
// 		body, _ := ioutil.ReadAll(r.Body)
// 		err := json.Unmarshal(body, &dash)
// 		fmt.Println("email", dash)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(`{"error": "body not parsed"}`))
// 			return
// 		}
// 		teach, ok := teacher.Exist(dash.Email)
// 		if !ok {
// 			w.WriteHeader(500)
// 			w.Write([]byte(`{"error": "Email Invaid"}`))
// 			return
// 		}
// 		json.NewEncoder(w).Encode(teach)
// 	})
// }
