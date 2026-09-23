# projet-red-cyberpunk

# 🌃 NEON CITY

## Projet RED — Ynov 2026/2027

### 🎮 Présentation

**Année 2087.**

Neon City est une mégalopole contrôlée par de puissantes corporations. La ville est divisée entre des quartiers riches ultra-technologiques et des zones pauvres contrôlées par des gangs.

Le joueur incarne un **mercenaire cybernétique** qui accepte des contrats pour gagner des **crédits**, améliorer ses **implants** et devenir progressivement plus puissant.

**Neon City** est un RPG développé en **Go**, jouable directement dans le terminal.

### 🧬 Classes

Le joueur peut choisir entre trois classes :

* **Netrunner** — spécialisé dans le hacking et les compétences technologiques.
* **Mercenaire** — équilibré et spécialisé dans le combat.
* **Cyborg** — possède beaucoup de PV et est spécialisé dans les améliorations cybernétiques.

### ⚔️ Fonctionnalités

* Création et personnalisation du personnage
* Exploration de Neon City
* Rencontres aléatoires
* Combats au tour par tour
* Compétences cybernétiques
* Gestion de l'énergie
* XP et progression du niveau
* Inventaire et équipements cybernétiques
* Marchand et système de crédits
* Atelier cybernétique et fabrication
* Missions et récompenses

## 👥 Équipe

* **Alexis DUMAIRE**
* **Selwanss Miar**
* **Jawad IBRAHIM**

## 🛠️ Installation

### Prérequis

* **Go**
* **Git**

Vérifier que Go est installé :

```bash
go version
```

### Cloner le projet

```bash
git clone <URL_DU_DEPOT>
cd NEON-CITY
```

## ▶️ Lancement

Depuis la racine du projet :

```bash
go run ./src
```

Le jeu se lance directement dans le terminal.

## 🧪 Tests et vérification

Avant de rendre le projet, nous vérifions les principales fonctionnalités du jeu directement en l'exécutant.

Les éléments testés comprennent notamment :

* Création d'un personnage et choix de sa classe
* Affichage des statistiques du personnage
* Gestion de l'inventaire
* Utilisation des objets
* Achat et vente chez le marchand
* Fabrication d'équipements chez le forgeron
* Équipement et remplacement des objets équipés
* Rencontres avec les ennemis
* Système de combat
* Utilisation des compétences et gestion de l'énergie
* Gains d'XP et montée de niveau
* Drops obtenus après les combats
* Missions et récompenses
* Mort et réapparition du personnage

Pour vérifier que le projet se compile et se lance correctement :

```bash
go run ./src
```

Les fonctionnalités sont ensuite testées directement dans le jeu en utilisant les différents menus et systèmes disponibles.

**Bienvenue dans Neon City. 🌃**
