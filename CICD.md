# CICD sur kube

Il y a énormement de façon de faire mais je vais selectionner une façon de faire. Cette méthode est pour une application qui n'a pas pour but d'être envoyé/distribué à plusieurs personnes,agences,.. c'est à dire que je ne vais pas parler des charts helms.


## Mise en place d'un répertoire github/gitlab

Deux branches main et develop.

Je commencerai par mettre mon code source sur github et gitlab accompagné d'un dockerfile.

## CICD via github/gitlab

Je ferai une CICD avec plusieurs étapes :

### Branche main :

- Build de l'application
- Lancement des tests
- Push sur docker.io l'image grâce au Dockerfile
- Deploy sur kubectl à l'aide de patch
- Tests hurls pour vérifier le bon fonctionnement de l'application 
- Rollback (manuel)

### Branche develop :

- Build de l'application
- Lancement des tests
- Push sur docker.io l'image grâce au Dockerfile (avec tag dev)
- Deploy sur kubectl sur un environnement de test à l'aide de patch
- Tests hurls pour vérifier le bon fonctionnement de l'application 
- Rollback (manuel)

## Kube

Mise en place des manifestes et ensuite l'application sera déployé de façon automatique
