# **Conseil**
- Utilise [VSCode](https://vscode.com) avec des extensions, il est très puissant
- Créer une conversation (ou plusieurs) spécialisé avec [ChatGPT](https://chatgpt.com) et donne lui tes messages d'erreur
- Utilise un kit d'extensions spécial python, il y en a plein
- N'hésite pas à passer plus de temps que nécessaire sur une étape ça ne sera jamais une perte de temps
- Amuse toi, sinon tu abandonneras trop vite

# Première étape

Apprendre et maitriser les bases de python.
Savoir :
- Boucles `for` et `while`
- Programmation orienté objet (`class` et `def`)
- Import et utilisation de librairies (avec `import`)

> **Aide**
> Faire un projet utilisant `random` pour simuler un lancer de X dé de Y faces, avec question avant pour savoir le nombre de dés et le nombre de faces. Donné ensuite chaque résultat et la somme totale

# Deuxième étape

Trouver une librairie supportant les `.dicom` et savoir utiliser les bases
Objectifs:
- Trouver une librairie pour manipuler les `.dicom`
- Savoir ouvrir un `.dicom` avec

> **Aide**
> Fais ta recherche sur [Github](https://github.com) en spécifiant le tag `lang:python`

# Troisième étape

Savoir utiliser **toutes** les fonctionnalités de la librairie trouvée à l'étape 2.
Savoir :
- Lire un `.dicom`
- Feuilleter les différents images (bouger la coupe)
- Enregistrer une image seule

> **Aide**
> Utilise VSCode avec Intellisense pour connaitre les différentes fonctions de la librairie et joue avec, n'hésite pas à demander à des IA de t'expliquer les erreurs

# Quatrième étape

Créer un environnement virtuel et installe les librairies dont tu auras besoin.
Objectifs :
- Savoir créer et lancer un environnemnt virtuel (venv)
- Installer les librairies

> **Aide**
> Explore les différents outils de venv et teste en plusieurs avant de trouver le meilleur pour toi (ou utilise le premier fais comme tu veux). Les librairies qui te seront les plus utiles seront : la librairie de manipulation de `.dicom` et `streamlit`

# Cinquième étape

Utilise `streamlit` pour créer une page simple
Objectif :
- Savoir faire et lancer un serveur web avec `streamlit`
- Créer une page basique (un hello world pour commencer)
- Savoir faire une requete au code python et avoir une réponse

> **Aide**
> Commence avec une page qui dit bonjour puis ajoute un bouton sur lequel lorsque tu cliques un message s'affiche dans la console et un texte apparait

# Sixième étape

Définis ton interface et tes fonctionnalités
Objectifs :
- Savoir où mettre quoi (bouton de submission à droite ? en bas ? Un ou plusieurs boutons ?)
- Permettre l'upload d'un fichier `.dicom`
- Afficher le `.dicom` selectionné

> **Aide**
> Teste toutes les fonctionnalités possibles et choisis ce que tu penses être le mieux, n'hésite pas à revenir sur tes choix pour les changer

# Septième étape

Affiner l'interface et ses fonctions
Objectifs:
- Envoyer le `.dicom` dans un dossier connu du serveur (ex: `upload`)
- Permettre de scroller dans le `.dicom`
- Lire le `.dicom` coté serveur (dire le nombre d'image par exemple)
- Permettre d'envoyer ce que tu penses que le patient a sur le `.dicom`

> **Aide**
> Utilise la librairie de l'étape deux et les fonctions de bases de python pour la manipulation de fichier (et soit créatif)

# Huitième étape (Projet dans le projet)

Création d'un modèle NLP
1. Se renseigner
    - Vérifie si des modèles n'existe déjà pas que tu pourrais réutiliser
    - Fais des tests avec des phrases simple
2. Mettre en application
    - Fais un fichier avec tous les titres de Radiopedia
    - Utilise le modèle NLP choisi pour qu'il sache quel titre est le plus proche du constat du médecin
    - Récupère le lien associé au titre

> **Aide**
> Fais un autre dossier puis explore [Github](https://github.com), Google est ton ami

# Neuvième étape

Utilsie Beautiful Soup (bs4) pour récupérer les données de la page radiopedia
Objectf :
- Se connecter à la page
- Récupérer son code source
- Affiner le code (récupérer que le texte, que les images, etc)
- Télécharger les images

> **Aide**
> Sur [Youtube](https://youtube.Com) il existe plein de tuto de web scraping, n'hésite pas à en regarder, et renseigne toi sur la légalité de ce que tu fais

# Dixième étape (Projet dans le projet)

Utilise `PIL` pour comparer les images
1. Maitrise `PIL`
    - Télécharge `PIL`
    - Fais joujou avec (créer des images random, rétrécis en d'autres, etc)
2. Mise en application
    - Utilise la librairie de l'étape 2 pour extraire chaque images du `.dicom` fourni et de radiopedia
    - Compare les images (différence de couleurs des pixels)
    - Fais une moyenne des différences
    - Affine ton algo pour détecter les différences

> **Aide**
> Pour le coup j'ai jamais fais donc c'est ta merde, [ChatGPT](https://chatgpt.com) est ton ami

# Onzième étape

Affichage du résultat
Objectif :
- Avoir un pourcentage de différence entre les images
- Retourner le pourcentage sur le site avec `streamlit`

> **Aide**
> Réutilise tes connaissances de l'étape 5 et part du principe que + de diff égal - de chances pour que ça corresponde
