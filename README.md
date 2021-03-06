# tsuserver

This is a server to be used with the game Attorney Online.
It is meant to be an alternative to the many servers floating
around, hoping for better performance and more flexibility.

## Features

* Multiple areas per server
* Simple user management

## How to use

Compile and run.

[tsuserver documentation](https://docs.google.com/document/d/10RH8MRbdiyKEEAzIQPB4dXJcvhROwdWXmPfgsuk3whU/edit)

## User Commands

* /area (number) - Changes user to that area, if blank will list current areas.

* /getarea [id] - Lists characters in a given area, or only the current one if id is blank.

* /getareas - Lists characters across all areas.

* /bg (background)- Changes the background to one in the backgrounds list.

* /charselect - Brings up the character select screen (Shrinks client)
	* If logged in as a mod you can /charselect (target)

* /doc - Gives you the doc in use in the current area, heavily logged to discourage/detect bombings
	* /setdoc (url) - Sets the doc in the current area.

* /g (message) - Sends a global message to everyone in the server
	* /global - Toggles Global off/on
	
* /help - Links to this readme.

* /lockarea (password) - Locks the area with (password), area will unlock when the last person leaves.

* /unlockarea - Unlocks the current area

* /motd - prints the MOTD to chat
	* If logged in as a mod you can change the MOTD. /motd (message)
	
* /need (message) - Sends an advert to everyone in the server
	* /adverts - Toggles adverts off/on
	
* /pm (target) - Sends a PM to the specified character or OOC name.
	* Character name only PMs the target in your current area.
	
* /pos (position) - Moves user to the specified position in court.
	* (wit, def, pro, jud, hlp, hld)
	
* /randomchar - Changes you to a random free character

* /roll (number) - Rolls a dice between 1 and 6 or 1 and number, min is 2, max is 9999

* /coinflip - Flips a coin

* /status (status) - If blank displays current area status.
	* Statuses are: idle, buildingopen, buildingfull, casingopen, casingfull, recess

* /switch (character folder) - Changes the user to the specified character.
	* Character must be in characters list.	
	
* /polls - Lists currently available polls for voting.

* /pollinfo (poll name) - Lists the info for a poll.

* /vote (poll name) (yes/no) - Casts a vote on the given poll.

## Mod Commands

* /announce - Makes a server wide announcement

* /ban (target) [reason] - Disconnects the target and adds them to the banlist, must be IP

* /bglock - Toggles the background lock in the area, preventing users from using /bg

* /dj (target) - Enables a user to play music

* /undj (target) - Disables a user from playing music

* /play (song.mp3) - Plays a song if it isnt in the music list.

* /gm (message) - Same as global but gives you a [M]od tag. (global mod)

* /iplist - Prints a list of all clients sorted by IP

* /judgelog - Prints a list of the last 10 characters that changed HP bars or WT/CE in current area

* /kick (target) - Disconnects the target from the server, can be IP, Character name or OOC name.
	* Character name only kicks the target in your current area.
	
* /lm (message) - Sends a mod message only in the area you are in. (local mod)
	
* /login (password) - Logs client in as moderator.

* /lockable - Toggles the current area between lockable and unlockable for normal users.

* /mute (target) - Mutes the target, can be IP, Character name or OOC name.
	* Character name only mutes the target in your current area.
	
* /reloadbans - Reloads the banlist from its file.

* /reloadmusic - Reloads the musiclist.
	
* /unmute (target) - Unmutes the target, can be IP, Character name or OOC name.
	* Character name only mutes the target in your current area.
	
* /newpoll (name) (description) - Creates a new poll people can vote in.
	* First word is the poll name, after that is the description
	* /newpoll Adachi Should Adachi be removed? - New poll called "Adachi" with the description.

* /closepoll (name) - Closes a poll, preventing people from voting.

* /pollresults (name) - Prints results of a given poll.

* /reloadpolls - Reloads polls from their file.


## License

This server is licensed under the GPLv3 license. See the
[LICENSE](LICENSE.md) file for more information.
