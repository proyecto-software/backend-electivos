package function

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Validator(rut string, logger *logrus.Entry, c *gin.Context) bool {
	res := strings.ReplaceAll(rut, ".", "")
	split := strings.Split(res, "-")
	number, err := strconv.Atoi(split[0])
	if err != nil {
		logger.Errorf("error en el rut")
		//c.JSON(400, "RUT invalido")
	}
	code, err := strconv.Atoi(split[1])
	if err != nil {
		if split[1] == "K" {
			code = 10
		} else {
			logger.Errorf("error en el codigo verificador")
			//c.JSON(400, "Codigo verificador invalido")
			return false
		}
	}
	if number > 5000000 && number < 30000000 {

		suma := 0
		multiplo := 2
		for i, item := range Reverse(split[0]) {
			if err != nil {
				logger.Errorf("error en el rut")
				c.JSON(200, "OK")
				return false
			}
			n, err := strconv.Atoi(string(item))
			if err != nil {
				logger.Errorf("error en el rut")
				return false
			}
			index := multiplo * n
			suma = suma + index
			if multiplo < 7 {
				multiplo = multiplo + 1
			} else {
				multiplo = 2
			}
			i = i
		}
		dv_esperado := 11 - (suma % 11)
		if code == 0 {
			code = 11
		}
		if code != dv_esperado {
			logger.Errorf("rut invalido")
			//c.JSON(400, "Formato valido, rut invalido")
			return false
		} else {
			logger.Errorf("rut valido")
			return true
		}

	} else {
		logger.Errorf("error en el rut")
		//c.JSON(400, "Error en el RUT")
		return false
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ValidatorCorreo(correo string, logger *logrus.Entry, c *gin.Context) (valid bool) {
	split := strings.Split(correo, "@")
	CorreoValido := []string{"alumnos.ucn.cl", "ucn.cl"}

	for i := 0; i < len(CorreoValido); i++ {
		if split[1] == CorreoValido[i] {
			valid = true
		}
	}
	if !valid {
		logger.Errorf("correo invalido")
		c.JSON(400, "Formato no valido, correo invalido")
		return false
	} else {
		return true
	}
}
