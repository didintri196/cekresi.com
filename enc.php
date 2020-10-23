<?php
// Script example.php
$options = getopt("f:");
// var_dump($options);
// $string = "JP9895941501";
$string = $options['f'];
$aKey = "79540e250fdb16afac03e19c46dbdeb3";
$sIv = "eb2bb9425e81ffa942522e4414e95bd0";
// encryption
// echo '<br>* * * encryption * * *' . PHP_EOL;
$ciphertext = openssl_encrypt($string, "aes-128-cbc", hex2bin($aKey), 0, hex2bin($sIv));
// echo '<br>ciphertext:      ' . $ciphertext . PHP_EOL;
$ciphertextUrlencoded = urlencode($ciphertext);
echo $ciphertextUrlencoded . PHP_EOL;
?>