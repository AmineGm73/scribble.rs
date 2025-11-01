package translations

func initFrenchTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "Ce site web nécessite JavaScript pour fonctionner correctement.")

	translation.put("start-the-game", "Prêt !")
	translation.put("force-start", "Forcer le Démarrage")
	translation.put("force-restart", "Forcer le Redémarrage")
	translation.put("game-not-started-title", "La partie n'a pas démarré")
	translation.put("waiting-for-host-to-start", "Veuillez attendre que l'hôte du salon démarre la partie.")

	translation.put("now-spectating-title", "Vous êtes maintenant spectateur")
	translation.put("now-spectating-text", "Vous pouvez quitter le mode spectateur en appuyant sur le bouton œil en haut.")
	translation.put("now-participating-title", "Vous participez maintenant")
	translation.put("now-participating-text", "Vous pouvez passer en mode spectateur en appuyant sur le bouton œil en haut.")

	translation.put("spectation-requested-title", "Mode spectateur demandé")
	translation.put("spectation-requested-text", "Vous serez spectateur après ce tour.")
	translation.put("participation-requested-title", "Participation demandée")
	translation.put("participation-requested-text", "Vous participerez après ce tour.")

	translation.put("spectation-request-cancelled-title", "Demande de mode spectateur annulée")
	translation.put("spectation-request-cancelled-text", "Votre demande de mode spectateur a été annulée, vous continuerez de participer.")
	translation.put("participation-request-cancelled-title", "Demande de participation annulée")
	translation.put("participation-request-cancelled-text", "Votre demande de participation a été annulée, vous continuerez d'être spectateur.")

	translation.put("round", "Manche")
	translation.put("toggle-soundeffects", "Activer/Désactiver les effets sonores")
	translation.put("toggle-pen-pressure", "Activer/Désactiver la pression du stylet")
	translation.put("change-your-name", "Pseudo")
	translation.put("randomize", "Aléatoire")
	translation.put("apply", "Appliquer")
	translation.put("save", "Sauvegarder")
	translation.put("toggle-fullscreen", "Activer/Désactiver le plein écran")
	translation.put("toggle-spectate", "Activer/Désactiver le mode spectateur")
	translation.put("show-help", "Afficher l'aide")
	translation.put("votekick-a-player", "Voter pour exclure un joueur")

	translation.put("last-turn", "(Dernier tour : %s)")

	translation.put("drawer-kicked", "Puisque le joueur exclu était en train de dessiner, aucun de vous n'obtiendra de points pour cette manche.")
	translation.put("self-kicked", "Vous avez été exclu")
	translation.put("kick-vote", "(%s/%s) joueurs ont voté pour exclure %s.")
	translation.put("player-kicked", "Le joueur a été exclu.")
	translation.put("owner-change", "%s est le nouveau propriétaire du salon.")

	translation.put("change-lobby-settings-tooltip", "Changer les paramètres du salon")
	translation.put("change-lobby-settings-title", "Paramètres du salon")
	translation.put("lobby-settings-changed", "Paramètres du salon modifiés")
	translation.put("advanced-settings", "Paramètres avancés")
	translation.put("chill", "Détendu")
	translation.put("competitive", "Compétitif")
	translation.put("chill-alt", "Bien qu'être rapide soit récompensé, ce n'est pas trop grave si vous êtes un peu plus lent.\nLe score de base est assez élevé, concentrez-vous sur le plaisir !")
	translation.put("competitive-alt", "Plus vous êtes rapide, plus vous obtiendrez de points.\nLe score de base est beaucoup plus bas et le déclin est plus rapide.")
	translation.put("score-calculation", "Score")
	translation.put("word-language", "Langue")
	translation.put("drawing-time-setting", "Temps de Dessin")
	translation.put("rounds-setting", "Manches")
	translation.put("max-players-setting", "Nombre Maximum de Joueurs")
	translation.put("public-lobby-setting", "Salon Public")
	translation.put("custom-words", "Mots Personnalisés")
	translation.put("custom-words-info", "Entrez vos mots supplémentaires, en les séparant par des virgules")
	translation.put("custom-words-per-turn-setting", "Mots Personnalisés Par Tour")
	translation.put("players-per-ip-limit-setting", "Limite de Joueurs Par IP")
	translation.put("save-settings", "Sauvegarder les paramètres")
	translation.put("input-contains-invalid-data", "Votre saisie contient des données invalides :")
	translation.put("please-fix-invalid-input", "Corrigez la saisie invalide et réessayez.")
	translation.put("create-lobby", "Créer un Salon")
	translation.put("create-public-lobby", "Créer un Salon Public")
	translation.put("create-private-lobby", "Créer un Salon Privé")

	translation.put("refresh", "Actualiser")
	translation.put("join-lobby", "Rejoindre le Salon")

	translation.put("message-input-placeholder", "Tapez vos suppositions et messages ici")

	translation.put("word-choice-warning", "Mot si vous ne choisissez pas à temps")
	translation.put("choose-a-word", "Choisissez un mot")
	translation.put("waiting-for-word-selection", "En attente de la sélection du mot")
	// This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "est en train de choisir un mot.")

	translation.put("close-guess", "'%s' est très proche.")
	translation.put("correct-guess", "Vous avez correctement deviné le mot.")
	translation.put("correct-guess-other-player", "'%s' a correctement deviné le mot.")
	translation.put("round-over", "Fin du tour, aucun mot n'a été choisi.")
	translation.put("round-over-no-word", "Fin du tour, le mot était '%s'.")
	translation.put("game-over-win", "Félicitations, vous avez gagné !")
	translation.put("game-over-tie", "C'est une égalité !")
	translation.put("game-over", "Vous avez terminé %s. avec %s points")

	translation.put("change-active-color", "Changer votre couleur active")
	translation.put("use-pencil", "Utiliser le crayon")
	translation.put("use-eraser", "Utiliser la gomme")
	translation.put("use-fill-bucket", "Utiliser le pot de peinture (Remplit la zone cible avec la couleur sélectionnée)")
	translation.put("change-pencil-size-to", "Changer la taille du crayon / de la gomme à %s")
	translation.put("clear-canvas", "Effacer le canevas")
	translation.put("undo", "Annuler la dernière modification que vous avez faite (Ne fonctionne pas après \""+translation.Get("clear-canvas")+"\")")

	translation.put("connection-lost", "Connexion perdue !")
	translation.put("connection-lost-text", "Tentative de reconnexion"+
		" ...\n\nAssurez-vous que votre connexion internet fonctionne.\nSi le "+
		"problème persiste, contactez le webmaster.")
	translation.put("error-connecting", "Erreur de connexion au serveur")
	translation.put("error-connecting-text",
		"Scribble.rs n'a pas pu établir de connexion socket.\n\nBien que votre connexion "+
		"internet semble fonctionner, soit le\nserveur, soit votre pare-feu n'a pas "+
		"été configuré correctement.\n\nPour réessayer, rechargez la page.")

	translation.put("message-too-long", "Votre message est trop long.")

	// Help dialog
	translation.put("controls", "Commandes")
	translation.put("pencil", "Crayon")
	translation.put("eraser", "Gomme")
	translation.put("fill-bucket", "Pot de peinture")
	translation.put("switch-tools-intro", "Vous pouvez changer d'outil en utilisant des raccourcis")
	translation.put("switch-pencil-sizes", "Vous pouvez aussi changer la taille du crayon en utilisant les touches %s à %s.")

	// Generic words
	// "close" as in "closing the window"
	translation.put("close", "Fermer")
	translation.put("no", "Non")
	translation.put("yes", "Oui")
	translation.put("system", "Système")

	translation.put("source-code", "Code Source")
	translation.put("help", "Aide")
	translation.put("submit-feedback", "Commentaires")
	translation.put("stats", "Statut")

	RegisterTranslation("fr", translation)

	return translation
}