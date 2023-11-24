package main

import (
	"flag"
	"os"
	"strings"
)

var (
	prem              os.FileMode = 0777
	modelPath                     = "./models"
	dataModel                     = "/dataModel"
	dataSources                   = "/dataSources"
	repositroy                    = "/repository.go"
	redisDS                       = "redisDS"
	vitessDS                      = "vitessDS"
	vittesFileName                = "vitess.go"
	redsiFileName                 = "redis.go"
	DBDSFileName                  = "DBDS.go"
	cahceDSFileName               = "cacheDS.go"
	receiversFileName             = "receivers.go"
)

var (
	hasModel      = flag.Bool("model", true, "model")
	hasDataSource = flag.Bool("dataSource", true, "dataSource")
	hasDB         = flag.Bool("DB", true, "DB")
	hasCache      = flag.Bool("cache", true, "cache")
	modelName     = flag.String("modelName", "", "modelName")
)

func main() {
	flag.Parse()
	if *modelName == "" {
		panic("modelName empty")
	}

	y := *modelName
	lowerModelName := strings.ToLower(string(y[0])) + y[1:]
	upperModelName := strings.ToUpper(string(y[0])) + y[1:]
	var rep = strings.NewReplacer(
		"{lowerModelName}", lowerModelName,
		"{upperModelName}", upperModelName,
		"{dataSources}", dataSources[1:],
		"{tableName}", lowerModelName+"s",
		"{dataModel}", dataModel[1:],
		"{redisDS}", redisDS,
		"{vitessDS}", vitessDS,
	)

	modelPath = modelPath + "/" + lowerModelName
	err := os.RemoveAll(modelPath)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(modelPath, prem)
	if err != nil {
		panic(err)
	}
	repoFile, err := os.Create(modelPath + repositroy)
	if err != nil {
		panic(err)
	}
	defer repoFile.Close()
	if *hasModel {
		err := os.Mkdir(modelPath+dataModel, prem)
		if err != nil {
			panic(err)
		}
		createAndWrite(rep, modelPath+dataModel+"/"+lowerModelName+".go", model)
	}
	if *hasDataSource {
		err := os.Mkdir(modelPath+dataSources, prem)
		if err != nil {
			panic(err)
		}

		if *hasDB {
			err := os.Mkdir(modelPath+dataSources+"/"+vitessDS, prem)
			if err != nil {
				panic(err)
			}

			createAndWrite(rep, modelPath+dataSources+"/"+vitessDS+"/"+vittesFileName, vitess)

			createAndWrite(rep, modelPath+dataSources+"/"+DBDSFileName, DBDS)

			createAndWrite(rep, modelPath+dataModel+"/"+receiversFileName, receiver)
		}

		if *hasCache {
			err := os.Mkdir(modelPath+dataSources+"/"+redisDS, prem)
			if err != nil {
				panic(err)
			}

			createAndWrite(rep, modelPath+dataSources+"/"+redisDS+"/"+redsiFileName, redis)

			createAndWrite(rep, modelPath+dataSources+"/"+cahceDSFileName, cacheDS)
		}

		if *hasCache && *hasDB {
			_, err = repoFile.WriteString(rep.Replace(fllRepositoyTem))

		} else if *hasDB && !*hasCache {
			_, err = repoFile.WriteString(rep.Replace(repositoyTemWithDB))
		} else if !*hasDB && *hasCache {
			_, err = repoFile.WriteString(rep.Replace(repositoyTemWithCache))
		}
		if err != nil {
			panic(err)
		}
	} else {
		_, err = repoFile.WriteString(rep.Replace(normalRepositoyTem))
		if err != nil {
			panic(err)
		}

	}

}

func createAndWrite(replacer *strings.Replacer, path string, str string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(replacer.Replace(str))
	if err != nil {
		panic(err)
	}
}
