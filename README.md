Roaster
=======
This is the source code of Roaster, a tool to edit the EPROMs of NBA Jam Tournament Edition rev 4.0. Currently, Roaster can patch images (with palettes) and ASCII text.

Usage
=====
1/ Extract the EPROMs content from your NBA Jam Tournament Edition rev 4.0 board. Place them in the nbajamte directory.
2/ When naming the ROMs, the names must be EXACTLY as follow. Use the NBA Jam kit documentation that came with the board to identify EPROMs:

* l1_nba_jam_tournament_game_rom_ug14.ug14
* l1_nba_jam_tournament_game_rom_ug16.ug16
* l1_nba_jam_tournament_game_rom_ug17.ug17
* l1_nba_jam_tournament_game_rom_ug18.ug18
* l1_nba_jam_tournament_game_rom_ug19.ug19
* l1_nba_jam_tournament_game_rom_ug20.ug20
* l1_nba_jam_tournament_game_rom_ug22.ug22
* l1_nba_jam_tournament_game_rom_ug23.ug23
* l1_nba_jam_tournament_game_rom_uj14.uj14
* l1_nba_jam_tournament_game_rom_uj16.uj16
* l1_nba_jam_tournament_game_rom_uj17.uj17
* l1_nba_jam_tournament_game_rom_uj18.uj18
* l1_nba_jam_tournament_game_rom_uj19.uj19
* l1_nba_jam_tournament_game_rom_uj20.uj20
* l1_nba_jam_tournament_game_rom_uj22.uj22
* l1_nba_jam_tournament_game_rom_uj23.uj23
* l1_nba_jam_tournament_u3_sound_rom.u3
* l1_nba_jam_tournament_u12_sound_rom.u12
* l1_nba_jam_tournament_u13_sound_rom.u13
* l4_nba_jam_tournament_game_rom_ug12.ug12
* l4_nba_jam_tournament_game_rom_uj12.uj12

3/ Adjust the patching instructions in feed.txt. This is where you can set the name of the player, player photo. The photos must be PNG 256 indexed color format. Additionaly, each player slot as a specific width/height so make sure to respect that.
4/ Run: go run src/main/main.go 
5/ You should see the following output:
Deinterlacing gfxrom
Deinterlacing mainrom   
6/ Burn back the EPROMs found in the output directory.
7/ Enjoy!

Disclaimer
==========
This is not an officially supported Google product
