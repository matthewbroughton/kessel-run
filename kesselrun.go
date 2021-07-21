{{$message := printf ""}}
{{$success := randInt 100}}
{{$reasonList := cslice "You crashed into" "You were captured by" "You were destroyed by" "You were scared off by" "You fell into" "You were shot down by" "You were distracted by" "You landed on"}}
{{$reason := ""}}
{{$obstacleList := cslice "the Emperor." "the Death Star." "the Slave 1." "a group of Zygerrian slavers." "a black hole." "a supernova." "the Sarlacc Pit." "Production Intern 1." "a seismic charge." "a piece of Alderaan." "the bigger fish." "Porgs."}}
{{$obstacle := ""}}
{{$failureList := cslice "Your crew mutinied." "You didn't pay your port fees." "Your hyperdrive failed." "You forgot to bypass the compressor." "You were eaten by a space worm." "You drank a bad glass of blue milk and got sick." "You left the handbrake on." "The Rathtar escaped." "You were stuck in the belly of a hyperspace whale." "You failed to deliver to Kanji club." "You fell out of hyperspace." "You were caught up to by Jabba's bounty hunters." "You were fired by Kathleen Kennedy." "Chewy forgot to punch it." "You went up against a Sicilian when death was on the line." "You simply walked into Mordor." "The Eagles came." "You woke the Balrog." "Smaug is fire ... Smaug is death." "Thrawn outplayed you." "You lost sabacc to Lando." "That was a moon. And you hit it." "You were frozen in carbonite." "You didn't use the Force, Luke." "You were on the wrong run." "You tried spinning. That's a good trick." "You were torn apart by buzz droids." "This isn't podracing." "You lost your lightsaber again. Obi-Wan is going to kill you." "You went where the light didn’t touch." "It was a trap." "The possibility of successfully navigating an asteroid field is approximately 3720:1!" "The Death Star II was fully operational." "You were told the odds." "Somehow, Palpatine returned." "Jar Jar was your copilot." "You were outpiloted by Darth Vader." "You were outpiloted by Han Solo." "Loki stabbed you in the back." "You didn't survive the Blip." "HK-47 called you a meatbag." "You died. But family never dies."}}
{{$distance := (toString (randInt 5 50))}}

{{if le $success 40}}
    {{$message = printf "You made the Kessel Run in %s parsecs!" $distance}}
{{else if le $success 70}}
    {{$reason = index $reasonList (randInt (len $reasonList))}}
    {{$obstacle = index $obstacleList (randInt (len $obstacleList))}}
    {{$message = printf "%s %s" $reason $obstacle}}
{{else}}
    {{$message = index $failureList (randInt (len $failureList))}}
{{end}}

{{sendMessage nil $message }}