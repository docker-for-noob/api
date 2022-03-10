# api

**Installation des dépendances**
```
go mod vendor
```

# Architecture
## cmd
Le dossier contient les points d'entrées de l'application, par exemple le router


## Internal
### core
Le dossier **core** contient le code métier de l'application 
 
Les fichier sont organisés de la manière suivante: 

#### Domain 

Ce dossier contient tout les models 
il contient des *struct* qui définisent toutes les entité qui seront utilié par le code métier

#### Ports

Ce dossier contient les iterface qui vont servir a communiquer avec les acteurs

il existe deux type de port :

> Les service, qui sont utilisé pour communiquer avec les inputs (Intéractions humaine)

> Les Repository, qui servent a communiquer avec les acteur de sortie (DB, cache, Queue)

#### Services

Ce sont les points d'entrée du "Core", chaque service doit implementer le "Port" corespondant.

### Handlers

Les handlers transforme les inputs en appelle a un service

### Repository

Les repository envois les models traité par un service du "Core" vers un module extérieur comme par exemple une base de données

[Tuto](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)

# Swagger

Installer swag  
Pour regénérer la doc de swagger éxécuter : `swag init -g "cmd/httpserver/main.go"`  
L'url pour voir la documentation API : http://localhost:8080/swagger/index.html