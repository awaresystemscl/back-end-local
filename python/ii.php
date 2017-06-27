<?php 

$argumento = 'El argumento mas feo del mundo';
$command = escapeshellcmd('/home/server/python/mail.py'.' "'.$argumento.'"');
$output = shell_exec($command);
echo $output;

?>