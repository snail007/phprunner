package main

import (
	"os"

	php "github.com/deuill/go-php"
)

func main() {
	if len(os.Args) != 2 {
		println(os.Args[0] + " filename.php")
		os.Exit(0)
	}
	engine, _ := php.New()
	context, _ := engine.NewContext()
	context.Output = os.Stdout
	err := context.Exec(os.Args[1])
	engine.Destroy()
	if err != nil {
		println(err.Error())
	}

}

// package main

// // #cgo CFLAGS: -I /home/pengmeng/linux_php_installer/libs/php-7.0.8/ -I /home/pengmeng/linux_php_installer/libs/php-7.0.8/main/ -I /home/pengmeng/linux_php_installer/libs/php-7.0.8/Zend/ -I /home/pengmeng/linux_php_installer/libs/php-7.0.8/TSRM/
// // #cgo LDFLAGS: -lphp7
// // #include "sapi/embed/php_embed.h"
// /*
// int print(char * script) {
//     PHP_EMBED_START_BLOCK(0,NULL);
//     zend_eval_string(script,NULL,"" TSRMLS_CC);
//     PHP_EMBED_END_BLOCK();
//     return 1;
// }
// */
// import "C"

// func main() {
// 	// if len(os.Args) != 2 {
// 	// 	println(os.Args[0] + " filename.php")
// 	// 	os.Exit(0)
// 	// }
// 	// s, e := ioutil.ReadFile(os.Args[1])
// 	// if e != nil {
// 	// 	println(e.Error())
// 	// 	os.Exit(0)
// 	// }
// 	// println(string(s))
// 	// C.print(C.CString(string(s)))
// }
