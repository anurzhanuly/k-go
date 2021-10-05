<?php
$stdout = fopen('php://stdout', 'w');
fwrite($stdout, "программа на PHP запущена\n");

// пишем клиенту
echo "Привет, {$_GET['name']}!\n";

// пишем в STDOUT
fwrite($stdout, "запрос обработан\n");