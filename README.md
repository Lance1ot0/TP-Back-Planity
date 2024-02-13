## Guide d'utilisation

Ce guide fournit des instructions pour mettre en place et exécuter le projet à l'aide de Docker Compose pour la base de données, ainsi que les étapes nécessaires pour lancer le serveur et le frontend.

### Installation

1. Assurez-vous d'avoir Docker et Docker Compose installés sur votre système.

### Mise en place du projet

1. Clonez ce dépôt sur votre machine locale :
    ```bash
    git clone https://github.com/Lance1ot0/TP-Back-Planity.git
    ```

### Docker Compose pour la base de données

1. À la racine du projet, exécutez la commande suivante pour lancer la base de données à l'aide de Docker Compose :
    ```bash
    docker-compose up -d
    ```

### Lancement du serveur

1. Accédez au répertoire backend du projet :
    ```bash
    cd backend
    ```

2. Exécutez la commande suivante pour lancer le serveur :
    ```bash
    go run main/main.go
    ```

### Lancement du frontend

1. Accédez au répertoire frontend du projet :
    ```bash
    cd frontend
    ```

2. Installez les dépendances en exécutant :
    ```bash
    npm install
    ```

3. Lancez le serveur de développement avec la commande :
    ```bash
    npm run dev
    ```

### Accès au projet

Une fois ces étapes effectuées, vous pouvez accéder à l'application à l'adresse [http://localhost:3000](http://localhost:3000) dans votre navigateur web.

---
**Remarque :** Assurez-vous que les ports nécessaires ne sont pas déjà utilisés par d'autres services sur votre machine.
