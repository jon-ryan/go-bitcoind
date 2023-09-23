bitcoind
===========

A Golang client library wrapping the bitcoind JSON RPC API


Installation
-----
	$ go get https://github.com/jon-ryan/go-bitcoind


Usage
----

	package main

	import (
		"github.com/jon-ryan/go-bitcoind"
		"log"
	)

	const (
		SERVER_HOST        = "You server host"
		SERVER_PORT        = port (int)
		USER               = "user"
		PASSWD             = "passwd"
		USESSL             = false
		WALLET_PASSPHRASE  = "WalletPassphrase"
	)

	func main() {
		bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
		if err != nil {
			log.Fatalln(err)
		}

		//walletpassphrase
		err = bc.WalletPassphrase(WALLET_PASSPHRASE, 3600)
		log.Println(err)

		// backupwallet
		err = bc.BackupWallet("/tmp/wallet.dat")
		log.Println(err)


		// dumpprivkey
		privKey, err := bc.DumpPrivKey("1KU5DX7jKECLxh1nYhmQ7CahY7GMNMVLP3")
		log.Println(err, privKey)

	}
	
Mores examples in example.go (in examples folder) 






##### Note on SSL support 

Note on ssl support : bitcoind library doesn't verify the server's certificate chain. That means that it accepts any certificate presented by the server and any host name in that certificate. In this mode, TLS is susceptible to man-in-the-middle attacks.

