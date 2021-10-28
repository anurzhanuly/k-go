<?php
$stdout = fopen('php://stdout', 'w');
fwrite($stdout, "программа на PHP запущена\n");

$name = $_GET['name'];
// пишем клиенту
echo "Салем еще раз, {$_GET['name']}!\n";

// пишем в STDOUT
fwrite($stdout, "запрос обработан\n");