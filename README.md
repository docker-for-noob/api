# Api

Partie Back-end du projet Docker for Noobs.  
*Par Camille ARSAC, Rémi COUFOURIER, Florian LEROY, Guillaume MARCEL, Steven NATIVEL, Cédric PIERRE-AUGUSTE et Arthur POULLIN.*

Nous travaillons avec une architecture hexagonale en golang.


## **Pour travailler en local**

Installation des dépendances :
```
go mod vendor
```

Build le projet :
```
docker-compose build
```

Lancer le projet :

```
docker-compose up -d
```

Url pour accéder à l'API : http://localhost:8080/ <br>
Url pour accéder à mongo-express : http://localhost:8081/


### **Pour arrêter le docker du projet**

```
docker-compose down
```


## **Pour voir vos dockers qui tourne sur votre machine**

```
docker ps -a
``` 
ou
```
docker-compose ps
```
