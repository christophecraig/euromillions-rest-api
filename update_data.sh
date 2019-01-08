wget -q https://www.fdj.fr/generated/game/euromillions/euromillions_4.zip -O ./euromillions.zip;
echo "downloaded from www.fdj.fr"
unzip -o ./euromillions.zip;
echo "extracted content to euromillions_4.csv"
/usr/bin/php $PWD/php-import/import_euromillions.php;
echo "imported csv into table results\r";
echo "building main.go\r";
go build;
echo "built. now serving on :8000";
./euromillions-rest-api;