<?php

$db = new PDO('mysql:dbname=euromillions;host=127.0.0.1', 'root', 'root');
$file = fopen('./euromillions_4.csv', 'r');
$csv = fgetcsv($file, 0, ';');
$db->exec('truncate table results');
echo "truncated table results\n";
while ($row = fgetcsv($file, 0, ';')) {
  $query = 'insert into results (b1, b2, b3, b4, b5, e1, e2, myMillion, date, weekday) values ('.
    $row[5].',
  '.$row[6].',
  '.$row[7].',
  '.$row[8].',
  '.$row[9].',
  '.$row[10].',
  '.$row[11].',
  "'.$row[73].'",
  STR_TO_DATE("'.$row[2].'", "%d/%m/%Y"),
  "'.trim($row[1]).'"
  )';
  // var_dump($query);
 $db->exec($query);
}
