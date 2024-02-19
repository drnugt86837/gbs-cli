package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateModuleStructure(moduleName string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	moduleDir := filepath.Join(currentDir, "module")

	if _, err := os.Stat(moduleDir); os.IsNotExist(err) {
		if err := os.MkdirAll(moduleDir, os.ModePerm); err != nil {
			return err
		}
	}

	moduleNameDir := filepath.Join(moduleDir, moduleName)

	if _, err := os.Stat(moduleNameDir); !os.IsNotExist(err) {
		return fmt.Errorf("Folder %s already exists, aborting.\n", moduleName)
	}

	if err := os.MkdirAll(moduleNameDir, os.ModePerm); err != nil {
		return err
	}
	fmt.Printf("Created folder: %s\n", moduleNameDir)

	controllerFileName := fmt.Sprintf("%sController.go", strings.Title(moduleName))
	controllerFilePath := filepath.Join(moduleNameDir, controllerFileName)
	controllerFileContent := fmt.Sprintf(`package %s

import (
	"github.com/gin-gonic/gin"
)

type %sController struct {
	%sService *%sService
}

func New%sController(service *%sService) *%sController {
	return &%sController{service}
}

func (c *%sController) GetAll(g *gin.Context) {
	result, err := c.%sService.GetAll()
	if err != nil {
		c.HandleError(g, err)
	} else {
		c.SuccessRes(g, result)
	}
}

func (c *%sController) Create(g *gin.Context) {
	var dto dto.Create%sDto
	if err := g.Bind(&dto); err != nil {
		c.HandleError(g, err)
		return
	}
	err := c.%sService.Create(&dto)
	if err != nil {
		c.HandleError(g, err)
	} else {
		c.SuccessRes(g, nil)
	}
}`,
		moduleName,
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
	)
	if err := ioutil.WriteFile(controllerFilePath, []byte(controllerFileContent), 0644); err != nil {
		return err
	}
	fmt.Printf("Created %s file: %s\n", controllerFileName, controllerFilePath)

	serviceFileName := fmt.Sprintf("%sService.go", strings.Title(moduleName))
	serviceFilePath := filepath.Join(moduleNameDir, serviceFileName)
	serviceFileContent := fmt.Sprintf(`package %s

import (
	// Add your service dependencies here
)

type %sService struct {
	// Add your service dependencies here
}

func New%sService() *%sService {
	return &%sService{}
}

func (s *%sService) GetAll() ([]YourModel, error) {
	// Your service logic for GetAll method here
}

func (s *%sService) Create(dto *dto.Create%sDto) error {
	// Your service logic for Create method here
	return nil
}`,
		moduleName,
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
	)
	if err := ioutil.WriteFile(serviceFilePath, []byte(serviceFileContent), 0644); err != nil {
		return err
	}
	fmt.Printf("Created %s file: %s\n", serviceFileName, serviceFilePath)

	routesFileName := fmt.Sprintf("%sRoutes.go", strings.Title(moduleName))
	routesFilePath := filepath.Join(moduleNameDir, routesFileName)
	routesFileContent := fmt.Sprintf(`package %s

import (
	"github.com/gin-gonic/gin"
)

func SetRoute(g *gin.Engine, baseGroup string) {
	%sGroup := g.Group(baseGroup + "/%s")
	controller := Init%sController()
	%sGroup.GET("", controller.GetAll)
	%sGroup.POST("", controller.Create)
}`,
		moduleName,
		moduleName,
		moduleName,
		strings.Title(moduleName),
		moduleName,
		moduleName,
	)
	if err := ioutil.WriteFile(routesFilePath, []byte(routesFileContent), 0644); err != nil {
		return err
	}
	fmt.Printf("Created %s file: %s\n", routesFileName, routesFilePath)

	dtoDir := filepath.Join(moduleNameDir, "dto")

	if _, err := os.Stat(dtoDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dtoDir, os.ModePerm); err != nil {
			return err
		}
	}

	createDtoFileName := fmt.Sprintf("Create%sDto.go", strings.Title(moduleName))
	createDtoFilePath := filepath.Join(dtoDir, createDtoFileName)
	createDtoFileContent := fmt.Sprintf(`
package dto

import (
)

type Create%sDto struct {
	// dto parameters
}`, strings.Title(moduleName))
	if err := ioutil.WriteFile(createDtoFilePath, []byte(createDtoFileContent), 0644); err != nil {
		return err
	}
	fmt.Printf("Created %s file: %s\n", createDtoFilePath, createDtoFilePath)

	wireFilePath := filepath.Join(moduleNameDir, "wire.go")
	wireFileContent := fmt.Sprintf(`//go:build wireinject
// +build wireinject

package %s

import (
	"github.com/google/wire"
)

func Init%sController() *%sController {
	wire.Build(
		New%sController,
		New%sService,
	)
	return &%sController{}
}`,
		moduleName,
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
		strings.Title(moduleName),
	)
	if err := ioutil.WriteFile(wireFilePath, []byte(wireFileContent), 0644); err != nil {
		return err
	}
	fmt.Printf("Created wire.go file: %s\n", wireFilePath)

	wireGenFilePath := filepath.Join(moduleNameDir, "wire_gen.go")
	wireGenFileContent := fmt.Sprintf(`// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package %s

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func Init%sController() *%sController {
	%sService := New%sService()
	%sController := New%sController(%sService)
	return %sController
}

// wire.go:`,
		moduleName,
		strings.Title(moduleName),
		strings.Title(moduleName),
		moduleName,
		strings.Title(moduleName),
		moduleName,
		strings.Title(moduleName),
		moduleName,
		moduleName,
	)
	if err := ioutil.WriteFile(wireGenFilePath, []byte(wireGenFileContent), 0644); err != nil {
		return err
	}
	fmt.Printf("Created wire.go file: %s\n", wireGenFilePath)

	return nil
}
