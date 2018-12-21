<?php

$db = new PDO('mysql:dbname=euromillions;host=127.0.0.1', 'root', 'root');
var_dump($db);
$file = fopen('/Users/christophe/Projects/euromillions/euromillions_4.csv', 'r');
$csv = fgetcsv($file, 0, ';');
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
  "'.$row[2].'",
  "'.trim($row[1]).'"
  )';
 $db->exec($query);
}
