<?php
use Spiral\Goridge;
require "vendor/autoload.php";

$rpc = new Goridge\RPC(new Goridge\SocketRelay("go", 6001));

echo $rpc->call("App.SumArray", "[1,2,3]");
