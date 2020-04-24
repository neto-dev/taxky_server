/*
 * Developer: Ernesto Valenzuela Vargas.
 * Created by: netodev
 * Contact: neto.dev@protonmail.com
 *
 */

/*
 * Revision History:
 *     Initial:      10/23/18  |  1:59 PM     Ernesto Valenzuela V
 */

package config

import (
	"os"
)

type DataBase struct {
	Dialect  string
	User     string
	Pass     string
	DbHost   string
	DataBase string
}

type SMTPSettings struct {
	Server string
	Port   string
	Auth   string
	UseTLS bool
}

type Config struct {
	Port       string
	Debug      bool
	DataBase   DataBase
	MailServer SMTPSettings
}

type Configurations struct {
	Development Config
	Test        Config
	Staging     Config
	Production  Config
}

type EnvironmentVariables struct {
	USER string
	PASS string
	HOST string
}

func Configuration() Configurations {
	/*
		Declaramos la variable s que sera de tipo Specification con la cual tendremos las propidades de la estructura
		para poder construir el objeto mas adelante

		We declare the variable s that will be of type Specification with which we will have the propieties of the structure
		to be able to build the object later.
	*/
	s := EnvironmentVariables{
		USER: os.Getenv("DATABASE_USER"),
		PASS: os.Getenv("DATABASE_PASS"),
		HOST: os.Getenv("DATABASE_HOST"),
	}

	var configurations = Configurations{
		Development: Config{
			Port:  ":3030",
			Debug: true,
			DataBase: DataBase{
				Dialect:  "mysql",
				User:     s.USER,
				Pass:     s.PASS,
				DbHost:   s.HOST,
				DataBase: "taxky_dev",
			},
			MailServer: SMTPSettings{
				Server: "email-smtp.us-west-2.amazonaws.com",
				Port:   "587",
				Auth:   "PLAIN", // PLAIN or CRAMMD5
				UseTLS: true,
			},
		}, Test: Config{
			Port:  ":7070",
			Debug: false,
			DataBase: DataBase{
				Dialect:  "mysql",
				User:     s.USER,
				Pass:     s.PASS,
				DbHost:   s.HOST,
				DataBase: "taxky_test",
			},
			MailServer: SMTPSettings{
				Server: "email-smtp.us-west-2.amazonaws.com",
				Port:   "465",
				Auth:   "PLAIN", // PLAIN or CRAMMD5
				UseTLS: true,
			},
		}, Staging: Config{
			Port:  ":7070",
			Debug: false,
			DataBase: DataBase{
				Dialect:  "mysql",
				User:     s.USER,
				Pass:     s.PASS,
				DbHost:   s.HOST,
				DataBase: "test_mvc_staging",
			},
			MailServer: SMTPSettings{
				Server: "email-smtp.us-west-2.amazonaws.com",
				Port:   "587",
				Auth:   "PLAIN", // PLAIN or CRAMMD5
				UseTLS: true,
			},
		}, Production: Config{
			Port:  ":6060",
			Debug: false,
			DataBase: DataBase{
				Dialect:  "mysql",
				User:     s.USER,
				Pass:     s.PASS,
				DbHost:   s.HOST,
				DataBase: "munkky",
			},
			MailServer: SMTPSettings{
				Server: "email-smtp.us-west-2.amazonaws.com",
				Port:   "25",
				Auth:   "PLAIN", // PLAIN or CRAMMD5
				UseTLS: true,
			},
		},
	}

	return configurations
}

var Environment = map[string]Config{
	"Development": Configuration().Development,
	"Test":        Configuration().Test,
	"Staging":     Configuration().Staging,
	"Production":  Configuration().Production,
}
