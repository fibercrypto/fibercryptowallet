echo '\033[1;37mUpdating languages: [EN_US, ES_ES, FR_FR]\033[0m'
echo '\033[1;37mSearching for strings to translate in QML sources...\033[0m'
lupdate ../../src/ui/*.qml ../../src/ui/Controls/*.qml ../../src/ui/Delegates/*.qml ../../src/ui/Dialogs/*.qml -source-language en_us -target-language en_us -ts FiberCryptoWallet_en.ts
lupdate ../../src/ui/*.qml ../../src/ui/Controls/*.qml ../../src/ui/Delegates/*.qml ../../src/ui/Dialogs/*.qml -source-language en_us -target-language es_es -ts FiberCryptoWallet_es.ts
lupdate ../../src/ui/*.qml ../../src/ui/Controls/*.qml ../../src/ui/Delegates/*.qml ../../src/ui/Dialogs/*.qml -source-language en_us -target-language fr_fr -ts FiberCryptoWallet_fr.ts
# echo '\033[1;37mSearching for strings to translate in Go sources...\033[0m'
# find ../../src -path *.go -type f -printf "%p " | lupdate -tr-function-alias translate+=Translate -source-language en_us -target-language en_us -ts FiberCryptoWallet_en.ts