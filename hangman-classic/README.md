# Hangman
## Descriptif du projet
Le programme est un jeu de pendu.
Le programme doit sélectionner un mot aléatoirement dans un fichier et doit demander à l'utilisateur de le trouver en comparant la lettre ou le mot donné par l'utilisateur.
Pour cela, le programme affiche une série d'underscore pour lui donner un indice.
Si le joueur trouve le mot correspondant, le programme lui dit qu'il a gagné et s'arrête.
Si la lettre donnée par le joueur est contenue dans le mot, le programme affiche la lettre dans le mot à la place de l'underscore correspondant à cette dernière.  
Si la lettre n'est pas dans le mot, le joueur perd une vie et un affichage du pendu évolutif en fonction du nombre de vie apparaît. Si le nombre de vies arrive à 0, alors le programme dit au joueur qu'il a perdu et il s'arrête. 
## Fonctionnalités disponibles
- Le programme récupère un mot au hasard dans un fichier dont le nom est donné en argument. 
- Le programme affiche au joueur un tableau d'underscore de la longueur du mot avec un nombre de lettres qui correspond au nombre de lettres du mot divisé par 2 et moins 1. Ces lettres sont aléatoires. 
- le programme va demander à l'utilisateur de donner soit une lettre soit le mot en entier. 
- Le programme va comparer la lettre ou le mot donné par le joueur au mot qu'il doit trouver. Si la lettre ou le mot ne correspondent pas au mot à trouver, alors l'utilisateur perd une vie. Sinon, si la lettre est dans le mot alors le programme l'affiche dans le tableau d'underscore. Si le mot ou le tableau d'underscore correspondent au mot à trouver, ou si le nombre de vies est égal à 0, ou si le joueur entre "exitgame", alors le programme s'arrête.
- on a aussi ajouté une fonction qui sert à afficher les lettres qui ont été jouées précedemment.
- Nous avons aussi quelques fonctionnalités cachées dans notre programme comme la fonction showword qui consiste à  donner le mot à trouver au joueur ou la fonction godmod qui permet de mettre un nombre de vies infini. 
## Installation du projet
- Pour installer le projet, il suffit d'aller sur le répo sur gitea et cliquer sur download_zip et d'extraire les fichiers. Ensuite, il faut aller dans Visual Studio Code, cliquer sur file et open folder et sélectionner le dossier.
- Ou alors nous pouvons ouvrir directement Visual Studio Code et entrer la commande "git clone [url du repo]".
## Lancement du projet
Pour lancer le projet, il suffit de se placer dans le dossier "hangman-classic" dans le terminal et d'entrer la commande "go run main.go [nom du fichier contenant les mots à trouver]". 
## Exemples
### lancement du jeu
 ```
 $ go run main.go words.txt
Votre ami José a été condamné à mort par pendaison, mais c'est votre unique fournisseur de saucisson donc vous voulez le sauver. Pour cela, vous devez retrouver le mot qui pourra l'inocenter. Vous devez donc compléter le mot caché avec les lettres manquantes.
[_ h _ i _ _ ]
essayez de trouver une lettre ou le mot en entier !
Entrez exitgame si vous voulez quitter le jeu
```
### une lettre juste donnée 
```
e
Bravo, cette lettre est dans le mot !
Il vous reste 10 vies.






=========

Vous avez joué les lettres suivantes :  e
[_ h _ i _ e ]
essayez de trouver une lettre ou le mot en entier !
Entrez exitgame si vous voulez quitter le jeu
```
### Si une lettre incorrecte est donnée
```
z
Dommage, cette lettre n'est pas dans le mot !
Il vous reste 9 vies.

      |
      |
      |
      |
      |
=========

Vous avez joué les lettres suivantes :  e z
[_ h _ i _ e ]
essayez de trouver une lettre ou le mot en entier !
Entrez exitgame si vous voulez quitter le jeu
```
### Si un mot est donné mais qu'il ne correspond pas à la longueur du mot à trouver
```
ddd
vous ne pouvez essayer de deviner qu'une lettre a la fois ou tout le mot d'un seul coup
Dommage, cette lettre n'est pas dans le mot !
Il vous reste 8 vies.
  +---+
      |
      |
      |
      |
      |
=========

Vous avez joué les lettres suivantes :  e z ddd
[_ h _ i _ e ]
essayez de trouver une lettre ou le mot en entier !
Entrez exitgame si vous voulez quitter le jeu
```
### Si le mot est trouvé lettre par lettre
```
c
Bravo, cette lettre est dans le mot !
Vous avez trouvés toutes les lettres du mot !
Bravo vous avez gagné !
José est sauvé et le saucisson aussi !
```
### Si le mot est trouvé d'un seul coup
```
etude
Bravo, vous avez gagné !
José est sauvé et le saucisson aussi !
```
### Si la commande "exitgame" est entrée
```
exitgame
Fin de partie, Merci d'avoir joué !
```
### Si le nombre de vies arrive a 0
```
z
Dommage, cette lettre n'est pas dans le mot !
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========

Vous n'avez plus de vies !
Game Over.
```
### fonctions bonus
#### showword
```
showword
Hé tu n'était pas censé faire ça !
Bon d'accord voici le mot : recreation
Vous avez joué les lettres suivantes :  showword
[_ _ c _ _ a _ i _ n ]
essayez de trouver une lettre ou le mot en entier !
Entrez exitgame si vous voulez quitter le jeu
```
#### godmode
```
godmode
godmode activé, vous avez  99999999999 vies
vous ne pouvez essayer de deviner qu'une lettre a la fois ou tout le mot d'un seul coup
Dommage, cette lettre n'est pas dans le mot !






=========

Il vous reste 99999999998 vies.
Vous avez joué les lettres suivantes :  showword godmode
[_ _ c _ _ a _ i _ n _]
essayez de trouver une lettre ou le mot en entier !
Entrez exitgame si vous voulez quitter le jeu
```
