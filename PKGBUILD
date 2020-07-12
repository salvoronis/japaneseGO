# This is an example PKGBUILD file. Use this as a start to creating your own, 
# and remove these comments. For more information, see 'man PKGBUILD'. 
# NOTE: Please fill out the license field for your package! If it is unknown, 
# then please put 'unknown'. 
# Maintainer: Your Name <youremail@domain.com> 
pkgname=kanji 
pkgver=0.1 
pkgrel=6
epoch= 
pkgdesc="Test to n5-n2 kanji knowlage" 
arch=('i386' 'x86_64') 
url="" 
license=('unknown') 
groups=() 
depends=('go') 
makedepends=() 
checkdepends=() 
optdepends=() 
provides=() 
conflicts=() 
replaces=() 
backup=() 
options=() 
install= 
changelog= 
source=("kanji::git+https://github.com/salvoronis/japaneseGO.git#branch=master") 
noextract=() 
md5sums=('SKIP') 
validpgpkeys=() 
build() {
	cd "kanji"
	make build
} 
package() { 
	echo "$pkgdir"/usr/bin/$pkgname
	cd kanji
	install  -Dm755 bin/kanjitest "$pkgdir"/usr/bin/$pkgname
}
